package main

import "fmt"

func main() {
	msg, err := Hello("attendee", "english")
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}
