package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

// 结构体转map[string]interface{}
// origin https://mp.weixin.qq.com/s/Q_CqHQxDZMpjLNJTp5kfQg

type UserInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// ToMap convert a single struct to map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer, but got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if tagValue := field.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// ToMap2 convert a embeded struct to map[string]interface{}
func ToMap2(in interface{}, tag string) (map[string]interface{}, error) {

}


func main() {
	userInfo := UserInfo{
		Name: "rancho",
		Age:  25,
	}

	// 序列化struct
	b, _ := json.Marshal(&userInfo)

	var m map[string]interface{}
	// 反序列化成map
	_ = json.Unmarshal(b, &m)

	// 输出验证
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// 输出类型
	for k, v := range m{
		// 不难发现，age是float64，而不是原先的int
		fmt.Printf("key: %v, value: %v, type: %T\n", k, v, v)
	}


	// 使用ToMap
	m2, _ := ToMap(&userInfo, "json")
	for k, v := range m2{
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}
