package cleanenv

// =================================================================
// Get by github.com/ilyakaznacheev/cleanenv and change yaml parser
// =================================================================

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"olympos.io/encoding/edn"

	"github.com/ThCompiler/go_game_constractor/pkg/structures"
)

type ConfigType int64

const (
	YAML ConfigType = iota
	JSON
	TOML
	EDN
	ENV
	XML
)

const (
	// DefaultSeparator is a default list and map separator character
	DefaultSeparator = ","
)

// Supported tags.
const (
	// TagEnv - name of the environment variable or a list of names.
	TagEnv = "env"

	// TagEnvLayout -value parsing layout (for types like time.Time).
	TagEnvLayout = "env-layout"

	// TagEnvDefault - default value.
	TagEnvDefault = "env-default"

	// TagEnvSeparator - custom list and map separator.
	TagEnvSeparator = "env-separator"

	// TagEnvDescription - environment variable description.
	TagEnvDescription = "env-description"

	// TagEnvUpd - flag to mark a field as updatable.
	TagEnvUpd = "env-upd"

	// TagEnvRequired - flag to mark a field as required.
	TagEnvRequired = "env-required"

	// TagEnvPrefix - flag to specify prefix for structure fields.
	TagEnvPrefix = "env-prefix"
)

// Setter is an interface for a custom value setter.
//
// To implement a custom value setter you need to add a SetValue function to your
// type that will receive a string raw value:
//
//	type MyField string
//
//	func (f *MyField) SetValue(s string) error {
//		if s == "" {
//			return fmt.Errorf("field value can't be empty")
//		}
//		*f = MyField("my field is: " + s)
//		return nil
//	}
type Setter interface {
	SetValue(string) error
}

// Updater gives an ability to implement custom update function for a field or a whole structure.
type Updater interface {
	Update() error
}

// ReadConfig reads configuration file and parses it depending on tags in structure provided.
// Then it reads and parses
//
// Example:
//
//	type ConfigDatabase struct {
//		Port     string `yaml:"port" env:"PORT" env-default:"5432"`
//		Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
//		Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
//		User     string `yaml:"user" env:"USER" env-default:"user"`
//		Password string `yaml:"password" env:"PASSWORD"`
//	}
//
//	var cfg ConfigDatabase
//
//	err := cleanenv.ReadConfig("config.yml", &cfg)
//	if err != nil {
//	    ...
//	}
func ReadConfig(path string, cfg interface{}) error {
	if err := parseFile(path, cfg); err != nil {
		return err
	}

	return readEnvVars(cfg, false)
}

func ReadConfigFromReader(reader io.Reader, ext ConfigType, cfg interface{}) error {
	if err := parseReader(reader, ext, cfg); err != nil {
		return err
	}

	return readEnvVars(cfg, false)
}

// ReadEnv reads environment variables into the structure.
func ReadEnv(cfg interface{}) error {
	return readEnvVars(cfg, false)
}

// UpdateEnv rereads (updates) environment variables in the structure.
func UpdateEnv(cfg interface{}) error {
	return readEnvVars(cfg, true)
}

// parseFile parses configuration file according to it's extension
//
// Currently following file extensions are supported:
//
// - yaml
//
// - json
//
// - toml
//
// - env
//
// - edn.
func parseFile(path string, cfg interface{}) error {
	// open the configuration file
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	// parse the file depending on the file type
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		err = parseYAML(f, cfg)
	case ".json":
		err = parseJSON(f, cfg)
	case ".toml":
		err = parseTOML(f, cfg)
	case ".edn":
		err = parseEDN(f, cfg)
	case ".env":
		err = parseENV(f, cfg)
	case ".xml":
		err = parseXML(f, cfg)
	default:
		return errorNotSupportedFileFormatAsString(ext)
	}

	if err != nil {
		return errors.Wrap(err, "config file parsing error")
	}

	return nil
}

// parseReader parses configuration from some reader according to it's extension
//
// Currently following extensions are supported:
//
// - yaml
//
// - json
//
// - toml
//
// - env
//
// - edn.
func parseReader(reader io.Reader, ext ConfigType, cfg interface{}) (err error) {
	// parse the file depending on the file type
	switch ext {
	case YAML:
		err = parseYAML(reader, cfg)
	case JSON:
		err = parseJSON(reader, cfg)
	case TOML:
		err = parseTOML(reader, cfg)
	case EDN:
		err = parseEDN(reader, cfg)
	case ENV:
		err = parseENV(reader, cfg)
	case XML:
		err = parseXML(reader, cfg)
	default:
		return errorNotSupportedFileFormatAsConfigType(ext)
	}

	if err != nil {
		return errors.Wrap(err, "config reader parsing error")
	}

	return nil
}

// parseYAML parses YAML from reader to data structure.
func parseYAML(r io.Reader, str interface{}) error {
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)

	return dec.Decode(str)
}

// parseJSON parses JSON from reader to data structure.
func parseJSON(r io.Reader, str interface{}) error {
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()

	return d.Decode(str)
}

// parseTOML parses TOML from reader to data structure.
func parseTOML(r io.Reader, str interface{}) error {
	_, err := toml.NewDecoder(r).Decode(str)

	return err
}

// parseEDN parses EDN from reader to data structure.
func parseEDN(r io.Reader, str interface{}) error {
	return edn.NewDecoder(r).Decode(str)
}

// parseXML parses XML from reader to data structure.
func parseXML(r io.Reader, str interface{}) error {
	return xml.NewDecoder(r).Decode(str)
}

// parseENV, in fact, doesn't fill the structure with environment variable values.
// It just parses ENV file and sets all variables to the environment.
// Thus, the structure should be filled at the next steps.
func parseENV(r io.Reader, _ interface{}) error {
	vars, err := godotenv.Parse(r)
	if err != nil {
		return err
	}

	for env, val := range vars {
		if err = os.Setenv(env, val); err != nil {
			return fmt.Errorf("set environment: %w", err)
		}
	}

	return nil
}

// parseSlice parses value into a slice of given type.
func parseSlice(valueType reflect.Type, value, sep string, layout *string) (*reflect.Value, error) {
	sliceValue := reflect.MakeSlice(valueType, 0, 0)
	if valueType.Elem().Kind() == reflect.Uint8 {
		sliceValue = reflect.ValueOf([]byte(value))
	} else if len(strings.TrimSpace(value)) != 0 {
		values := strings.Split(value, sep)
		sliceValue = reflect.MakeSlice(valueType, len(values), len(values))

		for i, val := range values {
			if err := parseValue(sliceValue.Index(i), val, sep, layout); err != nil {
				return nil, err
			}
		}
	}

	return &sliceValue, nil
}

const partNumber = 2

// parseMap parses value into a map of given type.
func parseMap(valueType reflect.Type, value, sep string, layout *string) (*reflect.Value, error) {
	mapValue := reflect.MakeMap(valueType)

	if len(strings.TrimSpace(value)) != 0 {
		pairs := strings.Split(value, sep)

		for _, pair := range pairs {
			kvPair := strings.SplitN(pair, ":", partNumber)
			if len(kvPair) != partNumber {
				return nil, errorInvalidMapItem(pair)
			}

			k := reflect.New(valueType.Key()).Elem()

			err := parseValue(k, kvPair[0], sep, layout)
			if err != nil {
				return nil, err
			}

			v := reflect.New(valueType.Elem()).Elem()

			err = parseValue(v, kvPair[1], sep, layout)
			if err != nil {
				return nil, err
			}

			mapValue.SetMapIndex(k, v)
		}
	}

	return &mapValue, nil
}

// structMeta is a structure metadata entity.
type structMeta struct {
	envList     []string
	fieldName   string
	fieldValue  reflect.Value
	defValue    *string
	layout      *string
	separator   string
	description string
	updatable   bool
	required    bool
}

// isFieldValueZero determines if fieldValue empty or not.
func (sm *structMeta) isFieldValueZero() bool {
	return sm.fieldValue.IsZero()
}

// parseFunc custom value parser function
type parseFunc func(*reflect.Value, string, *string) error

// Any specific supported struct can be added here
var validStructs = map[reflect.Type]parseFunc{
	reflect.TypeOf(time.Time{}): func(field *reflect.Value, value string, layout *string) error {
		var l string
		if layout != nil {
			l = *layout
		} else {
			l = time.RFC3339
		}

		val, err := time.Parse(l, value)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(val))

		return nil
	},

	reflect.TypeOf(url.URL{}): func(field *reflect.Value, value string, _ *string) error {
		val, err := url.Parse(value)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(*val))

		return nil
	},

	reflect.TypeOf(&time.Location{}): func(field *reflect.Value, value string, _ *string) error {
		loc, err := time.LoadLocation(value)
		if err != nil {
			return err
		}

		field.Set(reflect.ValueOf(loc))

		return nil
	},
}

type cfgNode struct {
	Val    interface{}
	Prefix string
}

// readStructMetadata reads structure metadata (types, tags, etc.)
func readStructMetadata(cfgRoot interface{}) ([]structMeta, error) {
	cfgStack := structures.NewStack[cfgNode]()
	cfgStack.Push(cfgNode{cfgRoot, ""})

	metas := make([]structMeta, 0)

	for !cfgStack.Empty() {
		node := cfgStack.MustPop()
		s := reflect.ValueOf(node.Val)
		sPrefix := node.Prefix

		// unwrap pointer
		if s.Kind() == reflect.Ptr {
			s = s.Elem()
		}

		// process only structures
		if s.Kind() != reflect.Struct {
			return nil, errorWrongTypeOfField(s.Kind())
		}

		metas = readTags(metas, s, cfgStack, sPrefix)
	}

	return metas, nil
}

func readTags(metas []structMeta, s reflect.Value, cfgStack structures.Stack[cfgNode], sPrefix string) []structMeta {
	typeInfo := s.Type()

	// read tags
	for idx := 0; idx < s.NumField(); idx++ {
		fType := typeInfo.Field(idx)

		var layout *string

		// process nested structure (except of supported ones)
		if fld := s.Field(idx); fld.Kind() == reflect.Struct {
			// add structure to parsing structures
			if _, found := validStructs[fld.Type()]; !found {
				prefix, _ := fType.Tag.Lookup(TagEnvPrefix)
				cfgStack.Push(cfgNode{fld.Interface(), sPrefix + prefix})

				continue
			}

			// process time.Time
			if l, ok := fType.Tag.Lookup(TagEnvLayout); ok {
				layout = &l
			}
		}

		defValue, separator, envList, upd, required := parseStruct(fType, sPrefix)

		metas = append(metas, structMeta{
			envList:     envList,
			fieldName:   s.Type().Field(idx).Name,
			fieldValue:  s.Field(idx),
			defValue:    defValue,
			layout:      layout,
			separator:   separator,
			description: fType.Tag.Get(TagEnvDescription),
			updatable:   upd,
			required:    required,
		})
	}

	return metas
}

func parseStruct(fType reflect.StructField, sPrefix string) (*string, string, []string, bool, bool) {
	var (
		defValue  *string
		separator string
	)

	if def, ok := fType.Tag.Lookup(TagEnvDefault); ok {
		defValue = &def
	}

	if sep, ok := fType.Tag.Lookup(TagEnvSeparator); ok {
		separator = sep
	} else {
		separator = DefaultSeparator
	}

	_, upd := fType.Tag.Lookup(TagEnvUpd)

	_, required := fType.Tag.Lookup(TagEnvRequired)

	envList := make([]string, 0)

	if envs, ok := fType.Tag.Lookup(TagEnv); ok && len(envs) != 0 {
		envList = strings.Split(envs, DefaultSeparator)
		if sPrefix != "" {
			for i := range envList {
				envList[i] = sPrefix + envList[i]
			}
		}
	}

	return defValue, separator, envList, upd, required
}

// readEnvVars reads environment variables to the provided configuration structure
func readEnvVars(cfg interface{}, update bool) error {
	metaInfo, err := readStructMetadata(cfg)
	if err != nil {
		return err
	}

	if updater, ok := cfg.(Updater); ok {
		if err := updater.Update(); err != nil {
			return err
		}
	}

	for _, meta := range metaInfo {
		// update only updatable fields
		if update && !meta.updatable {
			continue
		}

		var rawValue *string

		for _, env := range meta.envList {
			if value, ok := os.LookupEnv(env); ok {
				rawValue = &value

				break
			}
		}

		if rawValue == nil && meta.required && meta.isFieldValueZero() {
			return errorFieldNotProvided(meta.fieldName)
		}

		if rawValue == nil && meta.isFieldValueZero() {
			rawValue = meta.defValue
		}

		if rawValue == nil {
			continue
		}

		if err := parseValue(meta.fieldValue, *rawValue, meta.separator, meta.layout); err != nil {
			return err
		}
	}

	return nil
}

// parseValue parses value into the corresponding field.
// In case of maps and slices it uses provided separator to split raw value string
func parseValue(field reflect.Value, value, sep string, layout *string) error {
	if ok, err := parseFromInterface(&field, value); ok {
		return err
	}

	return parseFromBaseTypes(&field, value, sep, layout)
}

type parseFuncType func(valueType reflect.Type, value, sep string, layout *string, field *reflect.Value) error

var parsersMap = map[reflect.Kind]parseFuncType{
	reflect.Bool:    parseValueToBool,
	reflect.Int:     parseValueToInt81632,
	reflect.Int8:    parseValueToInt81632,
	reflect.Int16:   parseValueToInt81632,
	reflect.Int32:   parseValueToInt81632,
	reflect.Int64:   parseValueToInt81632,
	reflect.Uint:    parseValueToUint,
	reflect.Uint8:   parseValueToUint,
	reflect.Uint16:  parseValueToUint,
	reflect.Uint32:  parseValueToUint,
	reflect.Uint64:  parseValueToUint,
	reflect.Float32: parseValueToFloat,
	reflect.Float64: parseValueToFloat,
	reflect.Slice:   parseValueToSlice,
	reflect.Map:     parseValueToMap,
}

func parseFromBaseTypes(field *reflect.Value, value, sep string, layout *string) error {
	valueType := field.Type()

	if fun, ok := parsersMap[valueType.Kind()]; ok {
		return fun(valueType, value, sep, layout, field)
	} else {
		return parseFromOtherSupportedTypes(valueType, value, layout, field)
	}
}

func parseFromInterface(field *reflect.Value, value string) (bool, error) {
	if field.CanInterface() {
		if cs, ok := field.Interface().(Setter); ok {
			return true, cs.SetValue(value)
		} else if csp, ok := field.Addr().Interface().(Setter); ok {
			return true, csp.SetValue(value)
		}
	}

	return false, nil
}

func parseValueToInt64(valueType reflect.Type, value, _ string, _ *string, field *reflect.Value) error {
	if valueType == reflect.TypeOf(time.Duration(0)) {
		// try to parse time
		d, err := time.ParseDuration(value)
		if err != nil {
			return err
		}

		field.SetInt(int64(d))
	} else {
		// parse regular integer
		number, err := strconv.ParseInt(value, 0, valueType.Bits())
		if err != nil {
			return err
		}

		field.SetInt(number)
	}

	return nil
}

func parseFromOtherSupportedTypes(valueType reflect.Type, value string, layout *string, field *reflect.Value) error {
	if structParser, found := validStructs[valueType]; found {
		return structParser(field, value, layout)
	}

	return errorNotSupportedType(valueType.PkgPath(), valueType.Name())
}

func parseValueToMap(valueType reflect.Type, value, sep string, layout *string, field *reflect.Value) error {
	mapValue, err := parseMap(valueType, value, sep, layout)
	if err != nil {
		return err
	}

	field.Set(*mapValue)

	return nil
}

func parseValueToSlice(valueType reflect.Type, value, sep string, layout *string, field *reflect.Value) error {
	sliceValue, err := parseSlice(valueType, value, sep, layout)
	if err != nil {
		return err
	}

	field.Set(*sliceValue)

	return nil
}

func parseValueToBool(_ reflect.Type, value, _ string, _ *string, field *reflect.Value) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}

	field.SetBool(b)

	return nil
}

func parseValueToInt81632(valueType reflect.Type, value, _ string, _ *string, field *reflect.Value) error {
	number, err := strconv.ParseInt(value, 0, valueType.Bits())
	if err != nil {
		return err
	}

	field.SetInt(number)

	return nil
}

func parseValueToUint(valueType reflect.Type, value, _ string, _ *string, field *reflect.Value) error {
	number, err := strconv.ParseUint(value, 0, valueType.Bits())
	if err != nil {
		return err
	}

	field.SetUint(number)

	return nil
}

func parseValueToFloat(valueType reflect.Type, value, _ string, _ *string, field *reflect.Value) error {
	number, err := strconv.ParseFloat(value, valueType.Bits())
	if err != nil {
		return err
	}

	field.SetFloat(number)

	return nil
}

// GetDescription returns a description of environment variables.
// You can provide a custom header text.
func GetDescription(cfg interface{}, headerText *string) (string, error) {
	meta, err := readStructMetadata(cfg)
	if err != nil {
		return "", err
	}

	var header, description string

	if headerText != nil {
		header = *headerText
	} else {
		header = "Environment variables:"
	}

	for _, m := range meta {
		if len(m.envList) == 0 {
			continue
		}

		for idx, env := range m.envList {
			elemDescription := fmt.Sprintf("\n  %s %s", env, m.fieldValue.Kind())
			if idx > 0 {
				elemDescription += fmt.Sprintf(" (alternative to %s)", m.envList[0])
			}

			elemDescription += fmt.Sprintf("\n    \t%s", m.description)
			if m.defValue != nil {
				elemDescription += fmt.Sprintf(" (default %q)", *m.defValue)
			}

			description += elemDescription
		}
	}

	if description != "" {
		return header + description, nil
	}

	return "", nil
}

// Usage returns a configuration usage help.
// Other usage instructions can be wrapped in and executed before this usage function.
// The default output is STDERR.
func Usage(cfg interface{}, headerText *string, usageFuncs ...func()) func() {
	return FUsage(os.Stderr, cfg, headerText, usageFuncs...)
}

// FUsage prints configuration help into the custom output.
// Other usage instructions can be wrapped in and executed before this usage function
func FUsage(w io.Writer, cfg interface{}, headerText *string, usageFuncs ...func()) func() {
	return func() {
		for _, fn := range usageFuncs {
			fn()
		}

		_ = flag.Usage

		text, err := GetDescription(cfg, headerText)
		if err != nil {
			return
		}

		if len(usageFuncs) > 0 {
			_, _ = fmt.Fprintln(w)
		}

		_, _ = fmt.Fprintln(w, text)
	}
}
