package tests

import (
	"go_day_00/internal/app/calculations"
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    float64
	}{
		{
			name:    "test1",
			numbers: []int{1, 2, 3, 4, 2, 1, 4, 3, 5},
			want:    2.78,
		},
		{
			name:    "test2",
			numbers: []int{45, 53, 42, 93, 92, 32, 11, 23, 38, 42, 57, 42, 45, 53, 42, 93, 92, 38, 42, 84, 34},
			want:    52.05,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculations.Statistics{}
			s.Mean(tt.numbers)
			get := math.Round(s.Average*100) / 100
			if get != tt.want {
				t.Errorf("Calculated mean = %v, real mean %v", get, tt.want)
			}
		})
	}
}
