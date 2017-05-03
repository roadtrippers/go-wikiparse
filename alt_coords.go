package wikiparse

import (
	"regexp"
	"math"
	"fmt"
)

var altCoordsRE *regexp.Regexp

func init() {
	altCoordsRE = regexp.MustCompile(`(?mi){{(coords|coordinates)`)
}

func IsAltCoords(text string) bool {
	return altCoordsRE.MatchString(text)
}

func ParseAltCoords(text string) (Coord, error) {
	parts := partsFromText(text, altCoordsRE)

	if !IsAltCoords(text) || len(parts) < 3 {
		return Coord{}, ErrNoCoordFound
	}

	parts = cleanCoordParts(parts[1:])

	result, err := parseSexagesimal(parts)
	if err != nil {
		result, err = parseFloat(parts)
	}

	if math.Abs(result.Lat) > 90 {
		return result, fmt.Errorf("invalid latitude: %v", result.Lat)
	}
	if math.Abs(result.Lon) > 180 {
		return result, fmt.Errorf("invalid longitude: %v", result.Lon)
	}

	return result, err
}
