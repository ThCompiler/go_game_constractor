package scene

import (
	"encoding/xml"

	"github.com/ThCompiler/go_game_constractor/scg/expr/parser"
)

const nSceneMatcherFields = 2

type Matcher struct {
	Name    string
	ToScene string
}

type _matcher struct {
	Name    string `yaml:"name" json:"name" xml:"name"`
	ToScene string `yaml:"toScene,omitempty" json:"to_scene,omitempty" xml:"toScene,omitempty"`
}

func (m *Matcher) UnmarshalYAML(un func(interface{}) error) (err error) {
	return m.unmarshal(&parser.UnmarshalerYAML{UN: un})
}

func (m *Matcher) UnmarshalJSON(bs []byte) error {
	return m.unmarshal(&parser.UnmarshalerJSON{Bs: bs})
}

func (m *Matcher) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return m.unmarshal(&parser.UnmarshalerXML{D: d, Start: start})
}

func (m *Matcher) unmarshal(unm parser.Unmarshaler) (err error) {
	tmp := make(map[string]interface{})
	if err = unm.Unmarshal(&tmp); err != nil {
		if err = unm.Unmarshal(&m.Name); err == nil {
			return nil
		}

		return err
	}

	tmpMatcher := _matcher{}
	if err = unm.Unmarshal(&tmpMatcher); err == nil {
		m.Name = tmpMatcher.Name
		m.ToScene = tmpMatcher.ToScene

		return nil
	}

	if len(tmp) > nSceneMatcherFields {
		return ErrorTooManyFields
	}

	return err
}
