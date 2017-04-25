package wikiparse

import (
	"testing"
	"fmt"
)

type testInfoText struct {
	input string
	templateType string
	attributes map[string]string
	withoutInfobox string
}

var infoTestData = []testInfoText {
	testInfoText{
		`{{Infobox lake
 |name = Great Salt Lake
 |image = Great Salt Lake ISS 2003.jpg
 |caption = Satellite photo from August 2003 after five years of [[drought]], reaching near-record lows. Note the difference in colors between the northern and southern portions of the lake, the result of a railroad [[causeway]].
 |image_bathymetry =
 |caption_bathymetry =
 |location = [[Utah]], [[United States]]
 |coords = {{coord|41|10|N|112|35|W|region:US-UT_type:waterbody_scale:1000000|display=inline,title}}
 |type = [[Endorheic]], [[hypersaline lake|hypersaline]], generally 27% salinity
 |inflow = [[Bear River (Great Salt Lake)|Bear]], [[Jordan River (Utah)|Jordan]], [[Weber River|Weber]] rivers
 |outflow =
 |catchment = 21,500 sq mi (55,685 km²)
 |basin_countries = United States
 |length = 75 mi (120 km)
 |width = 28 mi (45 km)
 |area = 1,700 sq mi (4,400 km²)
 |depth = 16 ft (4.9 m), when lake is at average level
 |max-depth = 33 ft (10 m) average, high of {{convert|45|ft|m|abbr=on}} in 1987, low of {{convert|24|ft|m|abbr=on}} in 1963
 |volume = {{convert|15338693.6|acre.ft|km3|2|abbr=on|lk=in}}
 |residence_time =
 |shore =
 |elevation = historical average of 4,200 feet (1,283&nbsp;m), 4,190.3 feet (1,277&nbsp;m) as of 2016 July 7
 |islands = 8–15 (variable, see ''[[Great Salt Lake#Islands|Islands]]'')
 |cities = [[Salt Lake City|Salt Lake]] and [[Ogden, Utah|Ogden]] [[metropolitan area]]s.
}}`,
		"lake",
		map[string]string{
			"name" : "Great Salt Lake",
			"image" : "Great Salt Lake ISS 2003.jpg",
			"caption" : "Satellite photo from August 2003 after five years of [[drought]], reaching near-record lows. Note the difference in colors between the northern and southern portions of the lake, the result of a railroad [[causeway]].",
			"image_bathymetry" : "",
			"caption_bathymetry" : "",
			"location" : "[[Utah]], [[United States]]",
			"coords" : "{{coord|41|10|N|112|35|W|region:US-UT_type:waterbody_scale:1000000|display=inline,title}}",
			"type" : "[[Endorheic]], [[hypersaline lake|hypersaline]], generally 27% salinity",
			"inflow" : "[[Bear River (Great Salt Lake)|Bear]], [[Jordan River (Utah)|Jordan]], [[Weber River|Weber]] rivers",
			"outflow" : "",
			"catchment" : "21,500 sq mi (55,685 km²)",
			"basin_countries" : "United States",
			"length" : "75 mi (120 km)",
			"width" : "28 mi (45 km)",
			"area" : "1,700 sq mi (4,400 km²)",
			"depth" : "16 ft (4.9 m), when lake is at average level",
			"max-depth" : "33 ft (10 m) average, high of {{convert|45|ft|m|abbr=on}} in 1987, low of {{convert|24|ft|m|abbr=on}} in 1963",
			"volume" : "{{convert|15338693.6|acre.ft|km3|2|abbr=on|lk=in}}",
			"residence_time" : "",
			"shore" : "",
			"elevation" : "historical average of 4,200 feet (1,283&nbsp;m), 4,190.3 feet (1,277&nbsp;m) as of 2016 July 7",
			"islands" : "8–15 (variable, see ''[[Great Salt Lake#Islands|Islands]]'')",
			"cities" : "[[Salt Lake City|Salt Lake]] and [[Ogden, Utah|Ogden]] [[metropolitan area]]s.",
		},
		"",
	},
	testInfoText{
		`{{Infobox zoo
|zoo_name=Cincinnati Zoo
|image=Cincinnati Zoo.jpg
|location=3400 Vine St, [[Cincinnati]], [[Ohio]], U.S.
|coordinates={{Coord|39.145|N|84.508|W|display=inline,title|source:nlwiki}}
|area={{Convert|75|acre}}&lt;ref&gt;{{cite book |url=https://books.google.com/books?id=-I6TgnxWnj8C&amp;pg=PA491 |title=Frommer's USA |publisher=John Wiley &amp; Sons |date=Feb 17, 2009 |accessdate=2013-05-09 |author=Baird, David |pages=491|display-authors=etal}}&lt;/ref&gt;
|date_opened=1875&lt;ref name=&quot;zoo_about&quot;/&gt;
|num_animals= 1,896
|num_species=500+&lt;ref name=&quot;zoo_about&quot;/&gt;
|annual_visitors=1.2 million+&lt;ref name=&quot;zoo_about&quot;/&gt;
|members=[[Association of Zoos and Aquariums|AZA]],&lt;ref name=&quot;aza_list&quot;/&gt; [[World Association of Zoos and Aquariums|WAZA]]&lt;ref name=&quot;waza_list&quot;/&gt;
|website={{URL|http://www.cincinnatizoo.org}}
}}

The Cincinnati Zoo and Botanical Garden is the second-oldest [[zoo]] in the United States and is located in [[Cincinnati, Ohio|Cincinnati]], [[Ohio]]. It opened in 1875, just 14 months after the [[Philadelphia Zoo]] on July 1, 1874. The [[Cincinnati Zoo Historic Structures|Reptile House]] is the oldest zoo building in the United States, dating from 1875.
`,
		"zoo",
		map[string]string{
			"zoo_name": "Cincinnati Zoo",
			"image": "Cincinnati Zoo.jpg",
			"location": "3400 Vine St, [[Cincinnati]], [[Ohio]], U.S.",
			"coordinates": "{{Coord|39.145|N|84.508|W|display=inline,title|source:nlwiki}}",
			"area": "{{Convert|75|acre}}&lt;ref&gt;{{cite book |url=https://books.google.com/books?id=-I6TgnxWnj8C&amp;pg=PA491 |title=Frommer's USA |publisher=John Wiley &amp; Sons |date=Feb 17, 2009 |accessdate=2013-05-09 |author=Baird, David |pages=491|display-authors=etal}}&lt;/ref&gt;",
			"date_opened": "1875&lt;ref name=&quot;zoo_about&quot;/&gt;",
			"num_animals": "1,896",
			"num_species": "500+&lt;ref name=&quot;zoo_about&quot;/&gt;",
			"annual_visitors": "1.2 million+&lt;ref name=&quot;zoo_about&quot;/&gt;",
			"members": "[[Association of Zoos and Aquariums|AZA]],&lt;ref name=&quot;aza_list&quot;/&gt; [[World Association of Zoos and Aquariums|WAZA]]&lt;ref name=&quot;waza_list&quot;/&gt;",
			"website": "{{URL|http://www.cincinnatizoo.org}}",
		},
		`

The Cincinnati Zoo and Botanical Garden is the second-oldest [[zoo]] in the United States and is located in [[Cincinnati, Ohio|Cincinnati]], [[Ohio]]. It opened in 1875, just 14 months after the [[Philadelphia Zoo]] on July 1, 1874. The [[Cincinnati Zoo Historic Structures|Reptile House]] is the oldest zoo building in the United States, dating from 1875.
`,
	},
	testInfoText{
		`{{Infobox boxer
|name = Kid Chocolate
|nationality =   [[Cuba]]n
|realname = Eligio Sardiñas Montalvo
|image = Kid Chocolate.jpg
|nickname = The Cuban Bon Bon
| height = {{convert|1.68|m|ftin|abbr=on}}
| reach = {{convert|165|cm|in|abbr=on}}
|weight = [[Super Featherweight]]
|birth_date = January 6, 1910
|birth_place = [[Cerro, Havana|Cerro]], [[Havana]], [[Cuba]]
|death_date={{death date and age|1988|8|8|1910|1|6|mf=y}}
|death_place = [[Cuba]]
|style = Orthodox
|total = 152
|wins = 136
|KO = 51
|losses = 10
|draws = 6
|no contests =
}}
''For the boxer of the same nickname see [[Peter Quillin]].''

'''Eligio Sardiñas Montalvo''' (January 6, 1910 – August 8, 1988), better known as '''Kid Chocolate''', was a [[Cuba]]n [[boxing|boxer]] who enjoyed wild success both in the [[boxing ring]] and outside of it during the 1930s.

==Biography==
Sardiñas, also nicknamed ''The Cuban Bon Bon'', learned how to fight by watching old fight films in Cuba. He later sparred with boxers such as [[Benny Leonard]] and [[Jack Johnson (boxer)|Jack Johnson]], all world champions, before beginning an amateur [[boxing]] career. As an amateur, he allegedly won all 100 of his fights, 86 by [[knockout]], but this record was apparently fabricated for publicity purposes.

His professional boxing debut, officially, occurred on December 8, 1927, when he beat [[Johnny Cruz]] in six rounds in [[Havana]]. Although it has been claimed that he had 100 amateur fights and 21 KO wins as a pro in Cuba, this was a fabrication by his manager, Pincho Gutierrez.

Research by boxing historian Enrique Encinosa has uncovered 22 amateur bouts, verified through Cuban newspapers ''Diario de la Marina'' and ''La Noche'', as well as various books published by biographers or the Cuban government.

His first 9 bouts, including a five-round knockout win in a rematch with Cruz, were held in Cuba. In 1928, he moved to the [[United States]] and began campaigning in [[New York City]]. He won his first nine bouts there, five by knockout, and 12 of his first 13 fights in his new hometown. The only person to escape the ring without a defeat against Chocolate during that span was [[Joey Scalfaro]], who held him to a ten-round draw.

By 1929, Sardiñas was becoming a name to be reckoned with in boxing. He had 23 fights that year, and continued his undefeated run by winning each of them. He also began to box more competent opponents. Among the boxers he defeated were former world champion [[Fidel LaBarba]] (beaten by a decision in ten), future world champion [[Al Singer]] (also by a decision in ten), and fringe contenders [[Bushy Graham]], [[Vic Burrone]] and [[Gregorio Vidal]], all of whom, except for Graham, were beaten by decision. Graham was disqualified in the seventh round.`,
		"boxer",
		map[string]string{
			"name" : "Kid Chocolate",
			"nationality" : "[[Cuba]]n",
			"realname" : "Eligio Sardiñas Montalvo",
			"image" : "Kid Chocolate.jpg",
			"nickname" : "The Cuban Bon Bon",
			"height" : "{{convert|1.68|m|ftin|abbr=on}}",
			"reach" : "{{convert|165|cm|in|abbr=on}}",
			"weight" : "[[Super Featherweight]]",
			"birth_date" : "January 6, 1910",
			"birth_place" : "[[Cerro, Havana|Cerro]], [[Havana]], [[Cuba]]",
			"death_date" : "{{death date and age|1988|8|8|1910|1|6|mf=y}}",
			"death_place" : "[[Cuba]]",
			"style" : "Orthodox",
			"total" : "152",
			"wins" : "136",
			"KO" : "51",
			"losses" : "10",
			"draws" : "6",
			"no contests" : "",
		},
		`
''For the boxer of the same nickname see [[Peter Quillin]].''

'''Eligio Sardiñas Montalvo''' (January 6, 1910 – August 8, 1988), better known as '''Kid Chocolate''', was a [[Cuba]]n [[boxing|boxer]] who enjoyed wild success both in the [[boxing ring]] and outside of it during the 1930s.

==Biography==
Sardiñas, also nicknamed ''The Cuban Bon Bon'', learned how to fight by watching old fight films in Cuba. He later sparred with boxers such as [[Benny Leonard]] and [[Jack Johnson (boxer)|Jack Johnson]], all world champions, before beginning an amateur [[boxing]] career. As an amateur, he allegedly won all 100 of his fights, 86 by [[knockout]], but this record was apparently fabricated for publicity purposes.

His professional boxing debut, officially, occurred on December 8, 1927, when he beat [[Johnny Cruz]] in six rounds in [[Havana]]. Although it has been claimed that he had 100 amateur fights and 21 KO wins as a pro in Cuba, this was a fabrication by his manager, Pincho Gutierrez.

Research by boxing historian Enrique Encinosa has uncovered 22 amateur bouts, verified through Cuban newspapers ''Diario de la Marina'' and ''La Noche'', as well as various books published by biographers or the Cuban government.

His first 9 bouts, including a five-round knockout win in a rematch with Cruz, were held in Cuba. In 1928, he moved to the [[United States]] and began campaigning in [[New York City]]. He won his first nine bouts there, five by knockout, and 12 of his first 13 fights in his new hometown. The only person to escape the ring without a defeat against Chocolate during that span was [[Joey Scalfaro]], who held him to a ten-round draw.

By 1929, Sardiñas was becoming a name to be reckoned with in boxing. He had 23 fights that year, and continued his undefeated run by winning each of them. He also began to box more competent opponents. Among the boxers he defeated were former world champion [[Fidel LaBarba]] (beaten by a decision in ten), future world champion [[Al Singer]] (also by a decision in ten), and fringe contenders [[Bushy Graham]], [[Vic Burrone]] and [[Gregorio Vidal]], all of whom, except for Graham, were beaten by decision. Graham was disqualified in the seventh round.`,
	},
	testInfoText{
		`{{Primary sources|date=February 2014}}
{{Infobox zoo
|zoo_name=Monterey Bay Aquarium
|image=MontereyBayAquariumBackview.jpg
|image_caption=Back view of the aquarium (on the [[Pacific Ocean]])
|logo=Monterey Bay Aquarium Logo.svg
|logo_width=90px
|logo_caption=The aquarium's logo depicts [[Macrocystis pyrifera|giant kelp]].
|date_opened=October 20, 1984&lt;ref name=FAQ/&gt;
|location=[[Monterey, California]], US
|coordinates={{coord|36.618253|-121.901481|region:US_type:landmark|display=it}}
|area=
|floorspace=
|num_animals=35,000&lt;ref name=&quot;2005_press_kit&quot;/&gt;
|num_species=623 (plants and animals in 2005)&lt;ref name=&quot;2005_press_kit&quot;/&gt;
|largest_tank_vol={{Convert|1200000|gal|sing=on}}
|total_tank_vol=
|annual_visitors=2.08 Million&lt;ref&gt;{{cite news |last=Abel |first=David |url=https://www.bostonglobe.com/lifestyle/2016/08/02/aquariumside/aeNq1fhIiUd3i2vZlQ6RZJ/story.html |title=Top aquariums in the US, in terms of visitors |work=[[Boston Globe]] |date=2016-08-02 |accessdate=2016-12-02 }}&lt;/ref&gt;
|members=[[Association of Zoos and Aquariums|AZA]]&lt;ref name=&quot;aza_list&quot;/&gt;
|website={{URL|http://www.montereybayaquarium.org}}
}}

The '''Monterey Bay Aquarium''' ('''MBA''') is a non-profit [[public aquarium]] located in [[Monterey, California]], United States.  The aquarium was founded in 1984 and is located on the site of a former sardine cannery on [[Cannery Row]]. It has an annual attendance of around two million visitors. It holds thousands of plants and animals, representing more than 600 species on display. The aquarium benefits from a high circulation of fresh ocean water which is obtained through pipes which pump it in continuously from [[Monterey Bay]].

The centerpiece of the ''Ocean's Edge Wing'', is a {{convert|28|ft|m|adj=mid|-high}}, {{Convert|333000|gal|sing=on}} exhibit for viewing California coastal marine life. In this exhibit, the aquarium was the first in the world to grow live California [[Giant Kelp]]. Visitors are able to inspect the creatures of the kelp forest at several levels in the building. The largest exhibit in the aquarium is a {{Convert|1200000|gal|adj=on|sp=us}} the Open Sea exhibit (formerly the Outer Bay), which features one of the world's largest single-paned windows. It is one of the few aquariums to successfully care for  the [[ocean sunfish]] in captivity.

Sea life on exhibit includes [[stingray]]s, [[jellyfish]], [[sea otters]], sea horses, and numerous other native marine species, which can be viewed above and below the waterline. The Monterey Bay Aquarium is one of very few in the world to exhibit both [[Pacific bluefin tuna|bluefin]] and [[yellowfin tuna]]. For displaying jellyfish, it uses a [[Aquarium#Styles|Kreisel tank]], which creates a circular flow to support and suspend the jellies. The aquarium does not house mammals other than sea otters that were rescued through its [http://www.montereybayaquarium.org/conservation-and-science/our-priorities/thriving-ocean-wildlife/southern-sea-otters Sea Otter Program].

==History==
There had been a number of attempts to build an aquarium in the Monterey area dating back almost 100 years. In 1914 an aquarium was proposed to the city council by Frank Booth with a cost of $10,000. A bond issue was sponsored in an attempt to place an aquarium in the basement of the Pacific Grove Museum by [[Knut Hovden]] in 1925 and in 1944 an aquarium is suggested for Point Lobos State Reserve.&lt;ref name=&quot;Explore our History&quot;&gt;{{cite web | url=http://www.montereybayaquarium.org/about/our-history | title=Early Years | publisher=Monterey Bay Aquarium Foundation | accessdate=6 July 2013}}&lt;/ref&gt;

The aquarium occupies land at the end of [[Cannery Row]] (once Ocean View Avenue) in Monterey, at the site of the Hovden Cannery, a sardine cannery that helped to define the character of Monterey from the time it was built in 1916 to the day when it was the last cannery on the Row to close in 1973, after sardine fishing collapsed. This building was dismantled in 1980, but beginning in 2002 the Monterey Bay Aquarium has blown the original Hovden Cannery steam whistle at noon each day to commemorate it.&lt;ref name=&quot;timeline5&quot;&gt;
{{Cite web
|url=http://www.montereybayaquarium.org/about/our-history
|title=Explore Our History: 1980–1983
|work=montereybayaquarium.org
|publisher=Monterey Bay Aquarium
|accessdate=8 April 2012
}}&lt;/ref&gt;
[[File:Pacific Biological Laboratories.JPG|thumb|right|The Pacific Biological Laboratories of [[Ed Ricketts]], on [[Cannery Row]], [[Monterey, California]].]]
The aquarium's original building was designed by the architectural firm [[EHDD|Esherick Homsey Dodge &amp; Davis]] and opened on 20 October 1984. The aquarium's mission is &quot;to inspire conservation of the oceans.&quot;  The aquarium's initial financial backing was provided by [[David Packard]], co-founder of [[Hewlett-Packard]].  Packard, an avid [[blacksmith]], personally designed and created several exhibit elements for the aquarium at his forge in [[Big Sur]], including the wave machines in the Kelp Forest and aviary. His daughter, marine biologist Julie Packard, is currently Executive Director of the aquarium.

The aquarium was built in honor of the work of [[Edward Ricketts]] (1897–1948), a marine biologist who specialized in describing communities of organisms (which would also be the focus of aquarium tanks), and whose old laboratory ([[Pacific Biological Laboratories]]) (PCL)  and home resides next to the present MBA site.&lt;ref name=&quot;timeline2&quot;&gt;
{{Cite web
|url=http://www.montereybayaquarium.org/about/our-history
|title=Explore Our History: 1977
|work=montereybayaquarium.org
|publisher=Monterey Bay Aquarium
|accessdate=8 April 2012
}}&lt;/ref&gt; Ricketts, whose life was an inspiration for the eventual building of the aquarium, is famous as the &quot;Doc&quot; of [[John Steinbeck]]'s ''[[Cannery Row (novel)|Cannery Row]]'' and ''[[Sweet Thursday]]''. The aquarium itself contains a display of Ricketts' items, including some of his personal library. The shop also sells a variety of Steinbeck books.`,
		"zoo",
		map[string]string{
			"zoo_name" : "Monterey Bay Aquarium",
			"image" : "MontereyBayAquariumBackview.jpg",
			"image_caption" : "Back view of the aquarium (on the [[Pacific Ocean]])",
			"logo" : "Monterey Bay Aquarium Logo.svg",
			"logo_width" : "90px",
			"logo_caption" : "The aquarium's logo depicts [[Macrocystis pyrifera|giant kelp]].",
			"date_opened" : "October 20, 1984&lt;ref name=FAQ/&gt;",
			"location" : "[[Monterey, California]], US",
			"coordinates" : "{{coord|36.618253|-121.901481|region:US_type:landmark|display=it}}",
			"area" : "",
			"floorspace" : "",
			"num_animals" : "35,000&lt;ref name=&quot;2005_press_kit&quot;/&gt;",
			"num_species" : "623 (plants and animals in 2005)&lt;ref name=&quot;2005_press_kit&quot;/&gt;",
			"largest_tank_vol" : "{{Convert|1200000|gal|sing=on}}",
			"total_tank_vol" : "",
			"annual_visitors" : "2.08 Million&lt;ref&gt;{{cite news |last=Abel |first=David |url=https://www.bostonglobe.com/lifestyle/2016/08/02/aquariumside/aeNq1fhIiUd3i2vZlQ6RZJ/story.html |title=Top aquariums in the US, in terms of visitors |work=[[Boston Globe]] |date=2016-08-02 |accessdate=2016-12-02 }}&lt;/ref&gt;",
			"members" : "[[Association of Zoos and Aquariums|AZA]]&lt;ref name=&quot;aza_list&quot;/&gt;",
			"website" : "{{URL|http://www.montereybayaquarium.org}}",
		},
		`{{Primary sources|date=February 2014}}


The '''Monterey Bay Aquarium''' ('''MBA''') is a non-profit [[public aquarium]] located in [[Monterey, California]], United States.  The aquarium was founded in 1984 and is located on the site of a former sardine cannery on [[Cannery Row]]. It has an annual attendance of around two million visitors. It holds thousands of plants and animals, representing more than 600 species on display. The aquarium benefits from a high circulation of fresh ocean water which is obtained through pipes which pump it in continuously from [[Monterey Bay]].

The centerpiece of the ''Ocean's Edge Wing'', is a {{convert|28|ft|m|adj=mid|-high}}, {{Convert|333000|gal|sing=on}} exhibit for viewing California coastal marine life. In this exhibit, the aquarium was the first in the world to grow live California [[Giant Kelp]]. Visitors are able to inspect the creatures of the kelp forest at several levels in the building. The largest exhibit in the aquarium is a {{Convert|1200000|gal|adj=on|sp=us}} the Open Sea exhibit (formerly the Outer Bay), which features one of the world's largest single-paned windows. It is one of the few aquariums to successfully care for  the [[ocean sunfish]] in captivity.

Sea life on exhibit includes [[stingray]]s, [[jellyfish]], [[sea otters]], sea horses, and numerous other native marine species, which can be viewed above and below the waterline. The Monterey Bay Aquarium is one of very few in the world to exhibit both [[Pacific bluefin tuna|bluefin]] and [[yellowfin tuna]]. For displaying jellyfish, it uses a [[Aquarium#Styles|Kreisel tank]], which creates a circular flow to support and suspend the jellies. The aquarium does not house mammals other than sea otters that were rescued through its [http://www.montereybayaquarium.org/conservation-and-science/our-priorities/thriving-ocean-wildlife/southern-sea-otters Sea Otter Program].

==History==
There had been a number of attempts to build an aquarium in the Monterey area dating back almost 100 years. In 1914 an aquarium was proposed to the city council by Frank Booth with a cost of $10,000. A bond issue was sponsored in an attempt to place an aquarium in the basement of the Pacific Grove Museum by [[Knut Hovden]] in 1925 and in 1944 an aquarium is suggested for Point Lobos State Reserve.&lt;ref name=&quot;Explore our History&quot;&gt;{{cite web | url=http://www.montereybayaquarium.org/about/our-history | title=Early Years | publisher=Monterey Bay Aquarium Foundation | accessdate=6 July 2013}}&lt;/ref&gt;

The aquarium occupies land at the end of [[Cannery Row]] (once Ocean View Avenue) in Monterey, at the site of the Hovden Cannery, a sardine cannery that helped to define the character of Monterey from the time it was built in 1916 to the day when it was the last cannery on the Row to close in 1973, after sardine fishing collapsed. This building was dismantled in 1980, but beginning in 2002 the Monterey Bay Aquarium has blown the original Hovden Cannery steam whistle at noon each day to commemorate it.&lt;ref name=&quot;timeline5&quot;&gt;
{{Cite web
|url=http://www.montereybayaquarium.org/about/our-history
|title=Explore Our History: 1980–1983
|work=montereybayaquarium.org
|publisher=Monterey Bay Aquarium
|accessdate=8 April 2012
}}&lt;/ref&gt;
[[File:Pacific Biological Laboratories.JPG|thumb|right|The Pacific Biological Laboratories of [[Ed Ricketts]], on [[Cannery Row]], [[Monterey, California]].]]
The aquarium's original building was designed by the architectural firm [[EHDD|Esherick Homsey Dodge &amp; Davis]] and opened on 20 October 1984. The aquarium's mission is &quot;to inspire conservation of the oceans.&quot;  The aquarium's initial financial backing was provided by [[David Packard]], co-founder of [[Hewlett-Packard]].  Packard, an avid [[blacksmith]], personally designed and created several exhibit elements for the aquarium at his forge in [[Big Sur]], including the wave machines in the Kelp Forest and aviary. His daughter, marine biologist Julie Packard, is currently Executive Director of the aquarium.

The aquarium was built in honor of the work of [[Edward Ricketts]] (1897–1948), a marine biologist who specialized in describing communities of organisms (which would also be the focus of aquarium tanks), and whose old laboratory ([[Pacific Biological Laboratories]]) (PCL)  and home resides next to the present MBA site.&lt;ref name=&quot;timeline2&quot;&gt;
{{Cite web
|url=http://www.montereybayaquarium.org/about/our-history
|title=Explore Our History: 1977
|work=montereybayaquarium.org
|publisher=Monterey Bay Aquarium
|accessdate=8 April 2012
}}&lt;/ref&gt; Ricketts, whose life was an inspiration for the eventual building of the aquarium, is famous as the &quot;Doc&quot; of [[John Steinbeck]]'s ''[[Cannery Row (novel)|Cannery Row]]'' and ''[[Sweet Thursday]]''. The aquarium itself contains a display of Ricketts' items, including some of his personal library. The shop also sells a variety of Steinbeck books.`,
	},
}

func testOneParseInfobox(t *testing.T, ti testInfoText) {
	res, err := ParseInfobox(ti.input)

	if err != nil {
		t.Fatalf("Unexpected error on %v: %v\n", ti, err)
		fmt.Println(err)
	}

	if res.TemplateType != ti.templateType {
		t.Fatalf("Expected template type of %v, got %v\n", ti.templateType, res.TemplateType)
	}

	for k, v := range ti.attributes {
		val, ok := res.Attributes[k]
		switch {
		case !ok:
			t.Fatalf("Expected result to have attribute %s but it was missing", k)
		case v != val:
			t.Fatalf("Expected attribute %s to have value %s but had %s", k, v, res.Attributes[k])
		}
	}
}

func testOneWithoutInfobox(t *testing.T, ti testInfoText) {
	res := WithoutInfobox(ti.input)
	if res != ti.withoutInfobox {
		t.Fatalf("Expected to get `%s`, but got `%s`\n", ti.withoutInfobox, res)
	}
}

func TestParseInfobox(t *testing.T) {
	t.Parallel()
	for _, ti := range infoTestData {
		testOneParseInfobox(t, ti)
	}
}

func TestWithoutInfobox(t *testing.T) {
	t.Parallel()
	for _, ti := range infoTestData {
		testOneWithoutInfobox(t, ti)
	}
}
