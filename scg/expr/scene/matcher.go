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
	Name  string `yaml:"name" json:"name" xml:"name"`
	Regex string `yaml:"regex" json:"regex" xml:"regex"`
}

type SelectMatcher struct {
	Name    string   `yaml:"name" json:"name" xml:"name"`
	Selects []string `yaml:"selects_in" json:"selects" xml:"selects"`
}

type Matcher struct {
	name          string `yaml:"name" json:"name" xml:"name"`
	regexMatcher  *RegexMatcher
	selectMatcher *SelectMatcher
	typeMatch     typeMatcher
}

func (m *Matcher) getRegexMatcher() (*RegexMatcher, error) {
	if m.isRegexMatcher() {
		return m.regexMatcher, nil
	}
	return nil, errorIsNotRegexMatcher
}

func (m *Matcher) getSelectsMatcher() (*SelectMatcher, error) {
	if m.isRegexMatcher() {
		return m.selectMatcher, nil
	}
	return nil, errorIsNotSelectsMatcher
}

func (m *Matcher) getStandardMatcher() (string, error) {
	if m.isRegexMatcher() {
		return m.name, nil
	}
	return "", errorIsNotStandardMatcher
}

func (m *Matcher) isRegexMatcher() bool {
	return m.typeMatch&regex == 1
}

func (m *Matcher) isSelectMatcher() bool {
	return m.typeMatch&selects == 1
}

func (m *Matcher) isDefaultMatcher() bool {
	return m.typeMatch == standard
}

func (m *Matcher) UnmarshalYAML(n *yaml.Node) error {
	var err error
	if err = n.Decode(&m.regexMatcher); err == nil {
		m.typeMatch = regex
		return nil
	}
	m.regexMatcher = nil

	if err = n.Decode(&m.selectMatcher); err == nil {
		m.typeMatch = selects
		return nil
	}
	m.selectMatcher = nil

	if err = n.Decode(&m.name); err == nil {
		m.typeMatch = standard
		return nil
	}
	m.name = ""

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
