# Exercise 4

In this exercise, you will create reusable functionality for setting up test infrastructure in the form of a test helper function. More specifically, you will create a test helper for creating a temporary test directory.

1. Copy the [solution from the previous exercise](../03-test-fixtures/solution/tmpfile-fixture) that stands up test data in a temporary directory.
2. Refactoring the existing test code from `books_test.go` so that the code for creating/removing the temporary file can be reused across different test Go files. The contents of the test file can be provided as parameter.
3. Externalize the logic into a test helper in the package `testhelper`. Call the testhelper logic from the existing test file.
4. Execute the test case to ensure that it passes.
5. Reuse the test helper code in a second test case.
6. Execute both test cases to ensure that they pass.