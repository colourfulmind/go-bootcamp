package tests

import (
	"coins/internal/coins"
	"math/rand"
	"testing"
	"time"
)

var v, c = 10101, GenerateData(100)

func BenchmarkMinCoins(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = coins.MinCoins(v, c)
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = coins.MinCoins2(v, c)
	}
}

func GenerateData(x int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var coinsList = make([]int, x)

	for i := 0; i < x; i++ {
		coinsList[i] = rand.Int()
	}

	return coinsList
}
