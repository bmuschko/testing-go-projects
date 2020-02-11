package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Book struct {
	author string
	title  string
	year   int
}

func ReadBooks(csvFile io.Reader) ([]*Book, error) {
	var books []*Book
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return books, err
		}

		if len(line) < 3 {
			return books, fmt.Errorf("A book needs to have three attributes")
		}

		b := new(Book)
		b.author = line[0]
		b.title = line[1]
		year, err := strconv.Atoi(line[2])
		if err != nil {
			return books, err
		}
		b.year = year
		books = append(books, b)
	}

	return books, nil
}
