package tests

import (
	"testing"

	student "../SplitWhiteSpaces"
)

func TestSplitWhiteSpaces(t *testing.T) {
	tables := []struct {
		str    string
		answer []string
	}{
		{
			"The earliest foundations of what would become computer science predate the invention of the modern digital computer",
			[]string{"The", "earliest", "foundations", "of", "what", "would", "become", "computer", "science", "predate", "the", "invention", "of", "the", "modern", "digital", "computer"},
		},
		{
			"Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,",
			[]string{"Machines", "for", "calculating", "fixed", "numerical", "tasks", "such", "as", "the", "abacus", "have", "existed", "since", "antiquity,"},
		},
		{
			"aiding in computations such as multiplication and division .",
			[]string{"aiding", "in", "computations", "such", "as", "multiplication", "and", "division", "."},
		},
		{
			"Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.",
			[]string{"Algorithms", "for", "performing", "computations", "have", "existed", "since", "antiquity,", "even", "before", "the", "development", "of", "sophisticated", "computing", "equipment."},
		},
		{
			"Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]",
			[]string{"Wilhelm", "Schickard", "designed", "and", "constructed", "the", "first", "working", "mechanical", "calculator", "in", "1623.[4]"},
		},
		{
			"In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,",
			[]string{"In", "1673,", "Gottfried", "Leibniz", "demonstrated", "a", "digital", "mechanical", "calculator,"},
		},
	}

	for _, table := range tables {
		res := student.SplitWhiteSpaces(table.str)
		if len(res) != len(table.answer) {
			t.Errorf("SplitWhiteSpaces: [your answer - (%s)] != [(%s) - test's answer]", res, table.answer)
			break
		}
		for i := range res {
			if res[i] != table.answer[i] {
				t.Errorf("SplitWhiteSpaces: [your answer - (%s)] != [(%s) - test's answer]", res, table.answer)
				break
			}
		}
	}
}
