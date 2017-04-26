package wikiparse

import (
	"regexp"
	"strings"
)

var nowikiRE, commentRE *regexp.Regexp

func init() {
	nowikiRE = regexp.MustCompile(`(?ms)<nowiki>.*</nowiki>`)
	commentRE = regexp.MustCompile(`(?ms)<!--.*-->`)
}

func ParseTemplate(text string) (interface{}, error) {
	switch {
	case IsMFAdr(text):
		return ParseMFAdr(text)
	case IsAltCoords(text):
		return ParseAltCoords(text)
	case IsCoords(text):
		return ParseCoords(text)
	case IsConvert(text):
		return ParseConvert(text)
	case IsStartDate(text):
		return ParseStartDate(text)
	case IsEndDate(text):
		return ParseEndDate(text)
	case IsStartDateAndAge(text):
		return ParseStartDateAndAge(text)
	case IsURL(text):
		return ParseURL(text)
	}

	return text, nil
}

func partsFromText(text string) []string {
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	parts := strings.Split(cleaned, "|")

	for i, v := range parts {
		parts[i] = strings.TrimSpace(bracketReplacer.Replace(v))
	}

	return parts
}
