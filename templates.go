package wikiparse

import (
	"regexp"
	"strings"
)

var nowikiRE, commentRE *regexp.Regexp

func init() {
	nowikiRE = regexp.MustCompile(`(?ms)<nowiki>.*</nowiki>`)
	commentRE = regexp.MustCompile(`(?ms)(<|&lt;)!--.*?--(>|&gt;)`)
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

func templateBounds(text string, startRE *regexp.Regexp) (int, int) {
	idxs := startRE.FindAllStringIndex(text, -1)
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

func bracketsMatch(left, right byte) bool {
	return (left == '{' && right == '}') || (left == '[' && right == ']')
}

func templateAttributes(templateText string) []string {
	attrs := make([]string, 0)
	start := 2

	if len(templateText) <= start { return attrs }

	bracketStack := NewStack()

	for i := start; i < len(templateText) - 2; i++ {
		if templateText[i] == '{' || templateText[i] == '[' {
			bracketStack.Push(templateText[i])
		} else if (templateText[i] == '}' || templateText[i] == ']') &&
			bracketsMatch(bracketStack.Peek().(byte), templateText[i]) {
			bracketStack.Pop()
		} else if templateText[i] == '|' && bracketStack.Len() == 0 {
			attrs = append(attrs, templateText[start:i])
			start = i + 1
		}
	}

	attrs = append(attrs, templateText[start:len(templateText) - 2])

	return attrs
}

func attributesToMap(attrsArray []string) map[string]string {
	attrs := make(map[string]string)

	for _, prop := range attrsArray {
		propParts := strings.Split(prop, "=")

		if len(propParts) > 0 {
			attr := strings.TrimSpace(strings.Replace(propParts[0], "|", "", -1))
			val := ""

			if len(propParts) > 1 {
				val = strings.TrimSpace(strings.Join(propParts[1:], "="))
			}

			attrs[attr] = val
		}
	}

	return attrs
}

func partsFromText(text string, startRE *regexp.Regexp) []string {
	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	start, end := templateBounds(cleaned, startRE)
	parts := templateAttributes(cleaned[start:end])

	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}

	return parts
}
