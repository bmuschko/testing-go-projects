# Solution

Start by writing a test file that provide the CSV data as string. You can find the full solution [here](./string-fixture/books_test.go).

You can take the code from the previous step as starting point to read the CSV data from a temporary file. You can find the full solution [here](./tmpfile-fixture/books_test.go).

Finally, copy-paste the file `books.csv` to the `testdata` directory. Read the persistent file instead of the temporary file. You can find the full solution [here](./testdata-fixture/books_test.go).