package slices

import "testing"

func TestContains(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		value int
		want  bool
	}{
		{
			name:  "test1",
			slice: []int{1, 2, 3, 4, 5},
			value: 1,
			want:  true,
		},
		{
			name:  "test2",
			slice: []int{1, 2, 3, 4, 5},
			value: 99,
			want:  false,
		},
		{
			name:  "test3",
			value: 1,
			want:  false,
		},
		{
			name:  "test4",
			slice: []int{},
			value: 0,
			want:  false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := Contains(testCase.slice, testCase.value); got != testCase.want {
				t.Errorf("Contains(%v, %v) = %v; want %v", testCase.slice, testCase.value, got, testCase.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		value []int
		want  bool
	}{
		{
			name:  "test1",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{1, 2, 3, 4, 5},
			want:  true,
		},
		{
			name:  "test2",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{1, 2, 3},
			want:  true,
		},
		{
			name:  "test3",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{1, 7},
			want:  false,
		},
		{
			name:  "test4",
			slice: []int{},
			value: []int{1},
			want:  false,
		},
		{
			name:  "test5",
			value: []int{9},
			want:  false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := ContainsAll(testCase.slice, testCase.value); got != testCase.want {
				t.Errorf("ContainsAll(%v, %v) = %v; want %v", testCase.slice, testCase.value, got, testCase.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		value []int
		want  bool
	}{
		{
			name:  "test1",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{1, 2, 3},
			want:  true,
		},
		{
			name:  "test2",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{1, 7},
			want:  true,
		},
		{
			name:  "test3",
			slice: []int{1, 2, 3, 4, 5},
			value: []int{9, 7},
			want:  false,
		},
		{
			name:  "test4",
			slice: []int{},
			value: []int{1},
			want:  false,
		},
		{
			name:  "test5",
			value: []int{9},
			want:  false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := ContainsAny(testCase.slice, testCase.value); got != testCase.want {
				t.Errorf("ContainsAny(%v, %v) = %v; want %v", testCase.slice, testCase.value, got, testCase.want)
			}
		})
	}
}
