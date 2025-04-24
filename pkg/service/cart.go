package service

import (
	"curs/pkg/repository"
	"errors"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddInCart(productId, userId int) (int, error) {
	if cartId, _ := s.repo.CheckInCart(productId, userId); cartId > 0 {
		return -1, errors.New("item is already in cart")
	}
	return s.repo.AddInCart(productId, userId)
}

func (s *CartService) CheckInCart(productId, userId int) (int, error) {
	return s.repo.CheckInCart(userId, productId)
}
