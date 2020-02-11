package main

import "errors"

func Hello(name string, language string) (string, error) {
	if name == "" {
		name = "World"
	}

	prefix := ""

	switch language {
	case "english":
		prefix = "Hello"
	case "spanish":
		prefix = "Hola"
	case "german":
		prefix = "Hallo"
	default:
		return "", errors.New("need to provide a supported language")
	}

	return prefix + " " + name, nil
}
