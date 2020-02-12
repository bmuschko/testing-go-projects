# Solution

Add Testify with the help of the `go get` command.

```shell
$ go get github.com/stretchr/testify
```

The contents of the `go.mod` file should look as follows:

```
module github.com/bmuschko/hello

go 1.13

require github.com/stretchr/testify v1.4.0
```

Start by creating a `struct` for the mock object.

```go
type notifierMock struct {
	mock.Mock
}

func (m *notifierMock) SendMessage(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
```

With the `struct` in place, you can set up both test cases. For the first test case, tell the mock object to return `nil`. For the second test case, return an `error`. You can find the solution for both test cases [here](./shopping_test.go).