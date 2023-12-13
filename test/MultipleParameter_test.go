package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AddELement[T any](e1, e2 T) []T {
	var Elemets []T
	Elemets = append(Elemets, e1, e2)
	return Elemets
}

func TestMultipleParameter(t *testing.T) {
	elements := AddELement[int](1, 3)
	assert.Equal(t, elements[0], 1)
	assert.Equal(t, elements[1], 3)
}
