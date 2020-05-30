package tests

import (
	"testing"

	student "../atoi"
)

func TestAtoi(t *testing.T) {
	tables := []struct {
		s      string
		answer int
	}{
		{"", 0},
		{"-", 0},
		{"--123", 0},
		{"1", 1},
		{"-3", -3},
		{"8292", 8292},
		{"000000123", 123},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775808", -9223372036854775808},
		{"9223372036854775808", 0},
	}

	for _, table := range tables {
		res := student.Atoi(table.s)
		if res != table.answer {
			t.Errorf("Atoi (%s): [your answer - (%d)] != [(%d) - test's answer]", table.s, res, table.answer)
		}
	}
}
