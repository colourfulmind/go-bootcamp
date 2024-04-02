package tests

import (
	"main/internal/decorate"
	"main/internal/tree"
	"testing"
)

func TestDecorate(t *testing.T) {
	tests := []struct {
		name     string
		t        *tree.Node
		expected []bool
	}{
		{
			name:     "test1",
			expected: []bool{true, true, false, true, true, false, true},
		},
		{
			name:     "test2",
			expected: []bool{true, true, false, true, true, false, true, false, true, true, false},
		},
		{
			name:     "test3",
			expected: []bool{true, true, false, true, true, false, true, false, true, true, false, false, true, false},
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

	tests[1].t = tree.New(true)
	{
		tests[1].t.Left = tree.New(true)

		tests[1].t.Left.Left = tree.New(true)
		tests[1].t.Left.Left.Left = tree.New(false)
		tests[1].t.Left.Left.Right = tree.New(true)

		tests[1].t.Left.Right = tree.New(false)
		tests[1].t.Left.Right.Right = tree.New(true)

		tests[1].t.Right = tree.New(false)
		tests[1].t.Right.Left = tree.New(true)

		tests[1].t.Right.Right = tree.New(true)
		tests[1].t.Right.Right.Left = tree.New(false)
	}

	tests[2].t = tree.New(true)
	{
		tests[2].t.Left = tree.New(true)

		tests[2].t.Left.Left = tree.New(true)
		tests[2].t.Left.Left.Left = tree.New(false)
		tests[2].t.Left.Left.Right = tree.New(true)

		tests[2].t.Left.Right = tree.New(false)
		tests[2].t.Left.Right.Right = tree.New(true)
		tests[2].t.Left.Right.Right.Left = tree.New(false)

		tests[2].t.Right = tree.New(false)
		tests[2].t.Right.Left = tree.New(true)

		tests[2].t.Right.Right = tree.New(true)
		tests[2].t.Right.Right.Left = tree.New(false)
		tests[2].t.Right.Right.Left.Left = tree.New(true)
		tests[2].t.Right.Right.Left.Right = tree.New(false)
	}

	for _, tt := range tests {
		res := decorate.UnrollGarland(tt.t)
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
