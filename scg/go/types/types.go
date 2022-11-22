package types

import (
    "reflect"
    "strings"
)

func init() {
    AddTypes("bool", GoType{"bool", ToConvertFunc(ConvertStringToBool)})
    AddTypes("int", GoType{"int", ToConvertFunc(ConvertStringToInt)})
    AddTypes("int8", GoType{"int8", ToConvertFunc(ConvertStringToInt8)})
    AddTypes("int16", GoType{"int16", ToConvertFunc(ConvertStringToInt16)})
    AddTypes("int32", GoType{"int32", ToConvertFunc(ConvertStringToInt32)})
    AddTypes("int64", GoType{"int64", ToConvertFunc(ConvertStringToInt64)})
    AddTypes("uint", GoType{"uint", ToConvertFunc(ConvertStringToUint)})
    AddTypes("uint8", GoType{"uint8", ToConvertFunc(ConvertStringToUint8)})
    AddTypes("uint16", GoType{"uint16", ToConvertFunc(ConvertStringToUint16)})
    AddTypes("uint32", GoType{"uint32", ToConvertFunc(ConvertStringToUint32)})
    AddTypes("uint64", GoType{"uint64", ToConvertFunc(ConvertStringToUint64)})
    AddTypes("float32", GoType{"float32", ToConvertFunc(ConvertStringToFloat32)})
    AddTypes("float64", GoType{"float64", ToConvertFunc(ConvertStringToFloat64)})
    AddTypes("string", GoType{"string", ToConvertFunc(func(value string) (string, error) { return value, nil })})
    AddTypes("bytes", GoType{"[]byte", ToConvertFunc(ConvertStringToBytes)})
    AddTypes("byte", GoType{"byte", ToConvertFunc(ConvertStringToByte)})
    AddTypes("rune", GoType{"rune", ToConvertFunc(ConvertStringToRune)})
    AddTypes("complex128", GoType{"complex128", ToConvertFunc(ConvertStringToComplex128)})
    AddTypes("complex64", GoType{"complex64", ToConvertFunc(ConvertStringToComplex64)})
    AddTypes("strings", GoType{"[]string", ToConvertFunc(ConvertStringToStrings)})

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

type ConvertFunc func(value string) (any, error)

func ToConvertFunc[T any](convertFunc func(value string) (T, error)) ConvertFunc {
    return func(value string) (any, error) {
        return convertFunc(value)
    }
}

type GoType struct {
    GoName          string
    ConvertFunction ConvertFunc
}

var mapTypes = make(map[string]GoType)

var unmapTypes = make(map[string]string)

func AddTypes(name string, goType GoType) {
    mapTypes[name] = goType
    unmapTypes[goType.GoName] = name
}

func IsValidType(s string) bool {
    _, is := mapTypes[s]
    return is
}

func ToGoType(name string) string {
    return mapTypes[name].GoName
}

func GetSupportTypes() string {
    return types
}

func MustConvert[T any](str string) T {
    var res T
    if tp, is := unmapTypes[reflect.TypeOf(res).String()]; is {
        val, err := mapTypes[tp].ConvertFunction(str)
        if err != nil {
            panic(err)
        }
        res = val.(T)
    } else {
        panic(ErrorNotSupportedType)
    }
    return res
}
