package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

var globalVar = 10

// mock一个全局变量
func TestMockGlobalVar(t *testing.T) {
	patches := gomonkey.NewPatches()

	patches.ApplyGlobalVar(&globalVar, 100)
	defer patches.Reset()

	assert.Equal(t, 100, globalVar)
}

func doSomething(key string) bool {
	return false
}

// mock一个普通函数
// mock不生效, 是因为内联
// 需要禁用内联 go test -v -gcflags=-l ./gomonkey_test.go
func TestMockFunction(t *testing.T) {
	patches := gomonkey.NewPatches()

	patches.ApplyFunc(doSomething, func(string) bool {
		return true
	})
	defer patches.Reset()

	assert.True(t, doSomething("any"))
}

type MyObj struct {
}

func (m *MyObj) DoSomething(key string) bool {
	return false
}

// mock成员方法
func TestMockMethod(t *testing.T) {
	patches := gomonkey.NewPatches()

	patches.ApplyMethod(reflect.TypeOf(&MyObj{}), "DoSomething", func(*MyObj, string) bool {
		return true
	})
	defer patches.Reset()

	o := MyObj{}
	assert.True(t, o.DoSomething("any"))

	o2 := &MyObj{}
	assert.True(t, o2.DoSomething("any"))
}

type Action interface {
	Do(string) bool
}

type MyObj2 struct {
}

func (m *MyObj2) Do(key string) bool {
	return false
}

// mock一个interface
func TestMockInterface(t *testing.T) {
	patches := gomonkey.NewPatches()

	realObj := &MyObj2{}

	patches.ApplyMethod(reflect.TypeOf(realObj), "Do", func(*MyObj2, string) bool {
		return true
	})
	defer patches.Reset()

	assert.True(t, realObj.Do("any"))
}
