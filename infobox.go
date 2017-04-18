package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var infoboxRE *regexp.Regexp

func init() {
	infoboxRE = regexp.MustCompile(`(?mis){{infobox\s*(.[^\s|}]*)\s*(.*)\s*}}`)
}

type Infobox struct {
	TemplateType string
	Attributes map[string]string
}

func ParseInfobox(text string) (*Infobox, error) {
	res := Infobox{}
	res.Attributes = make(map[string]string)
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	matches := infoboxRE.FindAllStringSubmatch(cleaned, -1)

	if len(matches) == 0 || len(matches[0]) < 1 {
		return nil, errors.New("No Infobox found")
	}

	res.TemplateType = matches[0][1]

	properties := formatPropertyRE.FindAllStringSubmatch(matches[0][2], -1)

	for _, prop := range properties {
		attr := prop[1]
		val := strings.TrimSpace(prop[2])
		res.Attributes[attr] = val
	}

	return &res, nil
}

func WithoutInfobox(text string) (string) {
	return infoboxRE.ReplaceAllString(nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""),
		""), "")
}
