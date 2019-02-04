package main

type priceService interface {
	GetPrice(orderID int) float64
}

// online part of my service
type onlinePriceService struct{}

func (ops onlinePriceService) GetPrice(orderID int) float64 {
	return 0
}

// test part of my service for QA
type testPriceService struct{}

func (ops testPriceService) GetPrice(orderID int) float64 {
	return 0.11
}
