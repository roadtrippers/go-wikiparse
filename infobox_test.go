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
		if res.Attributes[k] != v {
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
