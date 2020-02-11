# Exercise 3

In this exercise, you will provide test data to a test case as input. You will implement different options for reading the data.

1. Navigate to the directory `start`. You will find an existing Go project.
2. Inspect the functionality of the project.
3. Write a test file named `books_test.go` with a single test case. The test case should use a `string` as input data. Copy the data from the file `books.csv`. The test case should verify that the data has been parsed correctly.
4. Execute the test case to ensure that it passes.
5. Modify the test case to read the data from a temporary file. The file should be created as fixture for test execution and deleted after the test has been run.
6. Execute the test case to ensure that it passes.
7. Modify the test case to read the data from a permanent file under the `testdata` directory.
8. Execute the test case to ensure that it passes.