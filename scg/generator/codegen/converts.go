package codegen

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/ChimeraCoder/gojson"
)

const NumberPartOfString = 2

// ConvertToGoStruct generate golang struct on any unknown types in string.
func ConvertToGoStruct(str interface{}, structName string) (string, error) {
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
		structName,
		[]string{"json"},
		true,
		true,
	)

	if er != nil {
		return "", er
	}

	res := strings.SplitN(string(result), "\n\n", NumberPartOfString)
	// convert out to string
	return res[1], nil
}
