package wikiparse

import "testing"

type startDateAndAgeTestInput struct {
	input string
	year int
	month int
	dayofmonth int
	dayfirst bool
	useparens bool
	linebreak bool
}

var startDateAndAgeTestData = []startDateAndAgeTestInput{
	startDateAndAgeTestInput{
		"{{Start date and age|2010|01|02}}",
		2010,
		1,
		2,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2008|Jan|09|df=no}}",
		2008,
		1,
		9,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|January|05|mf=yes}}",
		2003,
		1,
		5,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2010|1|2|df=yes}}",
		2010,
		1,
		2,
		true,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2008|Jan|9|df=y}}",
		2008,
		1,
		9,
		true,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|January|5|df=1}}",
		2003,
		1,
		5,
		true,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2001|9}}",
		2001,
		9,
		-1,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|Sep}}",
		2003,
		9,
		-1,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2005}}",
		2005,
		-1,
		-1,
		false,
		false,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|02|15|df=y|p=y}}",
		2003,
		2,
		15,
		true,
		true,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|paren=yes}}",
		2003,
		-1,
		-1,
		false,
		true,
		false,
	},
	startDateAndAgeTestInput{
		"{{Start date and age|2003|02|15|df=y|br=y}}",
		2003,
		2,
		15,
		true,
		false,
		true,
	},
}
func testOneStartDateAndAge(t *testing.T, ti startDateAndAgeTestInput) {
	startDate, err := ParseStartDateAndAge(ti.input)
	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", ti.input, err)
	} else if ti.year != startDate.Year {
		t.Fatalf("Expected year %v\n Got %v", ti.year, startDate.Year)
	} else if ti.month != startDate.Month {
		t.Fatalf("Expected month %v\n Got %v", ti.month, startDate.Month)
	} else if ti.dayofmonth != startDate.DayOfMonth {
		t.Fatalf("Expected dayofmonth %v\n Got %v", ti.dayofmonth, startDate.DayOfMonth)
	} else if ti.dayfirst != startDate.DayFirst {
		t.Fatalf("Expected dayfirst %v\n Got %v", ti.dayfirst, startDate.DayFirst)
	} else if ti.useparens != startDate.UseParens {
		t.Fatalf("Expected useparens %v\n Got %v", ti.useparens, startDate.UseParens)
	} else if ti.linebreak != startDate.LineBreak {
		t.Fatalf("Expected linebreak %v\n Got %v", ti.linebreak, startDate.LineBreak)
	}
}

func TestAllStartDateAndAges(t *testing.T) {
	t.Parallel()

	for _, ti := range startDateAndAgeTestData {
		testOneStartDateAndAge(t, ti)
	}
}

