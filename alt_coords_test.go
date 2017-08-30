package wikiparse

import "testing"

type altCoordsTestData struct {
	input string
	lat float64
	lon float64
}

var altCoordsTestInput = []altCoordsTestData{
	altCoordsTestData{
		"{{coords|35.0824099|-106.6764794}}",
		35.0824099,
		-106.6764794,
	},
	altCoordsTestData{
		`{{coords|11.937227|N|103.485260|E|region:KH|notes=<ref name="odc">{{cite web | url=http://www.opendevelopmentcambodia.net/natural_resource/protected-areas/?id=63&cat=1&map | title=Open Development Cambodia: Central Cardamom Mountains | accessdate=27 December 2013}}</ref>|display=inline, title}}`,
		11.937227,
		103.485260,
	},
}

func TestParseAltCoords(t *testing.T) {
	t.Parallel()

	for _, ti := range altCoordsTestInput {
		result, err := ParseAltCoords(ti.input)

		if err != nil {
			t.Fatalf("Unexpected error on %s: %v\n", ti.input, err)
		} else if result.Lat != ti.lat {
			t.Fatalf("Expected latitude %f\nGot %f", ti.lat, result.Lat)
		} else if result.Lon != ti.lon {
			t.Fatalf("Expected longitude %f\nGot %f", ti.lon, result.Lon)
		}
	}
}
