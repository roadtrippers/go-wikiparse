package wikiparse

import "testing"

type urlTestInput struct {
	input string
	theUrl string
	theDisplayUrl string
}

var urlTestData = []urlTestInput{
	urlTestInput{
		"{{ URL | EXAMPLE.com }}",
		"EXAMPLE.com",
		"example.com",
	},
	urlTestInput{
		"{{ URL | example.com }}",
		"example.com",
		"example.com",
	},
	urlTestInput{
		"{{ URL | http://example.com }}",
		"http://example.com",
		"example.com",
	},
	urlTestInput{
		"{{ URL | www.example.com }}",
		"www.example.com",
		"www.example.com",
	},
	urlTestInput{
		"{{ URL | http://www.example.com }}",
		"http://www.example.com",
		"www.example.com",
	},
	urlTestInput{
		"{{ URL | https://www.example.com }}",
		"https://www.example.com",
		"www.example.com",
	},
	urlTestInput{
		"{{ URL | ftp://www.example.com }}",
		"ftp://www.example.com",
		"www.example.com",
	},
	urlTestInput{
		"{{ URL | ftp://ftp.example.com }}",
		"ftp://ftp.example.com",
		"ftp.example.com",
	},
	urlTestInput{
		"{{ URL | http://www.example.com/ }}",
		"http://www.example.com/",
		"www.example.com",
	},
	urlTestInput{
		"{{ URL | http://www.example.com/path }}",
		"http://www.example.com/path",
		"www.example.com/path",
	},
	urlTestInput{
		"{{ URL | irc://irc.example.com/channel }}",
		"irc://irc.example.com/channel",
		"irc.example.com/channel",
	},
	urlTestInput{
		"{{ URL | www.example.com/foo }}",
		"www.example.com/foo",
		"www.example.com/foo",
	},
	urlTestInput{
		"{{ URL | http://www.example.com/path/ }}",
		"http://www.example.com/path/",
		"www.example.com/path/",
	},
	urlTestInput{
		"{{ URL | www.example.com/foo/ }}",
		"www.example.com/foo/",
		"www.example.com/foo/",
	},
	urlTestInput{
		"{{ URL | 1=http://www.example.com/path?section=17 }}",
		"http://www.example.com/path?section=17",
		"www.example.com/path?section=17",
	},
	urlTestInput{
		"{{ URL | 1=www.example.com/foo?page=42 }}",
		"www.example.com/foo?page=42",
		"www.example.com/foo?page=42",
	},
	urlTestInput{
		"{{ URL | www.example.com | example.com }}",
		"www.example.com",
		"example.com",
	},
}

func testOneUrl(t *testing.T, ti urlTestInput) {
	urlData, err := ParseURL(ti.input)

	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", ti.input, err)
	} else if ti.theUrl != urlData.Url {
		t.Fatalf("Expected url %v\nGot %v", ti.theUrl, urlData.Url)
	} else if ti.theDisplayUrl != urlData.DisplayAs {
		t.Fatalf("Expected display url %v\nGot %v", ti.theDisplayUrl, urlData.DisplayAs)
	}
}

func TestAllUrls(t *testing.T) {
	t.Parallel()

	for _, ti := range urlTestData {
		testOneUrl(t, ti)
	}
}
