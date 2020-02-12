# Exercise 8

In this exercise, you will simplify an existing test case for a function that writes a JSON file by introducing a golden file.

1. Navigate to the directory `start`. You will find an existing Go project.
2. Inspect the functionality of the project and execute the test in `books_test.go`.
3. Create a new golden file named `books.golden` in the directory `testdata`. Populate the file with the value of the variable `expectedJson`.
4. Read the golden file from the test and use it for comparing the actual with the expected JSON data.
5. (Optional) Discuss: How would you implement a solution that updates the data in the golden file if the provided input data or JSON structure changes? Implement a solution if you have extra time.