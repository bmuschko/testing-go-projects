package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyName(t *testing.T) {
	name := ""
	language := "english"
	expected := "Hello World"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEnglish(t *testing.T) {
	name := "Ben"
	language := "english"
	expected := "Hello Ben"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSpanish(t *testing.T) {
	name := "John"
	language := "spanish"
	expected := "Hola John"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGerman(t *testing.T) {
	name := "Mary"
	language := "german"
	expected := "Hallo Mary"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnsupportedLanguage(t *testing.T) {
	name := "Maurice"
	language := "french"
	expected := "need to provide a supported language"
	_, err := Hello(name, language)

	assert.NotNil(t, err)
	actual := err.Error()
	assert.Equal(t, expected, actual)
}
