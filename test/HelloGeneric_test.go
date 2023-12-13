package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Print[T any](message T) T {
	fmt.Println(message)
	return message
}

func TestHelloWorldGenerics(t *testing.T) {
	message := "Hello World"
	result := Print[string](message)

	assert.Equal(t, message, result)
}
