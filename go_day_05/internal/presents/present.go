package presents

type Present struct {
	Value int
	Size  int
}

func GetNCoolestPresents(presents PresentHeap, n int) PresentHeap {
	if n > len(presents) {
		n = len(presents)
	}
	return SortPresents(presents)[:n]
}

func SortPresents(presents PresentHeap) PresentHeap {
	if presents.Len() < 2 {
		return presents
	}
	pivot := presents[0]
	var l, g PresentHeap
	for i := 1; i < len(presents); i++ {
		if presents.Less(0, i) {
			l = append(l, presents[i])
		} else {
			g = append(g, presents[i])
		}
	}
	res := append(SortPresents(l), pivot)
	res = append(res, SortPresents(g)...)
	return res
}
