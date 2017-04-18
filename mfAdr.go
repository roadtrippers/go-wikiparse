package wikiparse

import (
	"regexp"
	"errors"
	"strings"
)

var adrRE, formatPropertyRE *regexp.Regexp

func init() {
	adrRE = regexp.MustCompile(`(?mis){{mf-adr\s*(.*)\s*}}`)
	formatPropertyRE = regexp.MustCompile(`(?i)\|([a-zA-Z\-_\s]+)=(.*)`)
}

type MFAdr struct {
	Street string
	City string
	Region string
	PostalCode string
	Nation string
}

func ParseMFAdr(text string) (*MFAdr, error) {
	res := MFAdr{}

	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	matches := adrRE.FindAllStringSubmatch(cleaned, -1)

	if len(matches) == 0 || len(matches[0]) == 0 {
		return nil, errors.New("No address found")
	}

	properties := formatPropertyRE.FindAllStringSubmatch(matches[0][1], -1)

	if len(properties) == 0 {
		return nil, errors.New("No address components found")
	}

	for _, prop := range properties {
		switch strings.TrimSpace(prop[1]) {
		case "street":
			res.Street = strings.TrimSpace(prop[2])
		case "city":
			res.City = strings.TrimSpace(prop[2])
		case "region":
			res.Region = strings.TrimSpace(prop[2])
		case "pocode":
			res.PostalCode = strings.TrimSpace(prop[2])
		case "nation":
			res.Nation = strings.TrimSpace(prop[2])
		}
	}

	return &res, nil
}

