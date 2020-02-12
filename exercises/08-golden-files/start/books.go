package main

import (
	"encoding/json"
	"io/ioutil"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
}

func WriteBooks(books []Book, filename string) error {
	jsonData, err := json.MarshalIndent(books, "", "    ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, jsonData, 0644)
}