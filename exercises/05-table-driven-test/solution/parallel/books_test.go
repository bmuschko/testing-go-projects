package main

import (
	"os"
	"testing"
)

func TestReadBooks(t *testing.T) {
	cases := []struct {
		fixture string
		err     bool
		name    string
		records int
	}{
		{
			fixture: "testdata/empty.csv",
			err:     false,
			name:    "Empty",
			records: 0,
		},
		{
			fixture: "testdata/invalid.csv",
			err:     true,
			name:    "Invalid",
			records: 1,
		},
		{
			fixture: "testdata/valid.csv",
			err:     false,
			name:    "Valid",
			records: 3,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			csvFile, err := os.Open(tc.fixture)
			if err != nil {
				t.Fatal("Can't open CSV file")
			}
			books, err := ReadBooks(csvFile)

			if err != nil && !tc.err {
				t.Errorf("Expected an error for file %s", tc.fixture)
			}

			if tc.records != len(books) {
				t.Errorf("Unexpected number of books, got: %d, want: %d.", len(books), tc.records)
			}
		})
	}
}
