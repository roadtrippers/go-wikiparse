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

var startDateStartRE *regexp.Regexp

func init() {
	startDateStartRE = regexp.MustCompile(`(?mi){{start date|`)
}

func IsStartDate(text string) bool {
	return startDateStartRE.MatchString(text)
}

// ParseStartDate accepts a Start Date template string and produces a TimeStruct holding the parsed
// date/time components. If a component is not included, the value will be -1 with the exception of
// Offset which will be an empty string.
func ParseStartDate(text string) (*TimeStruct, error) {
	result := TimeStruct{
		Year: -1,
		Month: -1,
		DayOfMonth: -1,
		Hour: -1,
		Min: -1,
		Sec: -1,
		DayFirst: false,
	}

	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	parts := strings.Split(cleaned, "|")

	for i, v := range parts {
		parts[i] = bracketReplacer.Replace(v)
	}

	if !IsStartDate(cleaned) || len(parts) == 0 {
		return nil, errors.New("No start date found")
	}

	var (
		value int
		err error
	)

	for i, v := range parts[1:] {
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
