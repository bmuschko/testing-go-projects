package main

import (
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Open("books.csv")
	if err != nil {
		panic(err)
	}
	books, err := ReadBooks(csvFile)
	if err != nil {
		panic(err)
	}
	for _, b := range books {
		fmt.Printf("%s by %s | %d\n", b.title, b.author, b.year)
	}
}
