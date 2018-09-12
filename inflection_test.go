package inflection_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tjimsk/inflection"
	"testing"
)

func TestInflections(t *testing.T) {
	type testData struct {
		singular string
		plural   string
	}

	data := []testData{
		testData{"ability", "abilities"},
		testData{"alga", "algae"},
		testData{"agency", "agencies"},
		testData{"analysis", "analyses"},
		testData{"archive", "archives"},
		testData{"axis", "axes"},
		testData{"basis", "bases"},
		testData{"buffalo", "buffaloes"},
		testData{"bus", "buses"},
		testData{"calf", "calves"},
		testData{"child", "children"},
		testData{"comment", "comments"},
		testData{"criterion", "criteria"},
		testData{"crisis", "crises"},
		testData{"datum", "data"},
		testData{"day", "days"},
		testData{"diagnosis", "diagnoses"},
		testData{"diagnosis_a", "diagnosis_as"},
		testData{"dwarf", "dwarves"},
		testData{"elf", "elves"},
		testData{"ellipsis", "ellipses"},
		testData{"emphasis", "emphases"},
		testData{"equipment", "equipment"},
		testData{"experience", "experiences"},
		testData{"fish", "fish"},
		testData{"fireman", "firemen"},
		testData{"half", "halves"},
		testData{"foobar", "foobars"},
		testData{"hero", "heroes"},
		testData{"index", "indices"},
		testData{"information", "information"},
		testData{"liquid", "liquids"},
		testData{"man", "men"},
		testData{"medium", "media"},
		testData{"mosquito", "mosquitoes"},
		testData{"mouse", "mice"},
		testData{"move", "moves"},
		testData{"movie", "movies"},
		testData{"news", "news"},
		testData{"newsletter", "newsletters"},
		testData{"node_child", "node_children"},
		testData{"old_news", "old_news"},
		testData{"ox", "oxen"},
		testData{"person", "people"},
		testData{"perspective", "perspectives"},
		testData{"photo", "photos"},
		testData{"product", "products"},
		testData{"query", "queries"},
		testData{"quiz", "quizzes"},
		testData{"safe", "saves"},
		testData{"salesperson", "salespeople"},
		testData{"scissors", "scissors"},
		testData{"series", "series"},
		testData{"shelf", "shelves"},
		testData{"species", "species"},
		testData{"spokesman", "spokesmen"},
		testData{"stadium", "stadia"},
		testData{"star", "stars"},
		testData{"STAR", "STARS"},
		testData{"Star", "Stars"},
		testData{"stock", "stocks"},
		testData{"STOCK", "STOCKS"},
		testData{"tomato", "tomatoes"},
		testData{"user", "users"},
		testData{"wife", "wives"},
		testData{"woman", "women"},
	}

	for _, td := range data {
		testPluralization(t, td.singular, td.plural)
		testSingularization(t, td.plural, td.singular)
	}
}

func testPluralization(t *testing.T, singular string, expected string) {
	plural := inflection.Pluralize(singular)
	assert.Equal(t, expected, plural, "wrong plural for %v", singular)
}

func testSingularization(t *testing.T, plural string, expected string) {
	singular := inflection.Singularize(plural)
	assert.Equal(t, expected, singular, "wrong singular for %v", plural)
}
