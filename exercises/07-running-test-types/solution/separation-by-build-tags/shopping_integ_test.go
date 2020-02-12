// +build !unit
// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationCheckoutShoppingCartEmailNotifier(t *testing.T) {
	notifier := EmailNotifier{}
	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, "Successfully purchased 10 books")

	assert.Nil(t, err)
}

func TestIntegrationCheckoutShoppingCartSMSNotifier(t *testing.T) {
	notifier := SMSNotifier{}
	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, "Successfully purchased 10 books")

	assert.Nil(t, err)
}
