package service

type Candies interface {
	BuyCandy(candy string, money, count int) (int, error)
}

type Service struct {
	Candies
}

func New() *Service {
	return &Service{
		Candies: NewCandies(),
	}
}
