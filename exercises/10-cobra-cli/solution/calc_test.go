package main

import (
	"bytes"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	calcAdd := &calcAddCmd{
		a: 1,
		b: 2,
		out: buffer,
	}
	err := calcAdd.run()

	if err != nil {
		t.Error("Should not produce an error")
	}

	expected := "Result: 3\n"
	actual := buffer.String()

	if expected != actual {
		t.Errorf("Unexpected result, got: %s, want: %s.", actual, expected)
	}
}
