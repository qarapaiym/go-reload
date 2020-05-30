package tests

import (
	"strings"
	"testing"

	student "../split"
)

func TestSplit(t *testing.T) {
	tables := []struct {
		str     string
		charset string
		answer  []string
	}{
		{
			"|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,",
			"|=choumi=|",
			strings.Split("|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,", "|=choumi=|"),
		},
		{
			"!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,",
			"!==!",
			strings.Split("!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,", "!==!"),
		},
		{
			"AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,",
			"AFJ",
			strings.Split("AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,", "AFJ"),
		},
		{
			"<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,",
			"<<==123==>>",
			strings.Split("<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,", "<<==123==>>"),
		},
	}
	for _, table := range tables {
		res := student.Split(table.str, table.charset)
		if len(res) != len(table.answer) {
			t.Errorf("SplitWhiteSpaces (%s): your answer (%s) != (%s)", table.str, res, table.answer)
			break
		}
		for i := 0; i < len(table.answer); i++ {
			if res[i] != table.answer[i] {
				t.Errorf("SplitWhiteSpaces (%s): [your answer - (%s)] != [(%s) - test's answer]", table.str, res, table.answer)
				break
			}
		}
	}
}
