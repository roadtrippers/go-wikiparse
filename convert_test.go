package wikiparse

import (
	"testing"
	"reflect"
)

type convertTestInput struct {
	input string
	inputValues []string
	inputUnits []string
	outputUnits []string
	options []string
}

var convertTestData = []convertTestInput{
	convertTestInput{
		"{{convert|1.5|lb|kg}}",
		[]string{
			"1.5",
		},
		[]string{
			"lb",
		},
		[]string{
			"kg",
		},
		[]string{
		},
	},
	convertTestInput{
		"{{convert|1|lb|kg|abbr=on}}",
		[]string{
			"1",
		},
		[]string{
			"lb",
		},
		[]string{
			"kg",
		},
		[]string{
			"abbr=on",
		},
	},
	convertTestInput{
		"{{cvt|1|lb|kg}}",
		[]string{
			"1",
		},
		[]string{
			"lb",
		},
		[]string{
			"kg",
		},
		[]string{
			"abbr=on",
		},
	},
	convertTestInput{
		"{{convert|10,000.1|C|F K}}",
		[]string{
			"10,000.1",
		},
		[]string{
			"°C",
		},
		[]string{
			"°F",
			"K",
		},
		[]string{
		},
	},
	convertTestInput{
		"{{convert|10 x 200 x 3000|m|ft|round=each}}",
		[]string{
			"10 x 200 x 3000",
		},
		[]string{
			"m",
		},
		[]string{
			"ft",
		},
		[]string{
			"round=each",
		},
	},
	convertTestInput{
		"{{convert|6|by|12|ft|m}}",
		[]string{
			"6 by 12",
		},
		[]string{
			"ft",
		},
		[]string{
			"m",
		},
		[]string{
		},
	},
	convertTestInput{
		"{{convert|2|ft|3|in|cm}}",
		[]string{
			"2",
			"3",
		},
		[]string{
			"ft",
			"in",
		},
		[]string{
			"cm",
		},
		[]string{
		},
	},
	convertTestInput{
		"{{convert|1|yd|2|ft|3|in}}",
		[]string{
			"1",
			"2",
			"3",
		},
		[]string{
			"yd",
			"ft",
			"in",
		},
		[]string{},
		[]string{},
	},
	convertTestInput{
		"{{convert|357,000|m²|sqmi|abbr=on}}",
		[]string{
			"357,000",
		},
		[]string{
			"m²",
		},
		[]string{
			"sq mi",
		},
		[]string{
			"abbr=on",
		},
	},
	convertTestInput{
		"{{convert|2|ha|sqm}}",
		[]string{
			"2",
		},
		[]string{
			"ha",
		},
		[]string{
			"sq mi",
		},
		[]string{
		},
	},
	convertTestInput{
		"{{convert|2|ha|xxx}}",
		[]string{
			"2",
			"xxx",
		},
		[]string{
			"ha",
		},
		[]string{
		},
		[]string{
		},
	},
}

func testOneConvert(t *testing.T, ti convertTestInput) {
	conv, err := ParseConvert(ti.input)
	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", ti.input, err)
	} else if !reflect.DeepEqual(conv.InputValues, ti.inputValues) {
		t.Fatalf("Expected input values %v\n Got %v", ti.inputValues, conv.InputValues)
	} else if !reflect.DeepEqual(conv.InputUnits, ti.inputUnits) {
		t.Fatalf("Expected input units %v\n Got %v", ti.inputUnits, conv.InputUnits)
	} else if !reflect.DeepEqual(conv.OutputUnits, ti.outputUnits) {
		t.Fatalf("Expected output units %v\n Got %v", ti.outputUnits, conv.OutputUnits)
	} else if !reflect.DeepEqual(conv.Options, ti.options) {
		t.Fatalf("Expected options %v\n Got %v", ti.options, conv.Options)
	}
}

func TestAllConverts(t *testing.T) {
	t.Parallel()

	for _, ti := range convertTestData {
		testOneConvert(t, ti)
	}
}
