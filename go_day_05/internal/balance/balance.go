package balance

import "main/internal/tree"

func AreToysBalanced(t *tree.Node) bool {
	return CountToys(t.Left) == CountToys(t.Right)
}

func CountToys(t *tree.Node) int {
	if t == nil {
		return 0
	}
	if t.HasToy {
		return CountToys(t.Left) + CountToys(t.Right) + 1
	}
	return CountToys(t.Left) + CountToys(t.Right)
}
