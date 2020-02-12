# Exercise 6

In this exercise, you will create a mock object for service to emulate specific behavior without reaching out to external services and/or components.

1. Navigate to the directory `start`. You will find an existing Go project.
2. Inspect the functionality of the project.
3. Add the external package Testify by running `go get github.com/stretchr/testify`.
4. Create a new test file named `shopping_test.go`.
5. Write a test case for the function `CheckoutShoppingCart`. Create a mock object for `Notifier` and provide it when creating `PurchaseService`. Emulate the use case where a customer could check out the shopping cart and a message could be sent successfully.
6. Execute the test case to ensure that it passes.
7. Write a test case for the function `CheckoutShoppingCart`. Create a mock object for `Notifier` and provide it when creating `PurchaseService`. Emulate the use case where a customer could check out the shopping cart but the notifier returns with an error.
8. Execute the test case to ensure that it passes.