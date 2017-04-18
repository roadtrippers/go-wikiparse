package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var infoboxRE, templateTypeRE, templateLabelRE *regexp.Regexp

func init() {
	infoboxRE = regexp.MustCompile(`(?mis){{infobox.*}}`)
	templateTypeRE = regexp.MustCompile(`(?i){{infobox (.*)`)
	templateLabelRE = regexp.MustCompile(`(?i)([a-zA-Z\-_]+)[ ]?=[ ]?(.*)`)
}

type Infobox struct {
	TemplateType string
	Attributes map[string]string
}

func ParseInfobox(text string) (*Infobox, error) {
	res := Infobox{}
	res.Attributes = make(map[string]string)
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	matches := infoboxRE.FindAllStringSubmatch(cleaned, 1)

	if len(matches) == 0 || len(matches[0]) < 1 {
		return nil, errors.New("No Infobox found")
	}

	theBox := matches[0][0]

	lines := strings.Split(theBox, "\n|")

	for idx, line := range lines {
		if idx == 0 {
			templateMatches := templateTypeRE.FindAllStringSubmatch(line, -1)
			if len(templateMatches) > 0 && len(templateMatches[0]) > 1 {
				res.TemplateType = templateMatches[0][1]
			}
		} else {
			templateLabelMatches := templateLabelRE.FindAllStringSubmatch(line, -1)
			if len(templateLabelMatches) > 0 && len(templateLabelMatches[0]) > 1 {
				attr := templateLabelMatches[0][1]
				val := templateLabelMatches[0][2]
				res.Attributes[attr] = val
			}
		}
	}

	return &res, nil
}
