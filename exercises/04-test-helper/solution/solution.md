# Solution

Create a directory named `testhelper`, navigate to the directory and create a new file named `file_helpers.go`.

```shell
$ mkdir testhelper
$ cd testhelper
$ touch file_helper.go
```

In the file `file_helper.go`, create two functions. Function 1 simply creates a temporary file and returns its. Function 2 creates a temporary file but will populate it automatically with plain-text content. You can find the full file [here](./testhelper/file_helpers.go).

You can now import the `testhelper` package in the test file `books_test.go`. Call the relevant functions. You should see that the setup/cleanup logic can be simplified significantly. You can find the full solution [here](./books_test.go).