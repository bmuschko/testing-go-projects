# Solution

Start by creating a `struct` that models the test case. The `struct` defines the given and expected values. Create a slice of test cases with the relevant data. To model the table-driven test simply iterate over the list of test cases, execute the code under test and verify the expected behavior. You can find the full solution [here](./table-driven/books_test.go).

There are two things you have to do to enable running the test cases in parallel. First, capture the test case as iteration value. Second, make the `Parallel()` function call inside of the subtest definition. You can find the full solution [here](./parallel/books_test.go).