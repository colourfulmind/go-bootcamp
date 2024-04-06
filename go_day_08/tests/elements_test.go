package tests

import (
	"errors"
	"fmt"
	"main/internal/elements"
	"testing"
)

func TestElementsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		elements []int
		want     int
	}{
		{
			name:     "test1",
			index:    3,
			elements: []int{0, 1, 2, 3, 4, 5},
			want:     3,
		},
		{
			name:     "test2",
			index:    6,
			elements: []int{213, 43, 231, 9432, 832, 3883, 22},
			want:     22,
		},
		{
			name:     "test3",
			index:    0,
			elements: []int{0, 1, 2, 3, 4, 5},
			want:     0,
		},
		{
			name:  "test4",
			index: 13,
			elements: []int{213, 4113, 231, 945432, 832, 3883, 22213,
				43, 231, 9432, 832, 3883, 22213, 43, 2431, 941732, 832,
				38583, 2052213, 4333, 2315, 9432, 832, 389983, 22, 9834},
			want: 43,
		},
		{
			name:     "test5",
			index:    0,
			elements: []int{0},
			want:     0,
		},
		{
			name:     "test6",
			index:    0,
			elements: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := elements.GetElement(tt.elements, tt.index)
			if got != tt.want || err != nil {
				t.Errorf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestElementsFail(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		elements []int
		err      error
	}{
		{
			name:     "test1",
			index:    3,
			elements: []int{},
			err:      errors.New(fmt.Sprintf("invalid data provided: %v", []int{})),
		},
		{
			name:     "test2",
			index:    3,
			elements: nil,
			err:      errors.New(fmt.Sprintf("invalid data provided: %v", []int{})),
		},
		{
			name:     "test3",
			index:    10,
			elements: []int{0, 1, 2, 3, 4, 5},
			err:      errors.New(fmt.Sprintf("index out of range: %d\n", 10)),
		},
		{
			name:     "test4",
			index:    -3,
			elements: []int{0, 1, 2, 3, 4, 5},
			err:      errors.New(fmt.Sprintf("index out of range: %d\n", -3)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := elements.GetElement(tt.elements, tt.index)
			if err == nil {
				t.Errorf("expected: %v, got: %v", tt.err, err)
			}
		})
	}
}
