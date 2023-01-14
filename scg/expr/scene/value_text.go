package scene

import (
	"encoding/xml"

	"gopkg.in/yaml.v3"

	"github.com/ThCompiler/go_game_constractor/scg/expr/parser"
)

type Value struct {
	Type        string
	FromContext string
}

type _value struct {
	Type        string `yaml:"type" json:"type" xml:"type"`
	FromContext string `yaml:"fromContext" json:"from_context" xml:"fromContext"`
}

func (v *Value) UnmarshalYAML(n *yaml.Node) (err error) {
	return v.unmarshal(&parser.UnmarshalerYAML{N: n})
}

func (v *Value) UnmarshalJSON(bs []byte) error {
	return v.unmarshal(&parser.UnmarshalerJSON{Bs: bs})
}

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return v.unmarshal(&parser.UnmarshalerXML{D: d, Start: start})
}

func (v *Value) unmarshal(unm parser.Unmarshaler) (err error) {
	tmp := make(map[string]interface{})
	if err = unm.Unmarshal(&tmp); err != nil {
		if err = unm.Unmarshal(&v.Type); err == nil {
			return nil
		}

		return err
	}

	tmpMatcher := _value{}
	if err = unm.Unmarshal(&tmpMatcher); err == nil {
		v.Type = tmpMatcher.Type
		v.FromContext = tmpMatcher.FromContext

		return nil
	}

	if len(tmp) > 1 {
		return ErrorTooManyFields
	}

	return err
}
