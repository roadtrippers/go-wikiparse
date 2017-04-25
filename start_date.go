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
	startDateStartRE = regexp.MustCompile(`(?mi){{start date`)
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

	var err error

	if result.Year, err = strconv.Atoi(parts[1]); err != nil {
		return nil, err
	}

	if len(parts) > 2 && len(parts[2]) > 0 {
		if result.Month, err = strconv.Atoi(parts[2]); err != nil { return nil, err }
	}

	if len(parts) > 3 && len(parts[3]) > 0 {
		if result.DayOfMonth, err = strconv.Atoi(parts[3]); err != nil { return nil, err }
	}

	if len(parts) > 4 && len(parts[4]) > 0 {
		if strings.Contains(parts[4], "df") {
			result.DayFirst = true
		} else if result.Hour, err = strconv.Atoi(parts[4]); err != nil {
			return nil, err
		}
	}

	if len(parts) > 5 && len(parts[5]) > 0 {
		if strings.Contains(parts[5], "df") {
			result.DayFirst = true
		} else if result.Min, err = strconv.Atoi(parts[5]); err != nil {
			return nil, err
		}
	}

	if len(parts) > 6 && len(parts[6]) > 0 {
		if strings.Contains(parts[6], "df") {
			result.DayFirst = true
		} else if result.Sec, err = strconv.Atoi(parts[6]); err != nil {
			return nil, err
		}
	}

	if len(parts) > 7 {
		if strings.Contains(parts[7], "df") {
			result.DayFirst = true
		} else {
			result.Offset = parts[7]
		}
	}

	if len(parts) > 8 {
		if strings.Contains(parts[8], "df") {
			result.DayFirst = true
		}
	}

	return &result, nil
}
