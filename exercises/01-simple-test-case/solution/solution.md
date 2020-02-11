# Solution

Navigate to the directory `start` and create a new test file.

```shell
$ cd start
$ touch hello_test.go
```

Implement test cases as shown in the file `hello_test.go` in the [solution](./standard-testing/hello_test.go) folder.

Execute tests with the following command. Use the option `-v` to render all test cases. All tests should pass.

```shell
$ go test -v
=== RUN   TestEmptyName
--- PASS: TestEmptyName (0.00s)
=== RUN   TestEnglish
--- PASS: TestEnglish (0.00s)
=== RUN   TestSpanish
--- PASS: TestSpanish (0.00s)
=== RUN   TestGerman
--- PASS: TestGerman (0.00s)
=== RUN   TestUnsupportedLanguage
--- PASS: TestUnsupportedLanguage (0.00s)
PASS
ok  	github.com/bmuschko/hello	0.187s
```

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

Change the existing test cases as shown in the file `hello_test.go` in the [solution](./testify/hello_test.go) folder.

Execute tests with the following command. Use the option `-v` to render all test cases. All tests should pass.

```shell
$ go test -v
=== RUN   TestEmptyName
--- PASS: TestEmptyName (0.00s)
=== RUN   TestEnglish
--- PASS: TestEnglish (0.00s)
=== RUN   TestSpanish
--- PASS: TestSpanish (0.00s)
=== RUN   TestGerman
--- PASS: TestGerman (0.00s)
=== RUN   TestUnsupportedLanguage
--- PASS: TestUnsupportedLanguage (0.00s)
PASS
ok  	github.com/bmuschko/hello	0.187s
```