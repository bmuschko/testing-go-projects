package main

func main() {
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

	err := WriteBooks(b, "books.json")
	if err != nil {
		panic(err)
	}
}