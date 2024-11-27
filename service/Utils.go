package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

func isExported(name string) bool {
	w, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(w)
}

func ConvertString2Int8(value string) (iValue int8, err error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	return int8(intValue), nil
}

func ConvertString2Int16(value string) (iValue int16, err error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	return int16(intValue), nil
}
func ConvertString2Int32(value string) (iValue int32, err error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	return int32(intValue), nil
}
func ConvertString2Int(value string) (iValue int, err error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	return int(intValue), nil
}
func ConvertString2Int64(value string) (iValue int64, err error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	return int64(intValue), nil
}

func ConvertInterface2Int8(v interface{}) (int8, error) {
	if v == nil {
		return -1, errors.New("ConvertInterface2Int8 with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return s, nil
	case int16:
		var s int16 = v
		return int8(s), nil
	case int32:
		var s int32 = v
		return int8(s), nil
	case int64:
		var s int64 = v
		return int8(s), nil
	case int:
		var s int = v
		return int8(s), nil
	case float32:
		var s float32 = v
		return int8(s), nil
	case float64:
		var s float64 = v
		return int8(s), nil
	case string:
		return ConvertString2Int8(v)
	case bool:
		return 0, nil
	default:
		fmt.Println("ConvertInterface2Int8 with untype", v)
		return v.(int8), nil
	}
}

func ConvertInterface2Int16(v interface{}) (int16, error) {
	if v == nil {
		return -1, errors.New("ConvertInterface2Int16 with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return int16(s), nil
	case int16:
		var s int16 = v
		return s, nil
	case int32:
		var s int32 = v
		return int16(s), nil
	case int64:
		var s int64 = v
		return int16(s), nil
	case int:
		var s int = v
		return int16(s), nil
	case float32:
		var s float32 = v
		return int16(s), nil
	case float64:
		var s float64 = v
		return int16(s), nil
	case string:
		return ConvertString2Int16(v)
	case bool:
		return 0, nil
	default:
		fmt.Println("ConvertInterface2Int16 with untype", v)
		return v.(int16), nil
	}
}

func ConvertInterface2Int32(v interface{}) (int32, error) {
	if v == nil {
		return -1, errors.New("ConvertInterface2Int32 with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return int32(s), nil
	case int16:
		var s int16 = v
		return int32(s), nil
	case int32:
		var s int32 = v
		return s, nil
	case int64:
		var s int64 = v
		return int32(s), nil
	case float32:
		var s float32 = v
		return int32(s), nil
	case float64:
		var s float64 = v
		return int32(s), nil
	case int:
		var s int = v
		return int32(s), nil
	case string:
		return ConvertString2Int32(v)
	case bool:
		return 0, nil
	default:
		fmt.Println("ConvertInterface2Int32 with untype", v)
		return v.(int32), nil
	}
}

func ConvertInterface2Int64(v interface{}) (int64, error) {
	if v == nil {
		return -1, errors.New("ConvertInterface2Int64 with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return int64(s), nil
	case int16:
		var s int16 = v
		return int64(s), nil
	case int32:
		var s int32 = v
		return int64(s), nil
	case int64:
		var s int64 = v
		return int64(s), nil
	case float32:
		var s float32 = v
		return int64(s), nil
	case float64:
		var s float64 = v
		return int64(s), nil
	case int:
		var s int = v
		return int64(s), nil
	case string:
		return ConvertString2Int64(v)
	case bool:
		return 0, nil
	default:
		fmt.Println("ConvertInterface2Int64 with untype", v)
		return v.(int64), nil
	}
}

func ConvertInterface2Int(v interface{}) (int, error) {
	if v == nil {
		return -1, errors.New("ConvertInterface2Int with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return int(s), nil
	case int16:
		var s int16 = v
		return int(s), nil
	case int32:
		var s int32 = v
		return int(s), nil
	case int64:
		var s int64 = v
		return int(s), nil
	case int:
		var s int = v
		return s, nil
	case float32:
		var s float32 = v
		return int(s), nil
	case float64:
		var s float64 = v
		return int(s), nil
	case string:
		return ConvertString2Int(v)
	case bool:
		return 0, nil
	default:
		fmt.Println("ConvertInterface2Int with untype", v)
		return v.(int), nil
	}
}

func ConvertInterface2String(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("ConvertInterface2String with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		return strconv.Itoa(int(s)), nil
	case int16:
		var s int16 = v
		return strconv.Itoa(int(s)), nil
	case int32:
		var s int32 = v
		return strconv.Itoa(int(s)), nil
	case int64:
		var s int64 = v
		return strconv.Itoa(int(s)), nil
	case int:
		var s int = v
		return strconv.Itoa(int(s)), nil
	case float32:
		var s float32 = v
		return strconv.Itoa(int(s)), nil
	case float64:
		var s float64 = v
		return strconv.Itoa(int(s)), nil
	case string:
		var s string = v
		return string(s), nil
	case bool:
		var s bool = v
		return strconv.FormatBool(s), nil
	default:
		fmt.Println("ConvertInterface2String with untype", v)
		return v.(string), nil
	}
}

func ConvertInterface2Bool(v interface{}) (bool, error) {
	if v == nil {
		return false, errors.New("ConvertInterface2Bool with nil")
	}
	switch v := v.(type) {
	case int8:
		var s int8 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case int16:
		var s int16 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case int32:
		var s int32 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case int64:
		var s int64 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case int:
		var s int = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case float32:
		var s float32 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case float64:
		var s float64 = v
		if s != 0 {
			return true, nil
		}
		return false, nil
	case string:
		var s string = v
		if s != "" {
			return strconv.ParseBool(s)
		}
		return false, nil
	case bool:
		var s bool = v
		return s, nil
	default:
		fmt.Println("ConvertInterface2Bool with untype", v)
		return false, nil
	}
}

func ConvertInterface2Int32Array(v interface{}) ([]int32, error) {
	if v == nil {
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	}
	switch v := v.(type) {
	case int8:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case int16:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case int32:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case int64:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case float32:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case float64:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case int:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case string:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case bool:
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	case []int32:
		var s []int32 = v
		return s, nil
	default:
		fmt.Println("ConvertInterface2Int32Array with untype", v)
		return nil, errors.New("ConvertInterface2Int32Array error with untype")
	}
}

func ConvertTime2Int64(value string) (iValue int64, err error) {
	timeValue, err := time.Parse(defauleDateFormat, value)
	if err != nil {
		return
	}
	return timeValue.Unix(), nil
}

const (
	defauleDateFormat string = "2006-01-02 15:04:05"
)
