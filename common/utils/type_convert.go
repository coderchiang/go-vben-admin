package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

//StrToUInt string 转int
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
// string到int64
func StringToint64(a string)(int64, error) {
	return strconv.ParseInt(a, 10, 64)
}
// string to float32、float64
func StringToFloat64(a string) (float64, error) {
	return strconv.ParseFloat(a, 64)
}
// int到string
func IntToString(a int) string {
	return strconv.Itoa(a)
}
// int64到string
func Int64ToString(a int64) string {
	return strconv.FormatInt(a, 10)
}
// float32/64 to string
func Float32ToString(a float32) string {
	return strconv.FormatFloat(3.1415, 'E', -1, 32)
}

func Float64ToString(a float64) string {
	return strconv.FormatFloat(3.1415, 'E', -1, 64)
}

// int和int64
func IntToInt64(a int) int64 {
	return int64(a)
}

// int64 to int
func Int64ToInt(a int64) int {
	return int(a)
}

func StringToStructOrMapOrArray(a string, b *interface{}) error {
	return json.Unmarshal([]byte(a), &b)
}

// 将结构体解析为字符串
func StructOrMapOrArrayToString(a interface{}) (string, error) {
	b, err := json.Marshal(a)
	return string(b), err
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}



