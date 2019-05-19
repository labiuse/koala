package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		actual, expected []int
	}{
		{
			actual:   FindRepeated([]int{1, 3, 3, 5, 5, 6, 6, 5, 3, 3}, 2),
			expected: []int{3, 5},
		},
		{
			actual:   FindRepeated([]int{}, 2),
			expected: []int{},
		},
		{
			actual:   FindRepeated([]int{1, 2}, 3),
			expected: []int{1, 2},
		},
		{
			actual:   FindRepeated([]int{1, 2}, 2),
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.actual, test.expected) {
			t.Errorf("wrong result, expected %v, got %v", test.expected, test.actual)
		}
	}

}
