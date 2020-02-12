# Exercise 9

In this exercise, you will the HTTP test utilities available in Go to emulate a response from a real-world HTTP server. With the experience gained, you'll be able to confidently write unit tests for production source code that would usually require an external endpoint.

1. Navigate to the directory `start`. You will find an existing Go project.
2. Inspect the functionality of the project.
3. Create a new test file named `http_test.go`.
4. Write a test case that uses `net/http/httptest` as a replacement to making a real HTTP call. The emulated HTTP call should respond with a HTTP 200 status code and the response body "Success!".
5. Verify the expected response with appropriate assertions.
6. (Optional) Discuss: How would you implement a test case that returns a different status code than HTTP 200?