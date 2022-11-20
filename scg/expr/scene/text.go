package scene

import (
    "github.com/ThCompiler/go_game_constractor/scg/go/types"
    "strings"
)

type Text struct {
    Text   string           `yaml:"string" json:"text" xml:"text"`
    TTS    string           `yaml:"tts" json:"tts" xml:"tts"`
    Values map[string]Value `yaml:"values,omitempty" json:"values,omitempty" xml:"values,omitempty"`
}

func (t *Text) IsValid() (bool, error) {
    err := t.checkValuesType()
    if err != nil {
        return false, err
    }
    err = t.checkTextOnContainsValues()
    return err == nil, err
}

func (t *Text) checkTextOnContainsValues() (err error) {
    err = nil
    for key := range t.Values {
        if !strings.Contains(t.Text, "{"+key+"}") {
            err = errorNotFoundValueInText(key)
            break
        }
        if !strings.Contains(t.TTS, "{"+key+"}") {
            err = errorNotFoundValueInText(key)
            break
        }
    }
    return
}

func (t *Text) checkValuesType() (err error) {
    err = nil
    for _, val := range t.Values {
        if !types.IsValidType(val.Type) {
            err = errorUnknownTypeOfValue(val.Type)
            break
        }
    }
    return
}
