package tests

import (
	"go_day_00/internal/app/calculations"
	"testing"
)

func TestMode(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "test1",
			numbers: []int{1, 2, 3, 4, 2, 1, 4, 3, 4, 5, 3, 4},
			want:    4,
		},
		{
			name:    "test2",
			numbers: []int{1, 2, 3, 4, 2, 1, 4, 3, 4, 5, 3},
			want:    3,
		},
		{
			name:    "test3",
			numbers: []int{45, 53, 42, 93, 92, 32, 11, 23, 38, 42, 57, 42, 45, 53, 42, 93, 92, 38, 42, 84, 34},
			want:    42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculations.Statistics{}
			s.Mode(tt.numbers)
			get := s.MostCommon
			if get != tt.want {
				t.Errorf("Calculated mode = %v, real mode = %v", get, tt.want)
			}
		})
	}
}
