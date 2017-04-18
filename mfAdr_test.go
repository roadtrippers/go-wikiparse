package wikiparse

import "testing"

type testMFAdrInput struct {
	input string
	street string
	city string
	region string
	pocode string
	nation string
}

var testMFAdrData = []testMFAdrInput{
	testMFAdrInput{
		`{{mf-adr
|street = 1255 Kingston Rd
|city   = Pickering
|region = Ontario
|pocode = L1V 1B5
|nation = Canada
}}`,
		"1255 Kingston Rd",
		"Pickering",
		"Ontario",
		"L1V 1B5",
		"Canada",
	},
}

func testOneMFAdr(t *testing.T, ti testMFAdrInput) {
	mfadr, err := ParseMFAdr(ti.input)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if mfadr.Street != ti.street {
		t.Fatalf("Expected street %v, got %v", ti.street, mfadr.Street)
	}

	if mfadr.City != ti.city {
		t.Fatalf("Expected city %v, got %v", ti.city, mfadr.City)
	}

	if mfadr.Region != ti.region {
		t.Fatalf("Expected region %v, got %v", ti.region, mfadr.Region)
	}

	if mfadr.PostalCode != ti.pocode {
		t.Fatalf("Expected pocode %v, got %v", ti.pocode, mfadr.PostalCode)
	}

	if mfadr.Nation != ti.nation {
		t.Fatalf("Expected nation %v, got %v", ti.nation, mfadr.Nation)
	}
}

func TestMFAdrs(t *testing.T) {
	t.Parallel()
	for _, ti := range testMFAdrData {
		testOneMFAdr(t, ti)
	}
}
