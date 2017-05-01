package wikiparse

import (
	"testing"
)

type geoboxTestInput struct {
	input string
	templateType string
	attributes map[string]string
}

var geoboxTestData = []geoboxTestInput{
	geoboxTestInput{
		`{{Geobox|River
		&lt;!-- *** Heading *** --&gt;
		| name = Burrum
		| native_name =
		| other_name =
		| category = River
		&lt;!-- *** Names **** --&gt;
		| etymology = [[Kabi people|Kabi]]: ''rocks interrupting river flow''&lt;ref name=QPN/&gt;
		| nickname =
		&lt;!-- *** Image *** --&gt;
		| image =
		| image_caption =
		| image_size =
		&lt;!-- *** Country *** --&gt;
		| country = [[Australia]]
		| state_type = [[States and territories of Australia|State]]
		| state = [[Queensland]]
		| region_type = [[Regions of Queensland|Region]]
		| region = [[Wide Bay-Burnett]]
		| region1 =
		| district =
		| municipality =
		&lt;!-- *** Family *** --&gt;
		| parent =
		| tributary_left =
		| tributary_left1 =
		| tributary_left2 =
		| tributary_left3 =
		| tributary_left4 =
		| tributary_right =
		| tributary_right1 =
		| tributary_right2 =
		| tributary_right3 =
		| city =
		| landmark =
		&lt;!-- *** River locations *** --&gt;
		| source = [[Lenthalls Dam|Lake Lenthall]]
		| source_location = [[Lenthalls Dam|Lake Lenthall]]
		| source_region =
		| source_country =
		| source_elevation = 24
		| source_lat_d =
		| source_lat_m =
		| source_lat_s =
		| source_lat_NS =
		| source_long_d =
		| source_long_m =
		| source_long_s =
		| source_long_EW =

		| source1 =
		| source1_location =
		| source1_region =
		| source1_country =
		| source1_elevation =
		| source1_lat_d =
		| source1_lat_m =
		| source1_lat_s =
		| source1_lat_NS =
		| source1_long_d =
		| source1_long_m =
		| source1_long_s =
		| source1_long_EW =

		| source_confluence =
		| source_confluence_location =
		| source_confluence_region =
		| source_confluence_country =
		| source_confluence_elevation =
		| source_confluence_lat_d =
		| source_confluence_lat_m =
		| source_confluence_lat_s =
		| source_confluence_lat_NS =
		| source_confluence_long_d =
		| source_confluence_long_m =
		| source_confluence_long_s =
		| source_confluence_long_EW =

		| mouth = {{QLDcity|Hervey Bay}}, [[Coral Sea]]
		| mouth_location = {{QLDcity|[[Burrum Heads]]}}
		| mouth_region =
		| mouth_country =
		| mouth_elevation = 0
		| mouth_lat_d = 25
		| mouth_lat_m = 10
		| mouth_lat_s = 46
		| mouth_lat_NS = S
		| mouth_long_d = 152
		| mouth_long_m = 37
		| mouth_long_s = 01
		| mouth_long_EW = E
		&lt;!-- *** Dimensions *** --&gt;
		| length = 31
		| width =
		| depth =
		| volume =
		| watershed = 3371
		| discharge =
		| discharge_location =
		| discharge_max =
		| discharge_min =
		&lt;!-- *** Free fields *** --&gt;
		| free = [[Burrum Coast National Park]]
		| free_type = [[National park]]
		| free1 = [[Lenthalls Dam|Lake Lenthall]]
		| free1_type = [[Reservoir]]
		&lt;!-- *** Maps *** --&gt;
		| pushpin_map = Australia Queensland
		| pushpin_map_relief = 1
		| pushpin_map_caption = Location of Burrum River [[river mouth|mouth]] in Queensland
		&lt;!-- *** Website *** --&gt;
		| website =
		| commons =
		&lt;!-- *** Footnotes *** --&gt;
		| footnotes =&lt;ref name=bonzle&gt;{{cite web|url=http://www.bonzle.com/c/a?a=p&amp;p=203255&amp;cmd=sp|title=Map of Burrum River, QLD|accessdate=23 June 2015|work=Bonzle Digital Atlas of Australia}}&lt;/ref&gt;
		}}`,
		"River",
		map[string]string{
			"name": "Burrum",
			"native_name": "",
			"other_name": "",
			"category": "River",
			"etymology": "[[Kabi people|Kabi]]: ''rocks interrupting river flow''&lt;ref name=QPN/&gt;",
			"nickname": "",
			"image": "",
			"image_caption": "",
			"image_size":"",
			"country": "[[Australia]]",
			"state_type": "[[States and territories of Australia|State]]",
			"state": "[[Queensland]]",
			"region_type": "[[Regions of Queensland|Region]]",
			"region": "[[Wide Bay-Burnett]]",
			"region1": "",
			"district": "",
			"municipality": "",
			"parent": "",
			"tributary_left": "",
			"tributary_left1": "",
			"tributary_left2": "",
			"tributary_left3": "",
			"tributary_left4": "",
			"tributary_right": "",
			"tributary_right1": "",
			"tributary_right2": "",
			"tributary_right3": "",
			"city": "",
			"landmark": "",
			"source": "[[Lenthalls Dam|Lake Lenthall]]",
			"source_location": "[[Lenthalls Dam|Lake Lenthall]]",
			"source_region": "",
			"source_country":"",
			"source_elevation": "24",
			"source_lat_d": "",
			"source_lat_m": "",
			"source_lat_s": "",
			"source_lat_NS": "",
			"source_long_d": "",
			"source_long_m": "",
			"source_long_s": "",
			"source_long_EW":"",
			"source1": "",
			"source1_location": "",
			"source1_region": "",
			"source1_country": "",
			"source1_elevation": "",
			"source1_lat_d": "",
			"source1_lat_m": "",
			"source1_lat_s": "",
			"source1_lat_NS":"",
			"source1_long_d": "",
			"source1_long_m": "",
			"source1_long_s": "",
			"source1_long_EW":"",
			"source_confluence": "",
			"source_confluence_location": "",
			"source_confluence_region": "",
			"source_confluence_country":"",
			"source_confluence_elevation": "",
			"source_confluence_lat_d": "",
			"source_confluence_lat_m": "",
			"source_confluence_lat_s": "",
			"source_confluence_lat_NS":"",
			"source_confluence_long_d": "",
			"source_confluence_long_m": "",
			"source_confluence_long_s": "",
			"source_confluence_long_EW":"",
			"mouth": "{{QLDcity|Hervey Bay}}, [[Coral Sea]]",
			"mouth_location": "{{QLDcity|[[Burrum Heads]]}}",
			"mouth_region":"",
			"mouth_country":"",
			"mouth_elevation": "0",
			"mouth_lat_d": "25",
			"mouth_lat_m": "10",
			"mouth_lat_s": "46",
			"mouth_lat_NS": "S",
			"mouth_long_d": "152",
			"mouth_long_m": "37",
			"mouth_long_s": "01",
			"mouth_long_EW": "E",
			"length": "31",
			"width": "",
			"depth": "",
			"volume":"",
			"watershed": "3371",
			"discharge": "",
			"discharge_location": "",
			"discharge_max":"",
			"discharge_min":"",
			"free": "[[Burrum Coast National Park]]",
			"free_type": "[[National park]]",
			"free1": "[[Lenthalls Dam|Lake Lenthall]]",
			"free1_type": "[[Reservoir]]",
			"pushpin_map": "Australia Queensland",
			"pushpin_map_relief": "1",
			"pushpin_map_caption": "Location of Burrum River [[river mouth|mouth]] in Queensland",
			"website":"",
			"commons":"",
			"footnotes": "&lt;ref name=bonzle&gt;{{cite web|url=http://www.bonzle.com/c/a?a=p&amp;p=203255&amp;cmd=sp|title=Map of Burrum River, QLD|accessdate=23 June 2015|work=Bonzle Digital Atlas of Australia}}&lt;/ref&gt;",
		},
	},
	geoboxTestInput{
		`      <text xml:space="preserve">{{Geobox|Valley
|name=Cottonwood Valley (Arizona/Nevada)
|native_name=
|image= Lake_Mohave_from_Spirit_Mountain_2.jpg
|image_size=295px
|image_caption=View slightly east of true-north. The expanse of the valley on left is Nevada. The view is from the peak [[Spirit Mountain (Nevada)]] of the [[Newberry Mountains (Nevada)|Newberry Mountains]], about 22 mi distant from the north end of [[Lake Mohave]].
|country = [[United States]]
|country_flag=1
|country_flag_type=1

|state = [[Arizona]]
|state1 = [[California]]
&lt;!--
|state_flag=1--&gt;

|region=(southeast)-'''[[Mojave Desert]]'''
&lt;!--|region_type=Regions--&gt;

|district=[[Mohave County, Arizona]]
|district1=[[Clark County, Nevada]]
|district2=
|district_type=County
|municipality=
| part =
| city = [[Cottonwood Cove, Nevada|Cottonwood Cove, NV]]
| city_type=Communities
| landmark =
| river = [[Colorado River]] &amp; [[Lake Mohave]]

&lt;!--Lake Mohave (coordinates, valley-center ~~slightly south of lake center) 35,26,14 N, 114,38,37 W--&gt;
&lt;!--Whitlock Cienega-(Whitlock Valley), coord, 32,33,10-N, 109,20,7-W--&gt;
| lat_d = 35 | lat_m = 26 | lat_s = 14 | lat_NS = N
| long_d = 114  | long_m = 38 | long_s = 37 | long_EW = W
| topo =
| topo_map =
| topo_maker =

| geology=
| border=[[Whitlock Mountains]]-W&lt;br&gt;[[Black Hills (Greenlee County)]]-N&lt;br&gt;[[Peloncillo Mountains (Cochise County)|Peloncillo Mountains]]-E&lt;br&gt;[[San Simon Valley]]-SW &amp; S
| orogeny=
| length_imperial=20| length_orientation=N-S
| width_imperial=16 | width_orientation= E-W

| highest            =
| highest_location   =
| highest_country    =
| highest_state      =
| highest_region     =
| highest_district   =
| highest_elevation  =
| highest_elevation_imperial=
| highest_lat_d      =
| highest_lat_m      =
| highest_lat_s      =
| highest_lat_NS     = N
| highest_long_d     =
| highest_long_m     =
| highest_long_s     =
| highest_long_EW    = W

&lt;!--Parks Lake, coord, 32.5678-N, 109.3031-W--&gt;
| lowest             = [[Lake Mohave]]
| lowest_location   = probably slightly south of lake center
| lowest_country    =
| lowest_state      =
| lowest_region     =
| lowest_district   =
| lowest_elevation  =
| lowest_elevation_imperial=
| lowest_lat_d      =
| lowest_lat_m      =

| lowest_lat_s      =
| lowest_lat_NS     = N
| lowest_long_d     =
| lowest_long_m     =
| lowest_long_s     =
| lowest_long_EW    = W
| map =USA Arizona location map.svg
| map_caption = Cottonwood Valley (Arizona/Nevada) &lt;br&gt; (in Arizona)
| map_background =
| map_locator = Arizona svg&lt;!-- *** Website *** --&gt;
| website =
&lt;!-- *** Footnotes *** --&gt;
| footnotes =
}}`,
		"Valley",
		map[string]string{
			"name":"Cottonwood Valley (Arizona/Nevada)",
			"native_name":"",
			"image":"Lake_Mohave_from_Spirit_Mountain_2.jpg",
			"image_size":"295px",
			"image_caption":"View slightly east of true-north. The expanse of the valley on left is Nevada. The view is from the peak [[Spirit Mountain (Nevada)]] of the [[Newberry Mountains (Nevada)|Newberry Mountains]], about 22 mi distant from the north end of [[Lake Mohave]].",
			"country" :"[[United States]]",
			"country_flag":"1",
			"country_flag_type":"1",
			"state" :"[[Arizona]]",
			"state1" :"[[California]]",
			"region":"(southeast)-'''[[Mojave Desert]]'''",
			"district":"[[Mohave County, Arizona]]",
			"district1":"[[Clark County, Nevada]]",
			"district2":"",
			"district_type":"County",
			"municipality":"",
			"part" :"",
			"city" :"[[Cottonwood Cove, Nevada|Cottonwood Cove, NV]]",
			"city_type":"Communities",
			"landmark" :"",
			"river" :"[[Colorado River]] &amp; [[Lake Mohave]]",
			"lat_d" :"35",
			"lat_m" :"26",
			"lat_s" :"14",
			"lat_NS" :"N",
			"long_d" :"114",
			"long_m" :"38",
			"long_s" :"37",
			"long_EW" :"W",
			"topo" :"",
			"topo_map" :"",
			"topo_maker" :"",
			"geology":"",
			"border":"[[Whitlock Mountains]]-W&lt;br&gt;[[Black Hills (Greenlee County)]]-N&lt;br&gt;[[Peloncillo Mountains (Cochise County)|Peloncillo Mountains]]-E&lt;br&gt;[[San Simon Valley]]-SW &amp; S",
			"orogeny":"",
			"length_imperial":"20",
			"length_orientation":"N-S",
			"width_imperial":"16",
			"width_orientation":"E-W",
			"highest"            :"",
			"highest_location"   :"",
			"highest_country"    :"",
			"highest_state"      :"",
			"highest_region"     :"",
			"highest_district"   :"",
			"highest_elevation"  :"",
			"highest_elevation_imperial":"",
			"highest_lat_d"      :"",
			"highest_lat_m"      :"",
			"highest_lat_s"      :"",
			"highest_lat_NS"     :"N",
			"highest_long_d"     :"",
			"highest_long_m"     :"",
			"highest_long_s"     :"",
			"highest_long_EW"    :"W",
			"lowest"             :"[[Lake Mohave]]",
			"lowest_location"   :"probably slightly south of lake center",
			"lowest_country"    :"",
			"lowest_state"      :"",
			"lowest_region"     :"",
			"lowest_district"   :"",
			"lowest_elevation"  :"",
			"lowest_elevation_imperial":"",
			"lowest_lat_d"      :"",
			"lowest_lat_m"      :"",
			"lowest_lat_s"      :"",
			"lowest_lat_NS"     :"N",
			"lowest_long_d"     :"",
			"lowest_long_m"     :"",
			"lowest_long_s"     :"",
			"lowest_long_EW"    :"W",
			"map" :"USA Arizona location map.svg",
			"map_caption" :"Cottonwood Valley (Arizona/Nevada) &lt;br&gt; (in Arizona)",
			"map_background" :"",
			"map_locator" :"Arizona svg",
			"website" :"",
			"footnotes" :"",
		},
	},
}

func testOneParseGeobox(t *testing.T, ti geoboxTestInput) {
	res, err := ParseGeobox(ti.input)

	if err != nil {
		t.Fatalf("Unexpected error on %v: %v\n", ti.input, err)
	} else if ti.templateType != res.TemplateType {
		t.Fatalf("Expected type of %s\nGot %s\n", ti.templateType, res.TemplateType)
	} else {
		for k, v := range ti.attributes {
			val, ok := res.Attributes[k]
			switch {
			case !ok:
				t.Fatalf("Expected to have attribute %s, but it was missing", k)
			case v != val :
				t.Fatalf("Expected attribute %s to have value %s\nGot %s\n", k, v, val)
			}
		}
	}


}

func TestParseGeobox(t *testing.T) {
	t.Parallel()
	for _, ti := range geoboxTestData {
		testOneParseGeobox(t, ti)
	}
}
