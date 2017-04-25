package wikiparse

import (
	"regexp"
	"strings"
	"errors"
	"strconv"
	"time"
)

type StartDateAndAge struct {
	Year int
	Month int
	DayOfMonth int
	DayFirst bool
	UseParens bool
	LineBreak bool
}

var monthToNumber = map[string]int{
	"1": 1,
	"Jan": 1,
	"January": 1,
	"2": 2,
	"Feb": 2,
	"February": 2,
	"3": 3,
	"Mar": 3,
	"March": 3,
	"4": 4,
	"Apr": 4,
	"April": 4,
	"5": 5,
	"May": 5,
	"6": 6,
	"Jun": 6,
	"June": 6,
	"7": 7,
	"Jul": 7,
	"July": 7,
	"8": 8,
	"Aug": 8,
	"August": 8,
	"9": 9,
	"Sept": 9,
	"September": 9,
	"10": 10,
	"Oct": 10,
	"October": 10,
	"11": 11,
	"Nov": 11,
	"November": 11,
	"12": 12,
	"Dec": 12,
	"December": 12,
}

var startDateAndAgeStartRE *regexp.Regexp

func init() {
	startDateAndAgeStartRE = regexp.MustCompile(`(?mi){{start date and age`)
}

func IsStartDateAndAge(text string) bool {
	return startDateAndAgeStartRE.MatchString(text)
}

// ParseStartDateAndAge accepts a Start Date And Age template string and produces a StartDateAndAge
// holding the parsed date/time components along with the display properties. If a component is not
// included, the value will be -1.
func ParseStartDateAndAge(text string) (*StartDateAndAge, error) {
	result := StartDateAndAge{
		Year: -1,
		Month: -1,
		DayOfMonth: -1,
		DayFirst: false,
		UseParens: false,
		LineBreak: false,
	}

	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	parts := strings.Split(cleaned, "|")

	for i, v := range parts {
		parts[i] = bracketReplacer.Replace(v)
	}

	if !IsStartDateAndAge(cleaned) || len(parts) == 0 {
		return nil, errors.New("No start date found")
	}

	var (
		err error
		//ok bool
	)

	for i, v := range parts[1:] {
		if strings.Contains(v, "df=") {
			if v != "df=no" {
				result.DayFirst = true
			}
		} else if strings.Contains(v, "mf=") {
			if v == "mf=no" {
				result.DayFirst = true
			}
		} else if strings.Contains(v, "p=") ||  strings.Contains(v, "paren=") {
			result.UseParens = true
		} else if strings.Contains(v, "br=") {
			result.LineBreak = true
		} else {
			switch i {
			case 0:
				if result.Year, err = strconv.Atoi(v); err != nil {
					return nil, err
				}
			case 1:
				if result.Month, err = strconv.Atoi(v); err != nil {
					var monthT time.Time

					// might have a month string
					if monthT, err = time.Parse("Jan", v); err != nil {
						if monthT, err = time.Parse("January", v); err != nil {
							return nil, err
						}
					}

					result.Month = int(monthT.Month())
				}
			case 2:
				if result.DayOfMonth, err = strconv.Atoi(v); err != nil {
					return nil, err
				}
			}
		}
	}

	return &result, nil
}
