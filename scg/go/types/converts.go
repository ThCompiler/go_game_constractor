package types

import (
    "fmt"
    "github.com/pkg/errors"
    "reflect"
    "strconv"
    "strings"
)

var (
    ErrorNotParsing = errors.New("can not parse type to other type")
)

func ConAnyToString(value any) (string, error) {
    switch tp := value.(type) {
    case bool:
        return strconv.FormatBool(tp), nil
    case int:
        return strconv.Itoa(tp), nil
    case int8:
        return strconv.FormatInt(int64(tp), 8), nil
    case int16:
        return strconv.FormatInt(int64(tp), 16), nil
    case int32:
        return strconv.FormatInt(int64(tp), 32), nil
    case int64:
        return strconv.FormatInt(tp, 64), nil
    case uint:
        return strconv.FormatUint(uint64(tp), 32), nil
    case uint8:
        return strconv.FormatUint(uint64(tp), 8), nil
    case uint16:
        return strconv.FormatUint(uint64(tp), 16), nil
    case uint32:
        return strconv.FormatUint(uint64(tp), 32), nil
    case uint64:
        return strconv.FormatUint(tp, 64), nil
    case float32:
        return strconv.FormatFloat(float64(tp), 'f', -1, 32), nil
    case float64:
        return strconv.FormatFloat(tp, 'f', -1, 64), nil
    case []byte:
        return string(tp), nil
    case complex64:
        return strconv.FormatComplex(complex128(tp), 'f', -1, 64), nil
    case complex128:
        return strconv.FormatComplex(tp, 'f', -1, 128), nil
    case []string:
        return strings.Join(tp, " "), nil
    }

    switch tp := value.(type) {
    case byte:
        return string([]byte{tp}), nil
    case rune:
        return string(tp), nil
    default:
        return fmt.Sprintf("%v", tp), nil
    }
}

func ConvertStringToBool(value string) (bool, error) {
    return strconv.ParseBool(value)
}

func ConvertStringToInt(value string) (int, error) {
    return strconv.Atoi(value)
}

func ConvertStringToInt8(value string) (int8, error) {
    res, err := strconv.ParseInt(value, 8, 10)
    return int8(res), err
}

func ConvertStringToInt8(value string) (int16, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[int16](tp), nil
    case string:
        res, err := strconv.ParseInt(tp, 16, 10)
        return int16(res), err
    case []byte:
        res, err := strconv.ParseInt(string(tp), 16, 10)
        return int16(res), err
    case []string:
        res, err := strconv.ParseInt(strings.Join(tp, ""), 16, 10)
        return int16(res), err
    }

    switch tp := value.(type) {
    case byte:
        return int16(tp), nil
    case rune:
        return int16(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type int16 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (int32, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[int32](tp), nil
    case string:
        res, err := strconv.ParseInt(tp, 32, 10)
        return int32(res), err
    case []byte:
        res, err := strconv.ParseInt(string(tp), 32, 10)
        return int32(res), err
    case []string:
        res, err := strconv.ParseInt(strings.Join(tp, ""), 32, 10)
        return int32(res), err
    }

    switch tp := value.(type) {
    case byte:
        return int32(tp), nil
    case rune:
        return int32(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type int32 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (int64, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[int64](tp), nil
    case string:
        return strconv.ParseInt(tp, 64, 10)
    case []byte:
        return strconv.ParseInt(string(tp), 64, 10)
    case []string:
        return strconv.ParseInt(strings.Join(tp, ""), 64, 10)
    }

    switch tp := value.(type) {
    case byte:
        return int64(tp), nil
    case rune:
        return int64(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type int64 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (uint, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[uint](tp), nil
    case string:
        res, err := strconv.ParseUint(tp, 64, 10)
        return uint(res), err
    case []byte:
        res, err := strconv.ParseUint(string(tp), 64, 10)
        return uint(res), err
    case []string:
        res, err := strconv.ParseUint(strings.Join(tp, ""), 64, 10)
        return uint(res), err
    }

    switch tp := value.(type) {
    case byte:
        return uint(tp), nil
    case rune:
        return uint(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type uint to %s", tp))
    }
}

func ConvertStringToInt8(value string) (uint8, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[uint8](tp), nil
    case string:
        res, err := strconv.ParseUint(tp, 8, 10)
        return uint8(res), err
    case []byte:
        res, err := strconv.ParseUint(string(tp), 8, 10)
        return uint8(res), err
    case []string:
        res, err := strconv.ParseUint(strings.Join(tp, ""), 8, 10)
        return uint8(res), err
    }

    switch tp := value.(type) {
    case byte:
        return uint8(tp), nil
    case rune:
        return uint8(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type uint8 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (uint16, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[uint16](tp), nil
    case string:
        res, err := strconv.ParseUint(tp, 16, 10)
        return uint16(res), err
    case []byte:
        res, err := strconv.ParseUint(string(tp), 16, 10)
        return uint16(res), err
    case []string:
        res, err := strconv.ParseUint(strings.Join(tp, ""), 16, 10)
        return uint16(res), err
    }

    switch tp := value.(type) {
    case byte:
        return uint16(tp), nil
    case rune:
        return uint16(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type uint16 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (uint32, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[uint32](tp), nil
    case string:
        res, err := strconv.ParseUint(tp, 32, 10)
        return uint32(res), err
    case []byte:
        res, err := strconv.ParseUint(string(tp), 32, 10)
        return uint32(res), err
    case []string:
        res, err := strconv.ParseUint(strings.Join(tp, ""), 32, 10)
        return uint32(res), err
    }

    switch tp := value.(type) {
    case byte:
        return uint32(tp), nil
    case rune:
        return uint32(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type uint32 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (uint64, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[uint64](tp), nil
    case string:
        return strconv.ParseUint(tp, 64, 10)
    case []byte:
        return strconv.ParseUint(string(tp), 64, 10)
    case []string:
        return strconv.ParseUint(strings.Join(tp, ""), 64, 10)
    }

    switch tp := value.(type) {
    case byte:
        return uint64(tp), nil
    case rune:
        return uint64(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type uint64 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (float32, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[float32](tp), nil
    case string:
        res, err := strconv.ParseFloat(tp, 32)
        return float32(res), err
    case []byte:
        res, err := strconv.ParseFloat(string(tp), 32)
        return float32(res), err
    case []string:
        res, err := strconv.ParseFloat(strings.Join(tp, ""), 32)
        return float32(res), err
    }

    switch tp := value.(type) {
    case byte:
        return float32(tp), nil
    case rune:
        return float32(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type float32 to %s", tp))
    }
}

func ConvertStringToInt8(value string) (float64, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[float64](tp), nil
    case string:
        return strconv.ParseFloat(tp, 64)
    case []byte:
        return strconv.ParseFloat(string(tp), 64)
    case []string:
        return strconv.ParseFloat(strings.Join(tp, ""), 64)
    }

    switch tp := value.(type) {
    case byte:
        return float64(tp), nil
    case rune:
        return float64(tp), nil
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type float64 to %s", tp))
    }
}

func ConAnyToComplex64(value any) (complex64, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[complex64](tp), nil
    case string:
        res, err := strconv.ParseComplex(tp, 64)
        return complex64(res), err
    case []byte:
        res, err := strconv.ParseComplex(string(tp), 64)
        return complex64(res), err
    case []string:
        res, err := strconv.ParseComplex(strings.Join(tp, ""), 64)
        return complex64(res), err
    }

    switch tp := value.(type) {
    case byte:
        res, err := strconv.ParseComplex(string(tp), 64)
        return complex64(res), err
    case rune:
        res, err := strconv.ParseComplex(string(tp), 64)
        return complex64(res), err
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type complex64 to %s", tp))
    }
}

func ConAnyToComplex128(value any) (complex128, error) {
    switch tp := value.(type) {
    case bool:
        if tp {
            return 1, nil
        } else {
            return 0, nil
        }
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint64, uint32, float32, float64, complex64, complex128:
        return conNumberTypeToNumberType[complex128](tp), nil
    case string:
        return strconv.ParseComplex(tp, 128)
    case []byte:
        return strconv.ParseComplex(string(tp), 128)
    case []string:
        return strconv.ParseComplex(strings.Join(tp, ""), 128)
    }

    switch tp := value.(type) {
    case byte:
        return strconv.ParseComplex(string(tp), 128)
    case rune:
        return strconv.ParseComplex(string(tp), 128)
    default:
        return 0, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type complex128 to %s", tp))
    }
}

func ConAnyToByte(value any) (byte, error) {
    res, err := ConAnyToInt32(value)
    if err != nil {
        err = errors.Wrap(ErrorNotParsing, fmt.Sprintf("with rune byte to %s", reflect.TypeOf(value).String()))
    }
    return byte(res), err
}

func ConAnyToRune(value any) (rune, error) {
    res, err := ConAnyToInt64(value)
    if err != nil {
        err = errors.Wrap(ErrorNotParsing, fmt.Sprintf("with rune rune to %s", reflect.TypeOf(value).String()))
    }
    return rune(res), err
}

func ConAnyToBytes(value any) ([]byte, error) {
    res, err := ConAnyToString(value)
    if err != nil {
        err = errors.Wrap(ErrorNotParsing, fmt.Sprintf("with rune []byte to %s", reflect.TypeOf(value).String()))
    }
    return []byte(res), err
}

func ConAnyToMapStrings(value any) (map[string]string, error) {
    return nil, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type map[string]string to %s", reflect.TypeOf(value).String()))
}

func ConAnyToMapString(value any) (map[string]any, error) {
    return nil, errors.Wrap(ErrorNotParsing, fmt.Sprintf("with type map[string]any to %s", reflect.TypeOf(value).String()))
}

func ConAnyToStrings(value any) ([]string, error) {
    res, err := ConAnyToString(value)
    if err != nil {
        err = errors.Wrap(ErrorNotParsing, fmt.Sprintf("with rune []string to %s", reflect.TypeOf(value).String()))
    }
    return []string{res}, err
}
