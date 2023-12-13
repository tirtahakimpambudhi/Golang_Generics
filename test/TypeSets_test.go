package test

import (
	"testing"
)

//	func sortAscending(slices []int) []int {
//		sort.Ints(slices)
//		return slices
//	}
//
//	func sortDescending(slices []int) []int {
//		sort.Sort(sort.Reverse(sort.IntSlice(slices)))
//		return slices
//	}
type Age int
type Ordered interface {
	~int | int8 | int16 | int32 | int64 | float32 | ~float64
}

func SortAscending[T Ordered](slice []T) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

func SortDescending[T Ordered](slice []T) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

func TestSort(t *testing.T) {
	slices := []int{9, 7, 5, 4, 6, 3, 1, 2, 8}
	ages := []Age{1, 3, 4, 5, 9, 8, 7, 6, 2}
	sliceFloat := []float64{1.0, 9.0, 8.0, 2.0, 4.0, 3.0, 5.0, 6.0, 7.0, 10.0}

	t.Log("Integer")
	//Sort Integer
	t.Log(SortAscending[int](slices))
	t.Log(SortDescending[int](slices))
	t.Log("Float")
	//Sort FLoat
	t.Log(SortAscending[float64](sliceFloat))
	t.Log(SortDescending[float64](sliceFloat))
	//Type Approximation
	t.Log("Age")
	t.Log(SortAscending(ages))
	t.Log(SortDescending(ages))
}
