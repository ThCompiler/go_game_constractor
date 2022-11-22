package scene

import (
    "encoding/xml"
    "github.com/ThCompiler/go_game_constractor/scg/expr/parser"
    "gopkg.in/yaml.v3"
)

type typeMatcher int64

const (
    selects = 1
    regex   = 2
)

type RegexMatcher struct {
    Name        string `yaml:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
    Regex       string `yaml:"regex" json:"regex" xml:"regex"`
    NameMatched string `yaml:"nameMatched" json:"name_matched" xml:"nameMatched"`
}

type SelectMatcher struct {
    Name           string   `yaml:"name" json:"name" xml:"name"`
    Selects        []string `yaml:"selects" json:"selects" xml:"selects"`
    ReplaceMessage string   `yaml:"replaceMessage" json:"replaceMessage" xml:"replaceMessage"`
}

type ScriptMatcher struct {
    regexMatcher  *RegexMatcher
    selectMatcher *SelectMatcher
    typeMatch     typeMatcher
}

func (m *ScriptMatcher) SetName(name string) {
    if m.IsRegexMatcher() {
        m.regexMatcher.Name = name
    } else {
        m.selectMatcher.Name = name
    }
}

func (m *ScriptMatcher) GetRegexMatcher() (*RegexMatcher, error) {
    if m.IsRegexMatcher() {
        return m.regexMatcher, nil
    }
    return nil, ErrorIsNotRegexMatcher
}

func (m *ScriptMatcher) GetSelectsMatcher() (*SelectMatcher, error) {
    if m.IsSelectMatcher() {
        return m.selectMatcher, nil
    }
    return nil, ErrorIsNotSelectsMatcher
}

func (m *ScriptMatcher) MustSelectsMatcher() SelectMatcher {
    if m.IsSelectMatcher() {
        return *m.selectMatcher
    }
    panic(ErrorIsNotSelectsMatcher)
}

func (m *ScriptMatcher) MustRegexMatcher() RegexMatcher {
    if m.IsRegexMatcher() {
        return *m.regexMatcher
    }
    panic(ErrorIsNotRegexMatcher)
}

func (m *ScriptMatcher) IsRegexMatcher() bool {
    return m.typeMatch == regex
}

func (m *ScriptMatcher) IsSelectMatcher() bool {
    return m.typeMatch == selects
}

func (m *ScriptMatcher) UnmarshalYAML(n *yaml.Node) (err error) {
    return m.unmarshal(&parser.UnmarshalerYAML{N: n})
}

func (m *ScriptMatcher) UnmarshalJSON(bs []byte) error {
    return m.unmarshal(&parser.UnmarshalerJSON{Bs: bs})
}

func (m *ScriptMatcher) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    return m.unmarshal(&parser.UnmarshalerXML{D: d, Start: start})
}

func (m *ScriptMatcher) unmarshal(unm parser.Unmarshaler) (err error) {
    tmp := make(map[string]interface{})
    if err = unm.Unmarshal(&tmp); err != nil {
        return err
    }

    if _, is := tmp["regex"]; is {
        if err = unm.Unmarshal(&m.regexMatcher); err == nil {
            m.typeMatch = regex
            return nil
        }
        m.regexMatcher = nil
    } else {
        if err = unm.Unmarshal(&m.selectMatcher); err == nil {
            m.typeMatch = selects
            return nil
        }
        m.selectMatcher = nil
    }

    if len(tmp) > 3 {
        return ErrorTooManyFields
    }

    return err
}
