package tests

import (
	"strings"
	"testing"

	student "../AdvancedSortWordArr"
)

func TestAdvancedSortWordArr(t *testing.T) {
	tables := []struct {
		array  []string
		f      func(a, b string) int
		answer []string
	}{
		{
			strings.Split("The earliest computing device undoubtedly consisted of the five fingers of each hand", " "),
			strings.Compare,
			[]string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"},
		},
		{
			strings.Split("The word digital comesfrom \"digits\" or fingers", " "),
			strings.Compare,
			[]string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"},
		},
		{
			strings.Split("a A 1 b B 2 c C 3", " "),
			strings.Compare,
			[]string{"1", "2", "3", "A", "B", "C", "a", "b", "c"},
		},
		{
			strings.Split("The computing consisted device each earliest fingers five hand of of the undoubtedly", " "),
			func(a, b string) int {
				return strings.Compare(b, a)
			},
			[]string{"undoubtedly", "the", "of", "of", "hand", "five", "fingers", "earliest", "each", "device", "consisted", "computing", "The"},
		},
		{
			strings.Split("\"digits\" The comesfrom digital fingers or word", " "),
			func(a, b string) int {
				return strings.Compare(b, a)
			},
			[]string{"word", "or", "fingers", "digital", "comesfrom", "The", "\"digits\""},
		},
		{
			strings.Split("1 2 3 A B C a b c", " "),
			func(a, b string) int {
				return strings.Compare(b, a)
			},
			[]string{"c", "b", "a", "C", "B", "A", "3", "2", "1"},
		},
	}
	for _, table := range tables {
		arr := table.array
		student.AdvancedSortWordArr(arr, table.f)
		if len(arr) != len(table.answer) {
			t.Errorf("SplitWhiteSpaces (%s): your answer (%s) != (%s)", table.array, arr, table.answer)
			break
		}
		for i := 0; i < len(table.answer); i++ {
			if arr[i] != table.answer[i] {
				t.Errorf("AdvancedSortWordArr (%s): [your answer - (%s)] != [(%s) - test's answer]", table.array, arr, table.answer)
				break
			}
		}
	}
}
