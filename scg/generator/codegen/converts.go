package codegen

import (
	"bytes"
	"encoding/json"
	"github.com/ChimeraCoder/gojson"
)

// ConvertToGoStruct generate golang struct on any unknown types in string
func ConvertToGoStruct(str interface{}, structName string, packageName string) (string, error) {
	// imagine any struct as json
	jsonStr, err := json.Marshal(str)
	if err != nil {
		return "", err
	}

	// Generate go structs from json
	r := bytes.NewReader(jsonStr)
	result, er := gojson.Generate(
		r,
		gojson.ParseJson,
		structName,
		packageName,
		[]string{"json"},
		true,
		true,
	)

	if er != nil {
		return "", er
	}

	// convert out to string
	return string(result[:]), nil
}
