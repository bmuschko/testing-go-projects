# Exercise 7

In this exercise, you will introduce integration tests alongside already existing unit tests. You will practice the ability to run unit tests and integration separately, as well as the whole test suite.

1. Copy the [solution from the previous exercise](../06-mocking/solution) that creates and uses mock objects to a new directory.
2. Introduce two integration test cases in `shopping_test.go` that make use of the actual implementations of `Notifier`, one for `EmailNotifier` and one for `SMSNotifier`. Differentiate unit and integration tests by appropriate function names. _Note:_ We are assuming that those implementation send real notification by email or SMS. In fact, they are only logging a message to avoid having to set up the integration with external services.
3. Execute all types of tests from the command line. Ensure that the correct tests have been run by providing the `-v` CLI option.
4. Execute only the integration tests. Ensure that the correct tests have been run by providing the `-v` CLI option.
5. Execute only the unit tests. Ensure that the correct tests have been run by providing the `-v` CLI option.
6. (Optional) Discuss: What would need to happen in order to use build tags to differentiate the test cases?