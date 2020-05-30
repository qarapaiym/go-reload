package tests

import (
	"testing"

	student "../atoibase"
)

func TestAtoiBase(task *testing.T) {
	tables := []struct {
		nbr    string
		base   string
		answer int
	}{
		{"bcbbbbaab", "abc", 12016},
		{"0001", "01", 1},
		{"00", "01", 0},
		{"saDt!I!sI", "CHOUMIisDAcat!", 11557277891},
		{"AAho?Ao", "WhoAmI?", 406772},
		{"thisinputshouldnotmatter", "abca", 0},
		{"125", "0123456789", 125},
		{"uoi", "choumi", 125},
		{"bbbbbab", "-ab", 0},
	}
	for _, table := range tables {
		res := student.AtoiBase(table.nbr, table.base)
		if res != table.answer {
			task.Errorf("AtoiBase (%s, %s): [your answer - (%d)] != [(%d) - test's answer]", table.nbr, table.base, res, table.answer)
		}
	}
}
