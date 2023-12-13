package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Bags[T comparable] []T

func (b *Bags[T]) AddItem(item T) {
	*b = append(*b, item)
}

func (b *Bags[T]) FindItem(item T) (T, error) {
	for _, t := range *b {
		if t == item {
			return t, nil
		}
	}
	return item, fmt.Errorf("'%v' Item Not Found", item)
}

func (b *Bags[T]) DeleteItem(item T) error {
	_, err := b.FindItem(item)
	if err != nil {
		return err
	} else {
		newBags := []T{}
		for _, t := range *b {
			if t != item {
				newBags = append(newBags, t)
			}
		}
		*b = newBags
		return nil
	}
}

func (b *Bags[T]) EditItem(oldItem, newItem T) error {
	for i, t := range *b {
		if t == oldItem {
			(*b)[i] = newItem
			return nil
		}
	}
	return fmt.Errorf("'%v' Item Not Found", oldItem)
}
func TestGenericsTypesAndMeth(t *testing.T) {
	mybags := Bags[int]{1, 2, 3, 4, 5, 6, 7}
	//ADD ITEM
	mybags.AddItem(8)
	assert.Equal(t, mybags[len(mybags)-1], 8)
	t.Log("ADD ITEM PASSED")
	//Remove Item Success
	err := mybags.DeleteItem(8)
	assert.Equal(t, nil, err)
	//Remove Item Failed
	err = mybags.DeleteItem(9)
	assert.NotEqual(t, nil, err)
	t.Log("REMOVE ITEM PASSED")
	//Find Item Success
	item, err := mybags.FindItem(7)
	assert.Equal(t, nil, err)
	assert.Equal(t, 7, item)
	//Find Item Failed
	_, err = mybags.FindItem(9)
	assert.NotEqual(t, nil, err)
	t.Log("FIND ITEM PASSED")

	//Edit Item Success
	err = mybags.EditItem(6, 7)
	assert.Equal(t, nil, err)
	//Edit Item Failed
	err = mybags.EditItem(9, 10)
	assert.NotEqual(t, nil, err)

	t.Log("EDIT ITEM PASSED")
}
