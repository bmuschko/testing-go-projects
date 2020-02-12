package main

import "fmt"

type PurchaseService struct {
	notifier Notifier
}

func (p PurchaseService) CheckoutShoppingCart(totalAmount int, message string) error {
	fmt.Printf("Charging payment method with total amount of %d\n", totalAmount)
	return p.notifier.SendMessage(message)
}