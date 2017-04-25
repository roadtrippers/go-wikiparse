package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var infoboxRE, infoboxStartRE *regexp.Regexp

func init() {
	infoboxRE = regexp.MustCompile(`(?mis){{infobox\s*(.[^\s|}]*)\s*(.*)\s*}}`)
	infoboxStartRE = regexp.MustCompile(`(?i){{infobox`)
}

type Infobox struct {
	TemplateType string
	Attributes map[string]string
}

func infoboxBounds(text string) (int, int) {
	idxs := infoboxStartRE.FindAllStringIndex(text, -1)
	if len(idxs) < 1 || len(idxs[0]) < 1 { return 0, 0 }

	start := idxs[0][0]
	end := start
	bracesCounter := 0

	for i := start; i < len(text); i++ {
		if text[i] == '{' {
			bracesCounter++
		} else if text[i] == '}' {
			bracesCounter--
		}

		if bracesCounter == 0 {
			end = i + 1
			break
		}
	}

	return start, end
}

func infoboxText(text string) string {
	start, end := infoboxBounds(text)
	return text[start:end]
}

func ParseInfobox(text string) (*Infobox, error) {
	res := Infobox{}
	res.Attributes = make(map[string]string)
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	infoboxString := infoboxText(cleaned)
	matches := infoboxRE.FindAllStringSubmatch(infoboxString, -1)

	if len(matches) == 0 || len(matches[0]) == 0 {
		return nil, errors.New("No Infobox found")
	}

	res.TemplateType = matches[0][1]

	properties := formatPropertyRE.FindAllStringSubmatch(matches[0][2], -1)

	// TODO: Parse microformats
	// Microfomats:
	// - start date
	// - start date and age
	// - end date
	// - URL
	for _, prop := range properties {
		attr := strings.TrimSpace(prop[1])
		val := strings.TrimSpace(prop[2])
		res.Attributes[attr] = val
	}

	return &res, nil
}

func WithoutInfobox(text string) (string) {
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	return strings.Replace(cleaned, infoboxText(cleaned), "", -1)
}
