package tests

import (
	"container/heap"
	"main/internal/knapsack"
	"main/internal/presents"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		p        *presents.PresentHeap
		W        int
		expected presents.PresentHeap
	}{
		{
			name: "test1",
			p: &presents.PresentHeap{presents.Present{Value: 5, Size: 1},
				presents.Present{Value: 4, Size: 5},
				presents.Present{Value: 3, Size: 2},
				presents.Present{Value: 5, Size: 2},
			},
			W: 3,
			expected: presents.PresentHeap{
				{5, 1},
				{5, 2},
			},
		},
		{
			name: "test2",
			p: &presents.PresentHeap{presents.Present{Value: 3, Size: 5},
				presents.Present{Value: 5, Size: 10},
				presents.Present{Value: 4, Size: 6},
				presents.Present{Value: 2, Size: 5},
			},
			W: 14,
			expected: presents.PresentHeap{
				{3, 5},
				{4, 6},
			},
		},
		{
			name:     "test3",
			p:        &presents.PresentHeap{},
			W:        14,
			expected: presents.PresentHeap{},
		},
		{
			name:     "test4",
			p:        &presents.PresentHeap{},
			W:        -14,
			expected: presents.PresentHeap{},
		},
	}
	for _, tt := range tests {
		heap.Init(tt.p)
		res := knapsack.GrabPresents(*tt.p, tt.W)
		if len(res) != len(tt.expected) {
			t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
		}
		for i := 0; i < len(res); i++ {
			if res[i] != tt.expected[i] {
				t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
			}
		}
	}
}
