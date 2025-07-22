package slices

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		want    int
		wantErr error
	}{
		{
			name:    "1",
			slice:   []int{1, 2, 3, 4, 5},
			want:    5,
			wantErr: nil,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "3",
			slice:   []int{999, 5, 46, 246, 15679, 154, 76649, 32, 45, 78, 66},
			want:    76649,
			wantErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Max(testCase.slice)
			assert.Equal(t, testCase.want, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.want, val)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		want    int
		wantErr error
	}{
		{
			name:    "1",
			slice:   []int{1, 2, 3, 4, 5},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "3",
			slice:   []int{999, 5, 46, 246, 15679, 154, 76649, 32, 45, 78, 66},
			want:    5,
			wantErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Min(testCase.slice)
			assert.Equal(t, testCase.want, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.want, val)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		want    int
		wantErr error
	}{
		{
			name:    "1",
			slice:   []int{1, 2, 3, 4, 5},
			want:    15,
			wantErr: nil,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "3",
			slice:   []int{999, 5, 46, 246, 15679, 154, 76649, 32, 45, 78, 66},
			want:    93999,
			wantErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Sum(testCase.slice)
			assert.Equal(t, testCase.want, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.want, val)
		})
	}
}

func TestAvg(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		want    int
		wantErr error
	}{
		{
			name:    "1",
			slice:   []int{1, 2, 3, 4, 5},
			want:    3,
			wantErr: nil,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "3",
			slice:   []int{999, 5, 46, 246, 15679, 154, 76649, 32, 45, 78, 66},
			want:    8545,
			wantErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Avg(testCase.slice)
			assert.Equal(t, testCase.want, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.want, val)
		})
	}
}

func TestCount(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		want    int
		wantErr error
	}{
		{
			name:    "1",
			slice:   []int{1, 2, 3, 4, 5, 5, 3, 2, 1, 1, 1},
			want:    5,
			wantErr: nil,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "3",
			slice:   []int{999, 5, 46, 246, 15679, 154, 76649, 32, 45, 78, 66, 999, 246, 154, 154, 999, 999},
			want:    11,
			wantErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Count_Distinct(testCase.slice)
			assert.Equal(t, testCase.want, err)
			if err != nil {
				return
			}
			assert.Equal(t, testCase.want, val)
		})
	}
}
