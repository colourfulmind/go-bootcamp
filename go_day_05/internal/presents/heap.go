package presents

type PresentHeap []Present

func (p PresentHeap) Len() int {
	return len(p)
}

func (p PresentHeap) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Size > p[j].Size
	}
	return p[i].Value < p[j].Value
}

func (p PresentHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PresentHeap) Push(x any) {
	*p = append(*p, x.(Present))
}

func (p *PresentHeap) Pop() any {
	old := *p
	n := old.Len()
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
