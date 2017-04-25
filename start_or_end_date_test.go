package wikiparse

import (
	"testing"
)

type dateTestInput struct {
	input string
	year int
	month int
	dayofmonth int
	hour int
	min int
	sec int
	offset string
	dayfirst bool
}

var startDateTestData = []dateTestInput{
	dateTestInput{
		"{{start date|1993}}",
		1993,
		-1,
		-1,
		-1,
		-1,
		-1,
		"",
		false,
	},
	dateTestInput{
		"{{start date|1993|02}}",
		1993,
		2,
		-1,
		-1,
		-1,
		-1,
		"",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24}}",
		1993,
		2,
		24,
		-1,
		-1,
		-1,
		"",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30}}",
		1993,
		2,
		24,
		8,
		30,
		-1,
		"",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|||+01:00}}",
		1993,
		2,
		24,
		8,
		-1,
		-1,
		"+01:00",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|||-07:00}}",
		1993,
		2,
		24,
		8,
		-1,
		-1,
		"-07:00",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30|23}}",
		1993,
		2,
		24,
		8,
		30,
		23,
		"",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30|23|Z}}",
		1993,
		2,
		24,
		8,
		30,
		23,
		"Z",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30|23|+01:00}}",
		1993,
		2,
		24,
		8,
		30,
		23,
		"+01:00",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30|23|-07:00}}",
		1993,
		2,
		24,
		8,
		30,
		23,
		"-07:00",
		false,
	},
	dateTestInput{
		"{{start date|1993|02|24|df=y}}",
		1993,
		2,
		24,
		-1,
		-1,
		-1,
		"",
		true,
	},
	dateTestInput{
		"{{start date|1993|02|24|08|30|df=yes}}",
		1993,
		2,
		24,
		8,
		30,
		-1,
		"",
		true,
	},
}

var endDateTestData = []dateTestInput{
	dateTestInput{
		"{{End date|1993|2|24|08|30}}",
		1993,
		2,
		24,
		8,
		30,
		-1,
		"",
		false,
	},
}

func verifyResults(t *testing.T, ti dateTestInput, timeStruct *TimeStruct, err error) {
	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", ti.input, err)
	} else if ti.year != timeStruct.Year {
		t.Fatalf("Expected year %v\n Got %v", ti.year, timeStruct.Year)
	} else if ti.month != timeStruct.Month {
		t.Fatalf("Expected month %v\n Got %v", ti.month, timeStruct.Month)
	} else if ti.dayofmonth != timeStruct.DayOfMonth {
		t.Fatalf("Expected dayofmonth %v\n Got %v", ti.dayofmonth, timeStruct.DayOfMonth)
	} else if ti.hour != timeStruct.Hour {
		t.Fatalf("Expected hour %v\n Got %v", ti.hour, timeStruct.Hour)
	} else if ti.min != timeStruct.Min {
		t.Fatalf("Expected min %v\n Got %v", ti.min, timeStruct.Min)
	} else if ti.sec != timeStruct.Sec {
		t.Fatalf("Expected sec %v\n Got %v", ti.sec, timeStruct.Sec)
	} else if ti.offset != timeStruct.Offset {
		t.Fatalf("Expected offset %v\n Got %v", ti.offset, timeStruct.Offset)
	} else if ti.dayfirst != timeStruct.DayFirst {
		t.Fatalf("Expected day first %v\nGot %v", ti.dayfirst, timeStruct.DayFirst)
	}
}

func testOneStartDate(t *testing.T, ti dateTestInput) {
	timeStruct, err := ParseStartDate(ti.input)
	verifyResults(t, ti, timeStruct, err)
}

func TestAllStartDates(t *testing.T) {
	t.Parallel()

	for _, ti := range startDateTestData {
		testOneStartDate(t, ti)
	}
}

func testOneEndDate(t *testing.T, ti dateTestInput) {
	timeStruct, err := ParseEndDate(ti.input)
	verifyResults(t, ti, timeStruct, err)
}

func TestAllEndDates(t *testing.T) {
	t.Parallel()

	for _, ti := range endDateTestData {
		testOneEndDate(t, ti)
	}
}
