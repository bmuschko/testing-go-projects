# Solution

The existing test cases use mock objects and therefore can be treated as unit tests. Rename the test functions to start with `TestUnit`.

Next, introduce two new test cases to should represent integration tests. The function names should start with `TestIntegration`. Instead of using a mock object, simply pass in the actual implementation of a `Notifier`. For example, the implementations could look as follows:

```go
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
```

We can now run the all tests with the following command:

```shell
$ go test ./... -v
=== RUN   TestUnitSuccessfulCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitSuccessfulCheckoutShoppingCart (0.00s)
    shopping_test.go:18: PASS:	SendMessage(string)
=== RUN   TestUnitFailedCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitFailedCheckoutShoppingCart (0.00s)
    shopping_test.go:31: PASS:	SendMessage(string)
=== RUN   TestIntegrationCheckoutShoppingCartEmailNotifier
Charging payment method with total amount of 99
Sending email purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartEmailNotifier (0.00s)
=== RUN   TestIntegrationCheckoutShoppingCartSMSNotifier
Charging payment method with total amount of 99
Sending SMS purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartSMSNotifier (0.00s)
PASS
ok  	github.com/bmuschko/shopping	0.266s
```

To run just the integration tests, use the following command:

```shell
$ go test ./... -v -run 'TestIntegration'
=== RUN   TestIntegrationCheckoutShoppingCartEmailNotifier
Charging payment method with total amount of 99
Sending email purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartEmailNotifier (0.00s)
=== RUN   TestIntegrationCheckoutShoppingCartSMSNotifier
Charging payment method with total amount of 99
Sending SMS purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartSMSNotifier (0.00s)
PASS
ok  	github.com/bmuschko/shopping	0.108s
```

To run just the unit tests, use the following command:

```shell
$ go test ./... -v -run 'TestUnit'
=== RUN   TestUnitSuccessfulCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitSuccessfulCheckoutShoppingCart (0.00s)
    shopping_test.go:18: PASS:	SendMessage(string)
=== RUN   TestUnitFailedCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitFailedCheckoutShoppingCart (0.00s)
    shopping_test.go:31: PASS:	SendMessage(string)
PASS
ok  	github.com/bmuschko/shopping	0.108s
```

You can find the full solution [here](./separation-by-function-name/shopping_test.go).

To use build tags, you will have to separate unit and integration tests into their own Go test files. In addition, add the appropriate build tags to the files to indicate the type of tests they represent. You can find the full solution in [shopping_unit_test.go](./separation-by-build-tags/shopping_unit_test.go) and [shopping_integration_test.go](./separation-by-build-tags/shopping_integration_test.go).

Running the unit tests via the build tag `unit`:

```shell
$ go test ./... -v -tags=unit
=== RUN   TestUnitSuccessfulCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitSuccessfulCheckoutShoppingCart (0.00s)
    shopping_unit_test.go:21: PASS:	SendMessage(string)
=== RUN   TestUnitFailedCheckoutShoppingCart
Charging payment method with total amount of 99
--- PASS: TestUnitFailedCheckoutShoppingCart (0.00s)
    shopping_unit_test.go:34: PASS:	SendMessage(string)
PASS
ok  	github.com/bmuschko/shopping	0.219s
```

Running the integration tests via the build tag `integration`:

```shell
$ go test ./... -v -tags=integration
=== RUN   TestIntegrationCheckoutShoppingCartEmailNotifier
Charging payment method with total amount of 99
Sending email purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartEmailNotifier (0.00s)
=== RUN   TestIntegrationCheckoutShoppingCartSMSNotifier
Charging payment method with total amount of 99
Sending SMS purchase notification
--- PASS: TestIntegrationCheckoutShoppingCartSMSNotifier (0.00s)
PASS
ok  	github.com/bmuschko/shopping	0.127s
```