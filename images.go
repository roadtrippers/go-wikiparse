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
var galleryRE *regexp.Regexp
var galleryFileRE *regexp.Regexp
var galleryImageRE *regexp.Regexp

func init() {
	fileRE = regexp.MustCompile(`\[File:([^\|\]]+)`)
	imageRE = regexp.MustCompile(`\[Image:([^\|\]]+)`)
	galleryRE = regexp.MustCompile(`(?mis)<gallery>.+</gallery>`)
	galleryFileRE = regexp.MustCompile(`File:([^\|\]]+)`)
	galleryImageRE = regexp.MustCompile(`Image:([^\|\]]+)`)
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
	galleryMatches := galleryRE.FindAllString(cleaned, 10000)

	rv := []string{}

	for _, x := range fileMatches {
		rv = append(rv, x[1])
	}

	for _, x := range imageMatches {
		rv = append(rv, x[1])
	}

	for _, x := range galleryMatches {
		// for each gallery, find the images and files after stripping comments
		noComments := commentRE.ReplaceAllString(x, "")

		galleryFileMatches := galleryFileRE.FindAllStringSubmatch(noComments, 10000)
		galleryImageMatches := galleryImageRE.FindAllStringSubmatch(noComments, 10000)

		for _, y := range galleryFileMatches {
			rv = append(rv, y[1])
		}

		for _, y := range galleryImageMatches {
			rv = append(rv, y[1])
		}
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
