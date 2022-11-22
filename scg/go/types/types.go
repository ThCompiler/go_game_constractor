package types

import "strings"

func init() {
    initTypes()
}

func initTypes() {
    arrayTypes := make([]string, len(mapTypes))
    index := 0
    for key := range mapTypes {
        arrayTypes[index] = key
        index++
    }
    types = strings.Join(arrayTypes, ", ")
}

var types = ""

type ConvertFunc func(value any) (any, error)

type GoType struct {
    GoName           string
    ConvertFunctions map[string]ConvertFunc
}

var mapTypes = map[string]string{
    "bool":        "bool",
    "int":         "int",
    "int8":        "int8",
    "int16":       "int16",
    "int32":       "int32",
    "int64":       "int64",
    "uint":        "uint",
    "uint8":       "uint8",
    "uint16":      "uint16",
    "uint32":      "uint32",
    "uint64":      "uint64",
    "float32":     "float32",
    "float64":     "float64",
    "string":      "string",
    "bytes":       "[]byte",
    "byte":        "byte",
    "rune":        "rune",
    "complex128":  "complex128",
    "complex64":   "complex64",
    "strings":     "[]string",
    "map_strings": "map[string]string",
    "map_string":  "map[string]any",
}

func AddTypes(name string, goType string) {
    mapTypes[name] = goType
}

func IsValidType(s string) bool {
    _, is := mapTypes[s]
    return is
}

func ToGoType(name string) string {
    return mapTypes[name]
}

func GetSupportTypes() string {
    return types
}
