# Solution

Create the [golden file](./testdata/books.golden) with the following content:

```json
[
    {
        "author": "John Grisham",
        "title": "The Guardians",
        "year": 2019
    },
    {
        "author": "Harper Lee",
        "title": "To Kill a Mockingbird",
        "year": 2005
    },
    {
        "author": "Leo Tolstoy",
        "title": "War and Peace",
        "year": 1982
    }
]
```

In the [test file](./books_test.go), read the golden file and use it for comparing it with the actual JSON data.

```go
expectedJsonData, err := ioutil.ReadFile("testdata/books.golden")
if err != nil {
	t.Fatalf("Failed to read golden JSON data: %s", err)
}
expectedJson := string(expectedJsonData)
```

If you wanted to implement update functionality, start by introducing a new command line option `-update` as part of `books_test.go`.

```go
var (
    update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
    flag.Parse()
    os.Exit(m.Run())
}
```

Next, you'd change the implementation of the test case to update the data in the golden file with the actual data produced by the production implementation whenever the `-update` has been provided e.g. `go test ./... -update`.