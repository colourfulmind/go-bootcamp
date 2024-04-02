package service

import (
	"cmd/app/app.go/internal/event"
)

var Prices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

type CandiesStruct struct{}

func NewCandies() *CandiesStruct {
	return &CandiesStruct{}
}

func (c *CandiesStruct) BuyCandy(candy string, money, count int) (int, error) {
	if Prices[candy] == 0 {
		return 0, event.WrongType
	} else if count <= 0 {
		return 0, event.WrongInput
	} else if money < Prices[candy]*count {
		return Prices[candy]*count - money, event.NotEnoughMoney
	}
	return money - Prices[candy]*count, nil
}
