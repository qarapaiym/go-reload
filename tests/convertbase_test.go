package tests

import (
	"testing"

	student "../ConvertBase"
)

func TestConvertBase(t *testing.T) {
	tables := []struct {
		nbr      string
		baseFrom string
		baseTo   string
		answer   string
	}{
		{"4506C", "0123456789ABCDEF", "choumi", "hccocimc"},
		{"babcbaaaaabcb", "abc", "0123456789ABCDEF", "9B611"},
		{"256850", "0123456789", "01", "111110101101010010"},
		{"uuhoumo", "choumi", "Zone01", "eeone0n"},
		{"683241", "0123456789", "0123456789", "683241"},
	}
	for _, table := range tables {
		res := student.ConvertBase(table.nbr, table.baseFrom, table.baseTo)
		if res != table.answer {
			t.Errorf("ConvertBase (%s, %s, %s): [your answer - (%s)] != [(%s) - test's answer]", table.nbr, table.baseFrom, table.baseTo, res, table.answer)
		}
	}
}
