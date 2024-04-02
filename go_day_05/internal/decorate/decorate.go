package decorate

import (
	"main/internal/tree"
)

func UnrollGarland(t *tree.Node) []bool {
	var res = []bool{t.HasToy}
	return append(res, CreateQueue([]*tree.Node{t}, true)...)
}

func CreateQueue(nodes []*tree.Node, level bool) []bool {
	if len(nodes) == 0 {
		return []bool{}
	}

	var unrolled []bool
	q := make([]*tree.Node, 0)
	l := len(nodes)

	nodes = ReverseQueue(nodes)

	for i := 0; i < l; i++ {
		if level {
			if nodes[i].Left != nil {
				q = append(q, nodes[i].Left)
				unrolled = append(unrolled, nodes[i].Left.HasToy)
			}
			if nodes[i].Right != nil {
				q = append(q, nodes[i].Right)
				unrolled = append(unrolled, nodes[i].Right.HasToy)
			}
		} else {
			if nodes[i].Right != nil {
				q = append(q, nodes[i].Right)
				unrolled = append(unrolled, nodes[i].Right.HasToy)
			}
			if nodes[i].Left != nil {
				q = append(q, nodes[i].Left)
				unrolled = append(unrolled, nodes[i].Left.HasToy)
			}
		}
	}

	unrolled = append(unrolled, CreateQueue(q, !level)...)
	return unrolled
}

func ReverseQueue(nodes []*tree.Node) []*tree.Node {
	var reversed = make([]*tree.Node, 0)
	for i := len(nodes) - 1; i >= 0; i-- {
		reversed = append(reversed, nodes[i])
	}
	return reversed
}
