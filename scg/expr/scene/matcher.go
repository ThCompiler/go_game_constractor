package scene

import (
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v3"
)

type typeMatcher int64

const (
	standard = 0
	selects  = 1
	regex    = 2
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

type Matcher struct {
	name          string
	regexMatcher  *RegexMatcher
	selectMatcher *SelectMatcher
	typeMatch     typeMatcher
}

func (m *Matcher) GetRegexMatcher() (*RegexMatcher, error) {
	if m.IsRegexMatcher() {
		return m.regexMatcher, nil
	}
	return nil, errorIsNotRegexMatcher
}

func (m *Matcher) GetSelectsMatcher() (*SelectMatcher, error) {
	if m.IsSelectMatcher() {
		return m.selectMatcher, nil
	}
	return nil, errorIsNotSelectsMatcher
}

func (m *Matcher) GetStandardMatcher() (string, error) {
	if m.IsDefaultMatcher() {
		return m.name, nil
	}
	return "", errorIsNotStandardMatcher
}

func (m *Matcher) MustSelectsMatcher() SelectMatcher {
	if m.IsSelectMatcher() {
		return *m.selectMatcher
	}
	panic(errorIsNotSelectsMatcher)
}

func (m *Matcher) MustRegexMatcher() RegexMatcher {
	if m.IsRegexMatcher() {
		return *m.regexMatcher
	}
	panic(errorIsNotRegexMatcher)
}

func (m *Matcher) MustStandardMatcher() string {
	if m.IsDefaultMatcher() {
		return m.name
	}
	panic(errorIsNotStandardMatcher)
}

func (m *Matcher) IsRegexMatcher() bool {
	return m.typeMatch == regex
}

func (m *Matcher) IsSelectMatcher() bool {
	return m.typeMatch == selects
}

func (m *Matcher) IsDefaultMatcher() bool {
	return m.typeMatch == standard
}

func (m *Matcher) UnmarshalYAML(n *yaml.Node) error {
	tmp := make(map[string]interface{})
	var err error
	if err = n.Decode(&tmp); err != nil {
		if err = n.Decode(&m.name); err == nil {
			m.typeMatch = standard
			return nil
		}
		m.name = ""
		return err
	}

	if len(tmp) == 2 {
		_, is := tmp["regex"]
		if err = n.Decode(&m.regexMatcher); err == nil && is {
			m.typeMatch = regex
			return nil
		}
		m.regexMatcher = nil

		if err = n.Decode(&m.selectMatcher); err == nil {
			m.typeMatch = selects
			return nil
		}
		m.selectMatcher = nil
	}

	return err
}

func (m *Matcher) UnmarshalJSON(bs []byte) error {
	var err error
	if err = json.Unmarshal(bs, &m.regexMatcher); err == nil {
		m.typeMatch = regex
		return nil
	}
	m.regexMatcher = nil

	if err = json.Unmarshal(bs, &m.selectMatcher); err == nil {
		m.typeMatch = selects
		return nil
	}
	m.selectMatcher = nil

	if err = json.Unmarshal(bs, &m.name); err == nil {
		m.typeMatch = standard
		return nil
	}
	m.name = ""

	return err
}

func (m *Matcher) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	if err = d.DecodeElement(&m.regexMatcher, &start); err == nil {
		m.typeMatch = regex
		return nil
	}
	m.regexMatcher = nil

	if err = d.DecodeElement(&m.selectMatcher, &start); err == nil {
		m.typeMatch = selects
		return nil
	}
	m.selectMatcher = nil

	if err = d.DecodeElement(&m.name, &start); err == nil {
		m.typeMatch = standard
		return nil
	}
	m.name = ""

	return err
}
