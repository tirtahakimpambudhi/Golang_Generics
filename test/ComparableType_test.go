package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Equal[T comparable](el1, el2 T) bool {
	return el1 == el2
}

func NotEqual[T comparable](el1, el2 T) bool {
	return el1 != el2
}

func TestComparableType(t *testing.T) {
	eq := Equal[string]("aku", "aku")
	neq := NotEqual[int](100, 200)

	assert.Equal(t, true, eq)
	assert.Equal(t, true, neq)
}
