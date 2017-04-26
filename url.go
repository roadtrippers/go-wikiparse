package wikiparse

import (
	"regexp"
	"strings"
	"errors"
	"net/url"
	"fmt"
)

type URLData struct {
	Url string
	DisplayAs string
}

var urlStartRE *regexp.Regexp

func init() {
	urlStartRE = regexp.MustCompile(`(?mi){{\s*url\s*\|`)
}

func IsURL(text string) bool {
	return urlStartRE.MatchString(text)
}

func ParseURL(text string) (*URLData, error) {
	result := URLData{}

	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	parts := strings.Split(cleaned, "|")

	if len(parts) < 2 || !IsURL(cleaned) {
		return nil, errors.New("No URL found")
	}

	for i, v := range parts {
		parts[i] = strings.TrimSpace(bracketReplacer.Replace(v))
	}

	if strings.Index(parts[1], "1=") == 0 {
		result.Url = parts[1][2:]
	} else {
		result.Url = parts[1]
	}

	if len(parts) > 2 {
		result.DisplayAs = parts[2]
	} else {
		// strip protocol specifier from Url for display
		urlWithScheme := result.Url
		if !strings.Contains(urlWithScheme, "://") || urlWithScheme[0:2] == "//" {
			urlWithScheme = fmt.Sprintf("//%s", urlWithScheme)
		}

		if parsed, err := url.Parse(urlWithScheme); err != nil {
			return nil, err
		} else {
			if len(parsed.Host) > 0 {
				result.DisplayAs = strings.ToLower(parsed.Host)
			}

			if len(parsed.Path) > 1 {
				result.DisplayAs = fmt.Sprintf("%s%s", result.DisplayAs, parsed.Path)
			}

			if len(parsed.RawQuery) > 0 {
				result.DisplayAs = fmt.Sprintf("%s?%s", result.DisplayAs, parsed.RawQuery)
			}

			if len(parsed.Fragment) > 0 {
				result.DisplayAs = fmt.Sprintf("%s#%s", result.DisplayAs, parsed.Fragment)
			}
		}
	}

	return &result, nil
}
