package tests

import (
	"main/internal/balance"
	"main/internal/tree"
	"testing"
)

func TestBalanced(t *testing.T) {
	tests := []struct {
		name     string
		t        *tree.Node
		expected bool
	}{
		{
			name:     "test1",
			expected: true,
		},
		{
			name:     "test2",
			expected: true,
		},
	}

	tests[0].t = tree.New(true)
	{
		tests[0].t.Left = tree.New(true)
		tests[0].t.Left.Left = tree.New(true)
		tests[0].t.Left.Right = tree.New(false)

		tests[0].t.Right = tree.New(false)
		tests[0].t.Right.Left = tree.New(true)
		tests[0].t.Right.Right = tree.New(true)
	}

	tests[1].t = tree.New(false)
	{
		tests[1].t.Left = tree.New(false)
		tests[1].t.Left.Left = tree.New(false)
		tests[1].t.Left.Right = tree.New(true)

		tests[1].t.Right = tree.New(true)
	}

	for _, tt := range tests {
		res := balance.AreToysBalanced(tt.t)
		if res != tt.expected {
			t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
		}
	}
}

func TestNotBalanced(t *testing.T) {
	tests := []struct {
		name     string
		t        *tree.Node
		expected bool
	}{
		{
			name:     "test1",
			expected: false,
		},
		{
			name:     "test2",
			expected: false,
		},
	}

	tests[0].t = tree.New(true)
	{
		tests[0].t.Left = tree.New(true)
		tests[0].t.Right = tree.New(false)
	}

	tests[1].t = tree.New(false)
	{
		tests[1].t.Left = tree.New(true)
		tests[1].t.Left.Right = tree.New(true)

		tests[1].t.Right = tree.New(false)
		tests[1].t.Right.Right = tree.New(true)
	}

	for _, tt := range tests {
		res := balance.AreToysBalanced(tt.t)
		if res != tt.expected {
			t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
		}
	}
}
