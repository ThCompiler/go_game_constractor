package scene

import (
	"encoding/xml"

	"github.com/ThCompiler/go_game_constractor/scg/expr/parser"
	"github.com/ThCompiler/go_game_constractor/scg/go/types"
)

const nLoadValueFields = 1

type Context struct {
	SaveValue *SaveValue  `yaml:"saveValue" json:"save_value" xml:"saveValue"`
	LoadValue []LoadValue `yaml:"loadValue" json:"load_value" xml:"loadValue"`
}

func (ct *Context) checkValuesType() (err error) {
	err = nil

	if ct.SaveValue != nil {
		if !types.IsValidType(ct.SaveValue.Type) {
			err = errorUnknownTypeOfValue(ct.SaveValue.Type)
		}
	}

	return
}

type SaveValue struct {
	Name string `yaml:"name" json:"name" xml:"name"`
	Type string `yaml:"type" json:"type" xml:"type"`
}

type LoadValue struct {
	Name string
	Type string
}

type _loadValue struct {
	Name string
}

func (lv *LoadValue) UnmarshalYAML(un func(interface{}) error) (err error) {
	return lv.unmarshal(&parser.UnmarshalerYAML{UN: un})
}

func (lv *LoadValue) UnmarshalJSON(bs []byte) error {
	return lv.unmarshal(&parser.UnmarshalerJSON{Bs: bs})
}

func (lv *LoadValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return lv.unmarshal(&parser.UnmarshalerXML{D: d, Start: start})
}

func (lv *LoadValue) unmarshal(unm parser.Unmarshaler) (err error) {
	tmp := make(map[string]interface{})
	if err = unm.Unmarshal(&tmp); err != nil {
		if err = unm.Unmarshal(&lv.Name); err == nil {
			return nil
		}

		return err
	}

	tmpMatcher := _loadValue{}
	if err = unm.Unmarshal(&tmpMatcher); err == nil {
		lv.Name = tmpMatcher.Name

		return nil
	}

	if len(tmp) > nLoadValueFields {
		return ErrorTooManyFields
	}

	return err
}
