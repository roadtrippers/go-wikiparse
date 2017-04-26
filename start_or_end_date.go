package wikiparse

import (
	"regexp"
	"strings"
	"errors"
	"strconv"
)

type TimeStruct struct {
	Year int
	Month int
	DayOfMonth int
	Hour int
	Min int
	Sec int
	Offset string
	DayFirst bool
}

var startDateStartRE, endDateStartRE *regexp.Regexp

func init() {
	startDateStartRE = regexp.MustCompile(`(?mi){{\s*start date\s*\|`)
	endDateStartRE = regexp.MustCompile(`(?mi){{\s*end date\s*\|`)
}

func IsStartDate(text string) bool {
	return startDateStartRE.MatchString(text)
}

func IsEndDate(text string) bool {
	return endDateStartRE.MatchString(text)
}

func timeStructFromParts(parts []string) (*TimeStruct, error) {
	var (
		value int
		err error
	)

	result := TimeStruct{
		Year: -1,
		Month: -1,
		DayOfMonth: -1,
		Hour: -1,
		Min: -1,
		Sec: -1,
		DayFirst: false,
	}

	for i, v := range parts {
		if strings.Contains(v, "df=") {
			result.DayFirst = true
		} else if len(v) > 0 {
			if value, err = strconv.Atoi(v); err != nil {
				if i == 6 {
					result.Offset = v
				} else {
					return nil, err
				}
			} else {
				switch i {
				case 0:
					result.Year = value
				case 1:
					result.Month = value
				case 2:
					result.DayOfMonth = value
				case 3:
					result.Hour = value
				case 4:
					result.Min = value
				case 5:
					result.Sec = value
				}
			}
		}
	}

	return &result, nil
}

// ParseStartDate accepts a Start Date template string and produces a TimeStruct holding the parsed
// date/time components. If a component is not included, the value will be -1 with the exception of
// Offset which will be an empty string.
func ParseStartDate(text string) (*TimeStruct, error) {
	parts := partsFromText(text)

	if len(parts) == 0 || !IsStartDate(text) {
		return nil, errors.New("No start date found")
	}

	return timeStructFromParts(parts[1:])
}

// ParseEndDate accepts a End Date template string and produces a TimeStruct holding the parsed
// date/time components. If a component is not included, the value will be -1 with the exception of
// Offset which will be an empty string.
func ParseEndDate(text string) (*TimeStruct, error) {
	parts := partsFromText(text)

	if len(parts) == 0 || !IsEndDate(text) {
		return nil, errors.New("No end date found")
	}

	return timeStructFromParts(parts[1:])
}
