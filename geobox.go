package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var geoboxStartRE, geoboxLineRE *regexp.Regexp

func init() {
	geoboxStartRE = regexp.MustCompile(`(?mis){{\s*geobox\s*\|\s*(.*)`)
	geoboxLineRE = regexp.MustCompile(`(?i)\|.*?=.*`)
}

type Geobox struct {
	TemplateType string
	Attributes map[string]string
}

func IsGeobox(text string) bool {
	return geoboxStartRE.MatchString(text)
}

func ParseGeobox(text string) (*Geobox, error) {
	res := Geobox{
		Attributes: make(map[string]string),
	}

	parts := partsFromText(text, geoboxStartRE)

	if len(parts) == 0 || !IsGeobox(text) {
		return nil, errors.New("No geobox found")
	}

	res.TemplateType = parts[1]
	res.Attributes = attributesToMap(parts[2:])

	return &res, nil
}

func geoboxText(text string) string {
	start, end := templateBounds(text, geoboxStartRE)
	return text[start:end]
}

func WithoutGeobox(text string) (string) {
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	return strings.Replace(cleaned, geoboxText(cleaned), "", -1)
}
