package service

import (
	"order-payment-system/internal/model"
	"order-payment-system/internal/repository"
)

type PaymentService struct {
	orderRepo *repository.OrderRepo
	userRepo  *repository.UserRepo
}

func NewPaymentService(orderRepo *repository.OrderRepo, userRepo *repository.UserRepo) *PaymentService {
	return &PaymentService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (p *PaymentService) GetOrder(orderNo string) (*model.Order, error) {
	return p.orderRepo.GetOrderByOrderNo(orderNo)
}

func (p *PaymentService) Settling(order *model.Order) error {
	return p.userRepo.Deduct(order.UserID, order.TotalPrice)
}
