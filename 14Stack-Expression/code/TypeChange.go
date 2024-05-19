package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//给我变量类型，判断处理
func M2int64(i interface{}) int64 {
	if i == nil {
		panic(errors.New("类型为空，无法判断"))
	}
	switch t := i.(type) {
	case int64:
		return t
	case int32:
		return int64(t)
	case int16:
		return int64(t)
	case int8:
		return int64(t)
	case int:
		return int64(t)
	case float64:
		return int64(t)
	case float32:
		return int64(t)
	case byte:
		return int64(t)
	default:
		//return int64(t)
		return Stoi64(M2string(i))
	}

}
func M2string(i interface{}) string {
	if i == nil {
		panic(errors.New("类型为空，无法判断"))
	}
	switch t := i.(type) {
	case string:
		return t
	default:
		return fmt.Sprintf("%v", i)
	}
}

//找到并且没有默认值，报错
func Stoi64(v string, def ...int64) int64 {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(errors.New("64位数转换失败"))
	}
}

func M2float64(i interface{}) float64 {
	if i == nil {
		panic(errors.New("类型为空，无法判断"))
	}
	switch t := i.(type) {
	case int64:
		return float64(t)
	case int32:
		return float64(t)
	case int16:
		return float64(t)
	case int8:
		return float64(t)
	case int:
		return float64(t)
	case float64:
		return float64(t)
	case float32:
		return float64(t)
	case byte:
		return float64(t)
	default:
		//return int64(t)
		return Stof64(M2string(i))
	}

}

//找到并且没有默认值，报错
func Stof64(v string, def ...float64) float64 {
	if n, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(errors.New("float 64位数转换失败"))
	}
}
