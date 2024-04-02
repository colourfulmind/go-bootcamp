package tests

import (
	"container/heap"
	"main/internal/presents"
	"testing"
)

func TestPresents(t *testing.T) {
	tests := []struct {
		name     string
		p        *presents.PresentHeap
		n        int
		expected presents.PresentHeap
	}{
		{
			name: "test1",
			p: &presents.PresentHeap{presents.Present{Value: 5, Size: 1},
				presents.Present{Value: 4, Size: 5},
				presents.Present{Value: 3, Size: 2},
				presents.Present{Value: 5, Size: 2},
			},
			n: 2,
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
			n: 1,
			expected: presents.PresentHeap{
				{5, 10},
			},
		},
		{
			name:     "test3",
			p:        &presents.PresentHeap{},
			n:        1,
			expected: presents.PresentHeap{},
		},
	}

	for _, tt := range tests {
		heap.Init(tt.p)
		res := presents.GetNCoolestPresents(*tt.p, tt.n)
		if len(res) != len(tt.expected) {
			t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
		}
		for i := range res {
			if res[i] != tt.expected[i] {
				t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
			}
		}
	}
}
