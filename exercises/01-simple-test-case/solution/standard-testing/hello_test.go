package main

import "testing"

func TestEmptyName(t *testing.T) {
	name := ""
	language := "english"
	expected := "Hello World"
	actual, err := Hello(name, language)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestEnglish(t *testing.T) {
	name := "Ben"
	language := "english"
	expected := "Hello Ben"
	actual, err := Hello(name, language)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestSpanish(t *testing.T) {
	name := "John"
	language := "spanish"
	expected := "Hola John"
	actual, err := Hello(name, language)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestGerman(t *testing.T) {
	name := "Mary"
	language := "german"
	expected := "Hallo Mary"
	actual, err := Hello(name, language)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestUnsupportedLanguage(t *testing.T) {
	name := "Maurice"
	language := "french"
	expected := "need to provide a supported language"
	_, err := Hello(name, language)

	if err == nil {
		t.Errorf("Should produce an error for an unsupported language")
	}

	actual := err.Error()

	if expected != actual {
		t.Errorf("Error message was incorrect, got: %s, want: %s.", actual, expected)
	}
}
