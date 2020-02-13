package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := NewRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
