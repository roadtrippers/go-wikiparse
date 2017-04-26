package wikiparse

import (
	"regexp"
	"strconv"
)

var altCoordsRE *regexp.Regexp

func init() {
	altCoordsRE = regexp.MustCompile(`(?mi){{(coords|coordinates)`)
}

func IsAltCoords(text string) bool {
	return altCoordsRE.MatchString(text)
}

func ParseAltCoords(text string) (Coord, error) {
	parts := partsFromText(text)

	if !IsAltCoords(text) || len(parts) < 3 {
		return Coord{}, ErrNoCoordFound
	}

	var err error
	result := Coord{}

	if result.Lat, err = strconv.ParseFloat(parts[1], 64); err != nil {
		return Coord{}, err
	}

	if result.Lon, err = strconv.ParseFloat(parts[2], 64); err != nil {
		return Coord{}, err
	}

	return result, nil
}
