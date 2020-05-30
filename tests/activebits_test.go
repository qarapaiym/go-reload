package tests

import (
	"testing"

	student "../ActiveBits"
)

func TestActiveBits(task *testing.T) {
	tables := []struct {
		n      int
		answer uint
	}{
		{15, 4},
		{17, 2},
		{4, 1},
		{11, 3},
		{9, 2},
		{12, 2},
		{2, 1},
	}
	for _, table := range tables {
		res := student.ActiveBits(table.n)
		if res != table.answer {
			task.Errorf("ActiveBits (%d): [your answer - (%d)] != [(%d) - test's answer]", table.n, res, table.answer)
		}
	}
}
