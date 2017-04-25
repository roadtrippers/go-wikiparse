package wikiparse

import (
	"testing"
)

type startDateTestInput struct {
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

var startDateTestData = []startDateTestInput{
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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
	startDateTestInput{
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

func testOneStartDate(t *testing.T, ti startDateTestInput) {
	startDate, err := ParseStartDate(ti.input)
	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", ti.input, err)
	} else if ti.year != startDate.Year {
		t.Fatalf("Expected year %v\n Got %v", ti.year, startDate.Year)
	} else if ti.month != startDate.Month {
		t.Fatalf("Expected month %v\n Got %v", ti.month, startDate.Month)
	} else if ti.dayofmonth != startDate.DayOfMonth {
		t.Fatalf("Expected dayofmonth %v\n Got %v", ti.dayofmonth, startDate.DayOfMonth)
	} else if ti.hour != startDate.Hour {
		t.Fatalf("Expected hour %v\n Got %v", ti.hour, startDate.Hour)
	} else if ti.min != startDate.Min {
		t.Fatalf("Expected min %v\n Got %v", ti.min, startDate.Min)
	} else if ti.sec != startDate.Sec {
		t.Fatalf("Expected sec %v\n Got %v", ti.sec, startDate.Sec)
	} else if ti.offset != startDate.Offset {
		t.Fatalf("Expected offset %v\n Got %v", ti.offset, startDate.Offset)
	} else if ti.dayfirst != startDate.DayFirst {
		t.Fatalf("Expected day first %v\nGot %v", ti.dayfirst, startDate.DayFirst)
	}
}

func TestAllStartDates(t *testing.T) {
	t.Parallel()

	for _, ti := range startDateTestData {
		testOneStartDate(t, ti)
	}
}
