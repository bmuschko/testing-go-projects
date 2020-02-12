package main

import (
	"github.com/bmuschko/books/testhelper"
	"io/ioutil"
	"os"
	"testing"
)

func setup(t *testing.T) *os.File {
	return testhelper.TmpFile(t, "", "books.json")
}

func cleanup(file *os.File) {
	os.Remove(file.Name())
}

func TestWriteBooks(t *testing.T) {
	file := setup(t)
	defer cleanup(file)

	b := []Book{
		{
			Title:  "The Guardians",
			Author: "John Grisham",
			Year:   2019,
		},
		{
			Title:  "To Kill a Mockingbird",
			Author: "Harper Lee",
			Year:   2005,
		},
		{
			Title:  "War and Peace",
			Author: "Leo Tolstoy",
			Year:   1982,
		},
	}

	err := WriteBooks(b, file.Name())
	if err != nil {
		t.Errorf("Unable to write JSON data: %s", err.Error())
	}

	actualJson, err := ioutil.ReadFile(file.Name())
	expectedJson := `[
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
]`
	if expectedJson != string(actualJson) {
		t.Errorf("Generated JSON doesn't meet expectations.\nExpected:\n%s\nActual:\n%s", expectedJson, actualJson)
	}
}
