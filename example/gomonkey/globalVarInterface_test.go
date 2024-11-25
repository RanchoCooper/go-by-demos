package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

type Actions interface {
	Writes(string) bool
}

type InterfaceItem struct {
}

var GlobalInterfaceObject Actions

func (ex *InterfaceItem) Writes(key string) bool {
	// do something
	return true
}

func TestMockGlobalVarInterface(t *testing.T) {
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&InterfaceItem{}), "Writes", func(_ *InterfaceItem, key string) bool {
		return false
	})
	defer patches.Reset()

	GlobalInterfaceObject = &InterfaceItem{}
	result := GlobalInterfaceObject.Writes("any_key")
	assert.False(t, result)
}
