package main

func main() {
	notifier := EmailNotifier{}
	purchaseService := PurchaseService{ notifier: notifier}
	purchaseService.CheckoutShoppingCart(99, "Successfully purchased 10 books")
}