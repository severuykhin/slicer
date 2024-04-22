package slicer

import (
	"reflect"
	"testing"
)

func TestMergeUnique(t *testing.T) {

	type testCase[T any] struct {
		key      string
		slice1   []T
		slice2   []T
		expected []T
	}

	type num int
	type char string

	cases := []testCase[any]{
		{
			key:      "1",
			slice1:   []any{1, 2, 3, 4},
			slice2:   []any{3, 4, 5, 6},
			expected: []any{1, 2, 3, 4, 5, 6},
		},
		{
			key:      "2",
			slice1:   []any{num(1), num(2), num(3), num(4)},
			slice2:   []any{num(3), num(4), num(5), num(6)},
			expected: []any{num(1), num(2), num(3), num(4), num(5), num(6)},
		},
		{
			key:      "3",
			slice1:   []any{"a", "b", "c", "d"},
			slice2:   []any{"c", "d", "e", "f"},
			expected: []any{"a", "b", "c", "d", "e", "f"},
		},
		{
			key:      "4",
			slice1:   []any{char("a"), char("b"), char("c"), char("d")},
			slice2:   []any{char("c"), char("d"), char("e"), char("f")},
			expected: []any{char("a"), char("b"), char("c"), char("d"), char("e"), char("f")},
		},
	}

	for _, tc := range cases {
		resultSlice := MergeUnique(tc.slice1, tc.slice2)
		if !reflect.DeepEqual(resultSlice, tc.expected) {
			t.Fatalf("Test case %s has result %+v but expected %+v", tc.key, resultSlice, tc.expected)
		}
	}

}
