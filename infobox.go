package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var  infoboxStartRE *regexp.Regexp

func init() {
	infoboxStartRE = regexp.MustCompile(`(?i){{infobox`)
}

type Infobox struct {
	TemplateType string
	Attributes map[string]string
}

func infoboxText(text string) string {
	start, end := templateBounds(text, infoboxStartRE)
	return text[start:end]
}

func IsInfobox(text string) bool {
	return infoboxStartRE.MatchString(text)
}

func ParseInfobox(text string) (*Infobox, error) {
	res := Infobox{
		Attributes: make(map[string]string),
	}

	parts := partsFromText(text, infoboxStartRE)

	if len(parts) == 0 || !IsInfobox(text) {
		return nil, errors.New("No Infobox found")
	}

	infoboxSplits := strings.Split(parts[0], " ")

	if len(infoboxSplits) < 2 {
		infoboxSplits = strings.Split(parts[0], "\t")
	}

	if len(infoboxSplits) > 1 {
		res.TemplateType = strings.TrimSpace(strings.Join(infoboxSplits[1:], " "))
	}
	res.Attributes = attributesToMap(parts[1:])

	return &res, nil
}

func WithoutInfobox(text string) (string) {
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	return strings.Replace(cleaned, infoboxText(cleaned), "", -1)
}
