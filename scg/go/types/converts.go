package types

import (
    "fmt"
    "github.com/pkg/errors"
    "strconv"
    "strings"
)

const (
    tenBase     = 10
    baseSize8   = 8
    baseSize16  = 16
    baseSize32  = 32
    baseSize64  = 64
    baseSize128 = 128
)

func ConvertStringToBool(value string) (bool, error) {
    res, err := strconv.ParseBool(value)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to bool", value))
}

func ConvertStringToInt(value string) (int, error) {
    res, err := strconv.Atoi(value)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to int", value))
}

func ConvertStringToInt8(value string) (int8, error) {
    res, err := strconv.ParseInt(value, tenBase, baseSize8)

    return int8(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to int8", value))
}

func ConvertStringToInt16(value string) (int16, error) {
    res, err := strconv.ParseInt(value, tenBase, baseSize16)

    return int16(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to int16", value))
}

func ConvertStringToInt32(value string) (int32, error) {
    res, err := strconv.ParseInt(value, tenBase, baseSize32)

    return int32(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to int32", value))
}

func ConvertStringToInt64(value string) (int64, error) {
    res, err := strconv.ParseInt(value, tenBase, baseSize64)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to int64", value))
}

func ConvertStringToUint(value string) (uint, error) {
    res, err := strconv.ParseUint(value, tenBase, baseSize32)

    return uint(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to uint", value))
}

func ConvertStringToUint8(value string) (uint8, error) {
    res, err := strconv.ParseUint(value, tenBase, baseSize8)

    return uint8(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to uint8", value))
}

func ConvertStringToUint16(value string) (uint16, error) {
    res, err := strconv.ParseUint(value, tenBase, baseSize16)

    return uint16(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to uint16", value))
}

func ConvertStringToUint32(value string) (uint32, error) {
    res, err := strconv.ParseUint(value, tenBase, baseSize32)

    return uint32(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to uint32", value))
}

func ConvertStringToUint64(value string) (uint64, error) {
    res, err := strconv.ParseUint(value, tenBase, baseSize64)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to uint64", value))
}

func ConvertStringToFloat32(value string) (float32, error) {
    res, err := strconv.ParseFloat(value, baseSize32)

    return float32(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to float32", value))
}

func ConvertStringToFloat64(value string) (float64, error) {
    res, err := strconv.ParseFloat(value, baseSize64)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to float64", value))
}

func ConvertStringToComplex64(value string) (complex64, error) {
    res, err := strconv.ParseComplex(value, baseSize64)

    return complex64(res), errors.Wrap(err, fmt.Sprintf("with try convert from %s to complex64", value))
}

func ConvertStringToComplex128(value string) (complex128, error) {
    res, err := strconv.ParseComplex(value, baseSize128)

    return res, errors.Wrap(err, fmt.Sprintf("with try convert from %s to complex128", value))
}

func ConvertStringToByte(value string) (byte, error) {
    res, err := ConvertStringToInt8(value)
    if err != nil {
        err = errors.Wrap(errors.Unwrap(err), fmt.Sprintf("with try convert from %s to byte", value))
    }

    return byte(res), err
}

func ConvertStringToRune(value string) (rune, error) {
    rs := []rune(value)
    if len(rs) == 0 {
        return 0, errors.Wrap(errors.New("empty string"), fmt.Sprintf("with try convert from %s to rune", value))
    }

    return rs[0], nil
}

func ConvertStringToBytes(value string) ([]byte, error) {
    return []byte(value), nil
}

func ConvertStringToStrings(value string) ([]string, error) {
    return strings.Split(value, " "), nil
}
