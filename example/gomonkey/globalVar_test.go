package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type Object struct {
	data map[string]string
	name string
}

var GlobalObject Object

func (ex *Object) Get(apiKey string) (string, error) {
	secret, ok := ex.data[apiKey]
	if !ok {
		return "", errors.Errorf("not found api-secret by api-key:%s", apiKey)
	}

	return secret, nil
}

func TestMockExchangeAuthGet(t *testing.T) {
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&Object{}), "Get", func(_ *Object, apiKey string) (string, error) {
		return "mock-secret", nil
	})
	defer patches.Reset()

	item, err := GlobalObject.Get("any_key")
	assert.NoError(t, err)
	assert.Equal(t, "mock-secret", item)
}
