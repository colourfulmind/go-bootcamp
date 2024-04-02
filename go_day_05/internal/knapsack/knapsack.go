package knapsack

import (
	"main/internal/presents"
	"math"
)

func GrabPresents(p presents.PresentHeap, W int) presents.PresentHeap {
	if W <= 0 {
		return presents.PresentHeap{}
	}
	n := p.Len()
	m := make([][]int, n+1)
	for i := range m {
		m[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if p[i-1].Size > j {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = int(math.Max(float64(m[i-1][j]), float64(m[i-1][j-p[i-1].Size]+p[i-1].Value)))
			}
		}
	}

	return ReceivePresents(m, n, W, p)
}

func ReceivePresents(m [][]int, n, W int, p presents.PresentHeap) presents.PresentHeap {
	var res presents.PresentHeap
	if m[n][W] == 0 {
		return presents.PresentHeap{}
	}
	if m[n-1][W] == m[n][W] {
		res = append(res, ReceivePresents(m, n-1, W, p)...)
	} else {
		res = append(ReceivePresents(m, n-1, W-p[n-1].Size, p), p[n-1])
	}
	return res
}
