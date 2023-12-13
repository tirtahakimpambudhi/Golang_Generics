package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	FirstData  T
	SecondData T
}

func (d *Data[T]) Change(firstData, secondData T) {
	d.FirstData = firstData
	d.SecondData = secondData
}

func TestGenericsStruct(t *testing.T) {
	message := Data[string]{
		FirstData:  "Hello",
		SecondData: "World",
	}
	message.Change("Hello", "Generics")

	assert.Equal(t, "Hello", message.FirstData)
	assert.Equal(t, "Generics", message.SecondData)
}
