package wikiparse

import "testing"

func TestParseURL(t *testing.T) {
	template := "{{URL|http://www.2pc.org|2pc.org}}"
	result, err := ParseTemplate(template)

	urlRes := result.(*URLData)

	if err != nil {
		t.Fatalf("Unexpected error on %v: %v", template, err)
	} else if urlRes.Url != "http://www.2pc.org" {
		t.Fatalf("Expected url to be %v\nGot %v\n", "http://www.2pc.org", urlRes.Url)
	} else if urlRes.DisplayAs != "2pc.org" {
		t.Fatalf("Expected url to be %v\nGot %v\n", "2pc.org", urlRes.DisplayAs)
	}
}
