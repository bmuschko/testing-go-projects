# Exercise 5

In this exercise, you will implement a table-driven test with different CSV files. Later, you will enhance the test logic to run the verification of the cases in parallel.

1. Copy the [solution from the previous exercise](../03-test-fixtures/solution/testdata-fixture) that reads test data from the `testdata` directory.
2. In the `testdata` directory, rename the existing CSV from `books.csv` to `valid.csv`. Add more CSV files to the `testdata` directory. One file should be empty, the other file should contain an invalid CSV structure. You should end up with three files named `empty.csv`, `invalid.csv`, `valid.csv`.
3. Turn the existing test code into a table-driven test. For each test case defined by the table, specify a `struct` that points to the `testdata` file, captures if the parsing can be handled without error and how many records to expect. Run each test case as a subtest. The assertions should only verify the number of expected records and whether they could be parsed properly.
4. Execute the test case to ensure that it passes.
5. Change the test case so that each of the cases can be executed in parallel.