package parser

import (
	"encoding/json"
	"encoding/xml"
)

type Unmarshaler interface {
	Unmarshal(value interface{}) error
}

type UnmarshalerYAML struct {
	UN func(interface{}) error
}

func (u *UnmarshalerYAML) Unmarshal(value interface{}) error {
	return u.UN(value)
}

type UnmarshalerJSON struct {
	Bs []byte
}

func (u *UnmarshalerJSON) Unmarshal(value interface{}) error {
	return json.Unmarshal(u.Bs, value)
}

type UnmarshalerXML struct {
	D     *xml.Decoder
	Start xml.StartElement
}

func (u *UnmarshalerXML) Unmarshal(value interface{}) error {
	return u.D.DecodeElement(value, &u.Start)
}

type UnmarshalFunc func(Unmarshaler) error

type MultiParser struct {
	Fun UnmarshalFunc
}

func (mp *MultiParser) UnmarshalYAML(un func(interface{}) error) (err error) {
	return mp.Fun(&UnmarshalerYAML{UN: un})
}

func (mp *MultiParser) UnmarshalJSON(bs []byte) error {
	return mp.Fun(&UnmarshalerJSON{Bs: bs})
}

func (mp *MultiParser) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return mp.Fun(&UnmarshalerXML{D: d, Start: start})
}
