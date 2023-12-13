package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func ChangeValue[T any](params GetterSetter[T], value T) T {
	params.SetValue(value)
	return params.GetValue()
}

func TestGenericsInterface(t *testing.T) {
	data := MyData[string]{Value: "Hello World"}
	ChangeValue(&data, "Hello Generics Interface")
	assert.Equal(t, "Hello Generics Interface", data.Value)
}
