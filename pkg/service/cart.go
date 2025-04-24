package service

import "curs/pkg/repository"

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddInCart(productId, userId int) (int, error) {
	return s.repo.AddInCart(productId, userId)
}
