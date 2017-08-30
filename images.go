package wikiparse

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"regexp"
	"strings"
)

var fileRE *regexp.Regexp
var imageRE *regexp.Regexp

func init() {
	fileRE = regexp.MustCompile(`\[File:([^\|\]]+)`)
	imageRE = regexp.MustCompile(`\[Image:([^\|\]]+)`)
}

// FindFiles finds all the File references from within an article
// body.
//
// This includes things in comments, as many I found were commented
// out.
func FindFiles(text string) []string {
	cleaned := nowikiRE.ReplaceAllString(text, "")
	fileMatches := fileRE.FindAllStringSubmatch(cleaned, 10000)
	imageMatches := imageRE.FindAllStringSubmatch(cleaned, 10000)

	rv := []string{}

	for _, x := range fileMatches {
		rv = append(rv, x[1])
	}

	for _, x := range imageMatches {
		rv = append(rv, x[1])
	}

	return rv
}

// URLForFile gets the wikimedia URL for the given named file.
func URLForFile(name string) string {
	m := md5.New()
	name = strings.Replace(name, " ", "_", -1)
	m.Write([]byte(name))
	h := hex.EncodeToString(m.Sum([]byte{}))

	return "http://upload.wikimedia.org/wikipedia/commons/" +
		string(h[0]) + "/" + h[0:2] + "/" + url.QueryEscape(name)
}
