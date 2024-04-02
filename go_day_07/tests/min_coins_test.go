package tests

import (
	"coins/internal/coins"
	"reflect"
	"testing"
)

func TestMinCoinsSuccess(t *testing.T) {
	tests := []struct {
		name  string
		val   int
		coins []int
		want  []int
	}{
		{
			name:  "test1",
			val:   13,
			coins: []int{1, 5, 10},
			want:  []int{10, 1, 1, 1},
		},
		{
			name:  "test2",
			val:   13,
			coins: []int{1, 1, 1, 5, 10},
			want:  []int{10, 1, 1, 1},
		},
		{
			name:  "test3",
			val:   13,
			coins: []int{},
			want:  []int{},
		},
		{
			name:  "test4",
			val:   -13,
			coins: []int{},
			want:  []int{},
		},
		{
			name:  "test5",
			val:   -13,
			coins: nil,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := coins.MinCoins(tt.val, tt.coins)
			got2 := coins.MinCoins2(tt.val, tt.coins)
			if !reflect.DeepEqual(got1, tt.want) && !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("expected: %v, got: %v && %v", tt.want, got1, got2)
			}
		})
	}
}

func TestMinCoinsFail(t *testing.T) {
	tests := []struct {
		name  string
		val   int
		coins []int
		want  []int
	}{
		{
			name:  "test1",
			val:   13,
			coins: []int{1, 6, 10},
			want:  []int{6, 6, 1},
		},
		{
			name:  "test2",
			val:   26,
			coins: []int{3, 6, 10, 15},
			want:  []int{10, 10, 6},
		},
		{
			name:  "test3",
			val:   26,
			coins: []int{3, 6, 10, 15, 25},
			want:  []int{10, 10, 6},
		},
		{
			name:  "test4",
			val:   13,
			coins: []int{10, 6, 1},
			want:  []int{6, 6, 1},
		},
		{
			name:  "test5",
			val:   13,
			coins: []int{5, 7, 12},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coins.MinCoins(tt.val, tt.coins)
			got2 := coins.MinCoins2(tt.val, tt.coins)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("unexpected result")

			}
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("got: %v, want: %v", got2, tt.want)
			}
		})
	}
}
