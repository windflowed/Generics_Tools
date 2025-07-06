package slices

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		value     int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			value:     10,
			index:     0,
			wantSlice: []int{10, 123, 100},
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			value:     178,
			index:     1,
			wantSlice: []int{123, 178, 124, 125},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100},
			value:   109,
			index:   12,
			wantErr: errors.New("index out of range"),
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 100},
			value:   321,
			index:   -1,
			wantErr: errors.New("index out of range"),
		},
		{
			name:      "index last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			value:     999,
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102, 102, 999},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sli, err := Add(testCase.slice, testCase.value, testCase.index)
			assert.Equal(t, testCase.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.wantSlice, sli)
		})
	}
}
