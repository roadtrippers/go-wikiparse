package wikiparse

import (
	"regexp"
	"errors"
	"strings"
	"fmt"
)

type Convert struct {
	InputValues []string
	InputUnits []string
	OutputUnits []string
	Options []string
}

var convertStartRE *regexp.Regexp
var bracketReplacer *strings.Replacer

var rangeIndicators = map[string]string{
	"-": "–",
	"–": "–",
	"&ndash;": "–",
	"and": "and",
	"&": "and",
	"and(-)": "–",
	"or": "or",
	"to": "to",
	"to(-)": "to",
	"to about": "to about",
	"+/-": "±",
	"±": "±",
	"&plusmin;": "±",
	"+": "+",
	",": ",",
	", and": ", and",
	", or": ", or",
	"by": "by",
	"x": "×",
	"×": "×",
	"&times;": "×",
}

var unitCodes = map[string]string{
	"km2": "km2",
	"m2": "m2",
	"cm2": "cm2",
	"mm2": "mm2",
	"ha": "ha",
	"sqmi": "sq mi",
	"acre": "acre",
	"sqyd": "sq yd",
	"sqft": "sq ft",
	"sqfoot": "sq ft",
	"sqin": "sq in",
	"sqnmi": "sq nmi",
	"dunam": "dunam",
	"tsubo": "tsubo",
	"kg/m3": "kg/m3",
	"g/m3": "g/m3",
	"lb/ft3": "lb/cu ft",
	"lb/yd3": "lb/cu yd",
	"GJ": "GJ",
	"MJ": "MJ",
	"kJ": "kJ",
	"hJ": "hJ",
	"daJ": "daJ",
	"J": "J",
	"dJ": "dJ",
	"cJ": "cJ",
	"mJ": "mJ",
	"µJ": "µJ",
	"uJ": "µJ",
	"nJ": "nJ",
	"Merg": "Merg",
	"kerg": "kerg",
	"erg": "erg",
	"TWh": "TWh",
	"TW.h": "TW·h",
	"GWh": "GWh",
	"GW.h": "GW·h",
	"MWh": "MWh",
	"MW.h": "MW·h",
	"kWh": "kWh",
	"kW.h": "kW·h",
	"Wh": "Wh",
	"W.h": "W·h",
	"GeV": "GeV",
	"MeV": "MeV",
	"keV": "keV",
	"eV": "eV",
	"meV": "meV",
	"Cal": "Cal",
	"Mcal": "Mcal",
	"kcal": "kcal",
	"cal": "cal",
	"mcal": "mcal",
	"ftpdl": "ft·pdl",
	"ftlbf": "ft·lbf",
	"ftlb-f": "ft·lbf",
	"inlbf": "in·lbf",
	"in-lbf": "in·lbf",
	"inozf": "in·ozf",
	"inoz-f": "in·ozf",
	"hph": "hph",
	"Btu": "Btu",
	"BTU": "BTU",
	"GtTNT": "gigatonne of TNT",
	"GtonTNT": "gigatonne of TNT",
	"MtTNT": "megatonne of TNT",
	"MtonTNT": "megatonne of TNT",
	"kTNT": "kilotonne of TNT",
	"ktonTNT": "kilotonne of TNT",
	"tTNT": "tonne of TNT",
	"tonTNT": "tonne of TNT",
	"Eh": "Eh",
	"Ry": "Ry",
	"toe": "toe",
	"BOE": "BOE",
	"cuftnaturalgas": "cubic foot of natural gas",
	"cufootnaturalgas": "cubic foot of natural gas",
	"latm": "l·atm",
	"Latm": "l·atm",
	"impgalatm": "imp gal·atm",
	"USgalatm": "US gal·atm",
	"usgalatm": "US gal·atm",
	"U.S.galatm": "US gal·atm",
	"GN": "GN",
	"MN": "MN",
	"kN": "kN",
	"N": "N",
	"mN": "mN",
	"µN": "µN",
	"uN": "µN",
	"nN": "nN",
	"Mdyn": "Mdyn",
	"kdyn": "kdyn",
	"dyn": "dyn",
	"dyne": "dyne",
	"mdyn": "mdyn",
	"t-f": "tf",
	"tf": "tf",
	"kg-f": "kgf",
	"kgf": "kgf",
	"g-f": "gf",
	"gf": "gf",
	"mg-f": "mgf",
	"mgf": "mgf",
	"pdl": "pdl",
	"LT-f": "LTf",
	"LTf": "LTf",
	"ST-f": "STf",
	"STf": "STf",
	"lb-f": "lbf",
	"lbf": "lbf",
	"gr-f": "grf",
	"grf": "grf",
	"Mm": "Mm",
	"km": "km",
	"m": "m",
	"cm": "cm",
	"mm": "mm",
	"µm": "µm",
	"um": "µm",
	"nm": "nm",
	"Å": "Å",
	"mi": "mi",
	"furlong": "furlong",
	"chain": "chain",
	"rd": "rd",
	"fathom": "fathom",
	"yd": "yd",
	"foot": "ft",
	"ft": "ft",
	"in": "in",
	"nmi": "nmi",
	"pc": "pc",
	"ly": "ly",
	"AU": "AU",
	"kg": "kg",
	"g": "g",
	"mg": "mg",
	"µg": "µg",
	"ug": "µg",
	"t": "t",
	"MT": "MT",
	"LT": "long ton",
	"long ton": "long ton",
	"ST": "short ton",
	"short ton": "short ton",
	"st": "st",
	"lb": "lb",
	"oz": "oz",
	"drachm": "drachm",
	"dram": "dram",
	"gr": "gr",
	"ozt": "ozt",
	"carat": "carat",
	"m/s": "m/s",
	"km/h": "km/h",
	"mph": "mph",
	"ft/s": "ft/s",
	"foot/s": "ft/s",
	"kn": "kn",
	"knot": "kn",
	"K": "K",
	"°C": "°C",
	"C": "°C",
	"°R": "°R",
	"R": "°R",
	"°F": "°F",
	"F": "°F",
	"C-change": "°C",
	"F-change": "°F",
	"Nm": "Nm",
	"N.m": "N·m",
	"kg.m": "kg·m",
	"lb.ft": "lb·ft",
	"kgf.m": "kgf·m",
	"lbf.ft": "lbf·ft",
	"m3": "m3",
	"cm3": "cm3",
	"cc": "cm3",
	"mm3": "mm3",
	"kl": "kl",
	"kL": "kL",
	"l": "l",
	"L": "L",
	"cl": "cl",
	"cL": "cL",
	"ml": "ml",
	"mL": "mL",
	"cuyd": "cu yd",
	"cuft": "cu ft",
	"cufoot": "cu ft",
	"cuin": "cu in",
	"impbbl": "imp bbl",
	"impbsh": "imp bsh",
	"impbu": "imp bsh",
	"impgal": "imp gal",
	"impqt": "imp qt",
	"imppt": "imp pt",
	"impoz": "imp fl oz",
	"impfloz": "imp fl oz",
	"USbbl": "US bbl",
	"U.S.bbl": "US bbl",
	"oilbbl": "bbl",
	"USbeerbbl": "US bbl",
	"usbeerbbl": "US bbl",
	"U.S.beerbbl": "US bbl",
	"USgal": "US gal",
	"U.S.gal": "US gal",
	"USqt": "US qt",
	"U.S.qt": "US qt",
	"USpt": "US pt",
	"U.S.pt": "US pt",
	"USoz": "US fl oz",
	"USfloz": "US fl oz",
	"U.S.oz": "US fl oz",
	"U.S.floz": "US fl oz",
	"USdrybbl": "US dry bbl",
	"U.S.drybbl": "US dry bbl",
	"USbsh": "US bsh",
	"U.S.bsh": "US bsh",
	"USbu": "US bu",
	"U.S.bu": "US bu",
	"USdrygal": "US dry gal",
	"U.S.drygal": "US dry gal",
	"USdryqt": "US dry qt",
	"U.S.dryqt": "US dry qt",
	"USdrypt": "US dry pt",
	"U.S.drypt": "US dry pt",
	"GPa": "GPa",
	"MPa": "MPa",
	"kPa": "kPa",
	"hPa": "hPa",
	"Pa": "Pa",
	"mPa": "mPa",
	"mbar": "mbar",
	"mb": "mbar",
	"dbar": "dbar",
	"bar": "bar",
	"kBa": "kBa",
	"Ba": "Ba",
	"atm": "atm",
	"Torr": "Torr",
	"mmHg": "mmHg",
	"inHg": "inHg",
	"psi": "psi",
	"km/l": "km/l",
	"km/L": "km/L",
	"l/100 km": "l/100 km",
	"L/100 km": "L/100 km",
	"l/km": "l/km",
	"L/km": "L/km",
	"mpgimp": "mpgimp",
	"mpgus": "mpgus",
	"mpgUS": "mpgUS",
	"mpgU.S.": "mpgU.S.",
	"impgal/mi": "impgal/mi",
	"usgal/mi": "usgal/mi",
	"USgal/mi": "USgal/mi",
	"U.S.gal/mi": "U.S.gal/mi",
	"PD/sqkm": "PD/sqkm",
	"/sqkm": "/sqkm",
	"PD/ha": "PD/ha",
	"/ha": "/ha",
	"PD/sqmi": "PD/sqmi",
	"/sqmi": "/sqmi",
	"PD/acre": "PD/acre",
	"/acre": "/acre",
	"$/lb": "$/lb",
	"$/kg": "$/kg",
	"$/ozt": "$/ozt",
	"miydftin": "mi yd ft it",
	"mift": "mi ft",
	"ydftin": "yd ft in",
	"ydft": "yd ft",
	"ftin": "ft in",
	"footin": "foot in",
	"handin": "hand in",
	"lboz": "lb oz",
	"stlb": "st lb",
	"stlboz": "st lb oz",
	"st and lb": "st lb",

}

func init() {
	convertStartRE = regexp.MustCompile(`(?mi){{(convert|cvt)`)
	bracketReplacer = strings.NewReplacer("}}", "", "{{", "")
}

func IsConvert(text string) bool {
	return convertStartRE.MatchString(text)
}

func ParseConvert(text string) (*Convert, error) {
	results := Convert{
		InputValues: make([]string, 0),
		InputUnits: make([]string, 0),
		OutputUnits: make([]string, 0),
		Options: make([]string, 0),
	}

	cleaned := nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), "")
	parts := strings.Split(cleaned, "|")

	for i, v := range parts {
		parts[i] = bracketReplacer.Replace(v)
	}

	if !IsConvert(cleaned) || len(parts) == 0 {
		return nil, errors.New("No Convert found")
	}

	// cvt variant always implies abbr=on
	if parts[0] == "cvt" {
		results.Options = append(results.Options, "abbr=on")
	}

	inRange := false
	gotInputUnit := false

	for _, v := range parts[1:] {
		// split the value; output formats can contain multiple units. ranges can, however, include the
		// "+"; since we don't use splitV for anything but units we are safe
		splitV := strings.Split(v, "+")
		if len(splitV) < 2 {
			splitV = strings.Split(v, " ")
		}

		if unit, ok := unitCodes[splitV[0]]; ok {
			// we matched a unit... is it input, or is it output?
			if gotInputUnit {
				// Output can contain multiple units, so add each one
				for _, ut := range splitV {
					results.OutputUnits = append(results.OutputUnits, unitCodes[ut])
				}
			} else {
				// Inputs only contain a single unit at a time
				results.InputUnits = append(results.InputUnits, unit)
				gotInputUnit = true
			}
		} else if strings.Contains(v, "=") {
			// we matched options
			results.Options = append(results.Options, v)
		} else if rng, ok := rangeIndicators[v]; ok {
			// we matched a range indicator
			results.InputValues[0] = fmt.Sprintf("%s %s", results.InputValues[0], rng)
			inRange = true
		} else {
			// assume its an input value
			if inRange {
				// ranges get merged into a single input
				results.InputValues[0] = fmt.Sprintf("%s %s", results.InputValues[0], v)
				inRange = false
			} else {
				results.InputValues = append(results.InputValues, v)
			}
			gotInputUnit = false
		}
	}

	return &results, nil
}
