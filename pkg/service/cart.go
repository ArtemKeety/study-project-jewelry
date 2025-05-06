package service

import (
	"curs/jewelrymodel"
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

func (s *CartService) GetCart(userId int) ([]jewelrymodel.Cart, error) {
	return s.repo.GetCart(userId)
}

func (s *CartService) RemoveInCart(userId, cartId int) (int, error) {
	return s.repo.RemoveInCart(userId, cartId)
}

func (s *CartService) UpdateItemCart(request jewelrymodel.CartRequest) (int, error) {
	return s.repo.UpdateItemCart(request)
}

func (s *CartService) ClearCart(userId int) (int, error) {
	return s.repo.ClearCart(userId)
}
