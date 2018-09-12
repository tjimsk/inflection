package inflection

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	pluralize   []*Rule
	singularize []*Rule
)

type Rule struct {
	singular   string
	plural     string
	singularRe *regexp.Regexp
	pluralRe   *regexp.Regexp
}

func (r *Rule) compile() (err error) {
	r.singularRe, err = regexp.Compile(r.singular)
	if err != nil {
		return err
	}

	r.pluralRe, err = regexp.Compile(r.plural)
	if err != nil {
		return err
	}

	return nil
}

var plurals = []*Rule{
	&Rule{singular: "([a-z])$", plural: "${1}s"},
	&Rule{singular: "s$", plural: "s"},
	&Rule{singular: "^(ax|test)is$", plural: "${1}es"},
	&Rule{singular: "(octop|vir)us$", plural: "${1}i"},
	&Rule{singular: "(octop|vir)i$", plural: "${1}i"},
	&Rule{singular: "(alias|status)$", plural: "${1}es"},
	&Rule{singular: "(bu)s$", plural: "${1}ses"},
	&Rule{singular: "(buffal|tomat)o$", plural: "${1}oes"},
	&Rule{singular: "([ti])um$", plural: "${1}a"},
	&Rule{singular: "([ti])a$", plural: "${1}a"},
	&Rule{singular: "sis$", plural: "ses"},
	&Rule{singular: "(?:([^f])fe|([lr])f)$", plural: "${1}${2}ves"},
	&Rule{singular: "(hive)$", plural: "${1}s"},
	&Rule{singular: "([^aeiouy]|qu)y$", plural: "${1}ies"},
	&Rule{singular: "(x|ch|ss|sh)$", plural: "${1}es"},
	&Rule{singular: "(matr|vert|ind)(?:ix|ex)$", plural: "${1}ices"},
	&Rule{singular: "^(m|l)ouse$", plural: "${1}ice"},
	&Rule{singular: "^(m|l)ice$", plural: "${1}ice"},
	&Rule{singular: "^(ox)$", plural: "${1}en"},
	&Rule{singular: "^(oxen)$", plural: "${1}"},
	&Rule{singular: "(quiz)$", plural: "${1}zes"},
}

var singulars = []*Rule{
	&Rule{plural: "s$", singular: ""},
	&Rule{plural: "(ss)$", singular: "${1}"},
	&Rule{plural: "(n)ews$", singular: "${1}ews"},
	&Rule{plural: "([ti])a$", singular: "${1}um"},
	&Rule{plural: "((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$", singular: "${1}sis"},
	&Rule{plural: "(^analy)(sis|ses)$", singular: "${1}sis"},
	&Rule{plural: "([^f])ves$", singular: "${1}fe"},
	&Rule{plural: "(hive)s$", singular: "${1}"},
	&Rule{plural: "(tive)s$", singular: "${1}"},
	&Rule{plural: "([lr])ves$", singular: "${1}f"},
	&Rule{plural: "([^aeiouy]|qu)ies$", singular: "${1}y"},
	&Rule{plural: "(s)eries$", singular: "${1}eries"},
	&Rule{plural: "(m)ovies$", singular: "${1}ovie"},
	&Rule{plural: "(c)ookies$", singular: "${1}ookie"},
	&Rule{plural: "(x|ch|ss|sh)es$", singular: "${1}"},
	&Rule{plural: "^(m|l)ice$", singular: "${1}ouse"},
	&Rule{plural: "(bus)(es)?$", singular: "${1}"},
	&Rule{plural: "(o)es$", singular: "${1}"},
	&Rule{plural: "(shoe)s$", singular: "${1}"},
	&Rule{plural: "(cris|test)(is|es)$", singular: "${1}is"},
	&Rule{plural: "^(a)x[ie]s$", singular: "${1}xis"},
	&Rule{plural: "(octop|vir)(us|i)$", singular: "${1}us"},
	&Rule{plural: "(alias|status)(es)?$", singular: "${1}"},
	&Rule{plural: "^(ox)en", singular: "${1}"},
	&Rule{plural: "(vert|ind)ices$", singular: "${1}ex"},
	&Rule{plural: "(matr)ices$", singular: "${1}ix"},
	&Rule{plural: "(quiz)zes$", singular: "${1}"},
	&Rule{plural: "(database)s$", singular: "${1}"},
}

var irregulars = []*Rule{
	&Rule{singular: "addendum", plural: "addenda"},
	&Rule{singular: "alga", plural: "algae"},
	&Rule{singular: "alumna", plural: "alumnae"},
	&Rule{singular: "alumnus", plural: "alumni"},
	&Rule{singular: "analysis", plural: "analyses"},
	&Rule{singular: "antenna", plural: "antennae"},
	&Rule{singular: "apparatus", plural: "apparatuses"},
	&Rule{singular: "appendix", plural: "appendices"},
	// &Rule{singular: "axis", plural: "axes"},
	&Rule{singular: "bacillus", plural: "bacilli"},
	&Rule{singular: "bacterium", plural: "bacteria"},
	&Rule{singular: "basis", plural: "bases"},
	&Rule{singular: "beau", plural: "beaux"},
	&Rule{singular: "bison", plural: "bison"},
	&Rule{singular: "buffalo", plural: "buffaloes"},
	&Rule{singular: "bureau", plural: "bureaus"},
	&Rule{singular: "bus", plural: "buses"},
	&Rule{singular: "cactus", plural: "cacti"},
	// &Rule{singular: "calf", plural: "calves"},
	&Rule{singular: "child", plural: "children"},
	&Rule{singular: "corps", plural: "corps"},
	&Rule{singular: "corpus", plural: "corpora"},
	// &Rule{singular: "crisis", plural: "crises"},
	&Rule{singular: "criterion", plural: "criteria"},
	&Rule{singular: "curriculum", plural: "curricula"},
	&Rule{singular: "datum", plural: "data"},
	&Rule{singular: "deer", plural: "deer"},
	&Rule{singular: "die", plural: "dice"},
	// &Rule{singular: "dwarf", plural: "dwarves"},
	&Rule{singular: "diagnosis", plural: "diagnoses"},
	&Rule{singular: "echo", plural: "echoes"},
	&Rule{singular: "elf", plural: "elves"},
	&Rule{singular: "ellipsis", plural: "ellipses"},
	&Rule{singular: "embargo", plural: "embargoes"},
	&Rule{singular: "emphasis", plural: "emphases"},
	&Rule{singular: "erratum", plural: "errata"},
	&Rule{singular: "fireman", plural: "firemen"},
	&Rule{singular: "fish", plural: "fish"},
	&Rule{singular: "focus", plural: "focuses"},
	&Rule{singular: "foot", plural: "feet"},
	&Rule{singular: "formula", plural: "formulas"},
	&Rule{singular: "fungus", plural: "fungi"},
	&Rule{singular: "genus", plural: "genera"},
	&Rule{singular: "goose", plural: "geese"},
	// &Rule{singular: "half", plural: "halves"},
	&Rule{singular: "hero", plural: "heroes"},
	&Rule{singular: "hippopotamus", plural: "hippopotami"},
	&Rule{singular: "hoof", plural: "hooves"},
	&Rule{singular: "hypothesis", plural: "hypotheses"},
	&Rule{singular: "index", plural: "indices"},
	&Rule{singular: "knife", plural: "knives"},
	&Rule{singular: "leaf", plural: "leaves"},
	&Rule{singular: "life", plural: "lives"},
	&Rule{singular: "loaf", plural: "loaves"},
	&Rule{singular: "louse", plural: "lice"},
	&Rule{singular: "man", plural: "men"},
	&Rule{singular: "matrix", plural: "matrices"},
	&Rule{singular: "means", plural: "means"},
	&Rule{singular: "medium", plural: "media"},
	&Rule{singular: "memorandum", plural: "memoranda"},
	&Rule{singular: "millennium", plural: "milennia"},
	&Rule{singular: "mombie", plural: "mombies"},
	&Rule{singular: "moose", plural: "moose"},
	&Rule{singular: "mosquito", plural: "mosquitoes"},
	&Rule{singular: "mouse", plural: "mice"},
	&Rule{singular: "move", plural: "moves"},
	&Rule{singular: "nebula", plural: "nebulaenebulas"},
	&Rule{singular: "neurosis", plural: "neuroses"},
	&Rule{singular: "nucleus", plural: "nuclei"},
	&Rule{singular: "oasis", plural: "oases"},
	&Rule{singular: "octopus", plural: "octopi"},
	&Rule{singular: "ovum", plural: "ova"},
	&Rule{singular: "ox", plural: "oxen"},
	&Rule{singular: "paralysis", plural: "paralyses"},
	&Rule{singular: "parenthesis", plural: "parentheses"},
	&Rule{singular: "person", plural: "people"},
	&Rule{singular: "phenomenon", plural: "phenomena"},
	&Rule{singular: "potato", plural: "potatoes"},
	&Rule{singular: "radius", plural: "radii"},
	&Rule{singular: "scarf", plural: "scarves"},
	&Rule{singular: "sex", plural: "sexes"},
	&Rule{singular: "self", plural: "selves"},
	&Rule{singular: "series", plural: "series"},
	&Rule{singular: "sheep", plural: "sheep"},
	&Rule{singular: "scissors", plural: "scissors"},
	&Rule{singular: "species", plural: "species"},
	&Rule{singular: "stimulus", plural: "stimuli"},
	&Rule{singular: "stratum", plural: "strata"},
	&Rule{singular: "syllabus", plural: "syllabi"},
	&Rule{singular: "symposium", plural: "symposia"},
	&Rule{singular: "synthesis", plural: "syntheses"},
	&Rule{singular: "synopsis", plural: "synopses"},
	&Rule{singular: "tableau", plural: "tableaux"},
	&Rule{singular: "that", plural: "those"},
	&Rule{singular: "thesis", plural: "theses"},
	&Rule{singular: "thief", plural: "thieves"},
	&Rule{singular: "this", plural: "these"},
	&Rule{singular: "tomato", plural: "tomatoes"},
	&Rule{singular: "tooth", plural: "teeth"},
	&Rule{singular: "torpedo", plural: "torpedoes"},
	&Rule{singular: "vertebra", plural: "vertebrae"},
	&Rule{singular: "veto", plural: "vetoes"},
	&Rule{singular: "vita", plural: "vitae"},
	&Rule{singular: "watch", plural: "watches"},
	&Rule{singular: "wife", plural: "wives"},
	&Rule{singular: "wolf", plural: "wolves"},
	&Rule{singular: "woman", plural: "women"},
	&Rule{singular: "zero", plural: "zeroes"},
}

var uncountables = []*Rule{
	&Rule{singular: "accommodation", plural: "accommodation"},
	&Rule{singular: "advertising", plural: "advertising"},
	&Rule{singular: "air", plural: "air"},
	&Rule{singular: "aid", plural: "aid"},
	&Rule{singular: "advice", plural: "advice"},
	&Rule{singular: "anger", plural: "anger"},
	&Rule{singular: "art", plural: "art"},
	&Rule{singular: "assistance", plural: "assistance"},
	&Rule{singular: "bread", plural: "bread"},
	&Rule{singular: "business", plural: "business"},
	&Rule{singular: "butter", plural: "butter"},
	&Rule{singular: "calm", plural: "calm"},
	&Rule{singular: "cash", plural: "cash"},
	&Rule{singular: "chaos", plural: "chaos"},
	&Rule{singular: "cheese", plural: "cheese"},
	&Rule{singular: "childhood", plural: "childhood"},
	&Rule{singular: "clothing", plural: "clothing"},
	&Rule{singular: "coffee", plural: "coffee"},
	&Rule{singular: "content", plural: "content"},
	&Rule{singular: "corruption", plural: "corruption"},
	&Rule{singular: "courage", plural: "courage"},
	&Rule{singular: "currency", plural: "currency"},
	&Rule{singular: "damage", plural: "damage"},
	&Rule{singular: "danger", plural: "danger"},
	&Rule{singular: "darkness", plural: "darkness"},
	&Rule{singular: "determination", plural: "determination"},
	&Rule{singular: "economics", plural: "economics"},
	&Rule{singular: "education", plural: "education"},
	&Rule{singular: "electricity", plural: "electricity"},
	&Rule{singular: "employment", plural: "employment"},
	&Rule{singular: "energy", plural: "energy"},
	&Rule{singular: "entertainment", plural: "entertainment"},
	&Rule{singular: "enthusiasm", plural: "enthusiasm"},
	&Rule{singular: "equipment", plural: "equipment"},
	&Rule{singular: "evidence", plural: "evidence"},
	&Rule{singular: "failure", plural: "failure"},
	&Rule{singular: "fame", plural: "fame"},
	&Rule{singular: "fire", plural: "fire"},
	&Rule{singular: "flour", plural: "flour"},
	&Rule{singular: "food", plural: "food"},
	&Rule{singular: "freedom", plural: "freedom"},
	&Rule{singular: "friendship", plural: "friendship"},
	&Rule{singular: "fuel", plural: "fuel"},
	&Rule{singular: "furniture", plural: "furniture"},
	&Rule{singular: "fun", plural: "fun"},
	&Rule{singular: "genetics", plural: "genetics"},
	&Rule{singular: "gold", plural: "gold"},
	&Rule{singular: "grammar", plural: "grammar"},
	&Rule{singular: "guilt", plural: "guilt"},
	&Rule{singular: "hair", plural: "hair"},
	&Rule{singular: "happiness", plural: "happiness"},
	&Rule{singular: "harm", plural: "harm"},
	&Rule{singular: "health", plural: "health"},
	&Rule{singular: "heat", plural: "heat"},
	&Rule{singular: "help", plural: "help"},
	&Rule{singular: "homework", plural: "homework"},
	&Rule{singular: "honesty", plural: "honesty"},
	&Rule{singular: "hospitality", plural: "hospitality"},
	&Rule{singular: "housework", plural: "housework"},
	&Rule{singular: "humour", plural: "humour"},
	&Rule{singular: "imagination", plural: "imagination"},
	&Rule{singular: "importance", plural: "importance"},
	&Rule{singular: "information", plural: "information"},
	&Rule{singular: "innocence", plural: "innocence"},
	&Rule{singular: "intelligence", plural: "intelligence"},
	&Rule{singular: "jealousy", plural: "jealousy"},
	&Rule{singular: "juice", plural: "juice"},
	&Rule{singular: "justice", plural: "justice"},
	&Rule{singular: "kindness", plural: "kindness"},
	&Rule{singular: "knowledge", plural: "knowledge"},
	&Rule{singular: "labour", plural: "labour"},
	&Rule{singular: "lack", plural: "lack"},
	&Rule{singular: "laughter", plural: "laughter"},
	&Rule{singular: "leisure", plural: "leisure"},
	&Rule{singular: "literature", plural: "literature"},
	&Rule{singular: "litter", plural: "litter"},
	&Rule{singular: "logic", plural: "logic"},
	&Rule{singular: "love", plural: "love"},
	&Rule{singular: "luck", plural: "luck"},
	&Rule{singular: "magic", plural: "magic"},
	&Rule{singular: "management", plural: "management"},
	&Rule{singular: "metal", plural: "metal"},
	&Rule{singular: "milk", plural: "milk"},
	&Rule{singular: "money", plural: "money"},
	&Rule{singular: "motherhood", plural: "motherhood"},
	&Rule{singular: "motivation", plural: "motivation"},
	&Rule{singular: "music", plural: "music"},
	&Rule{singular: "nature", plural: "nature"},
	&Rule{singular: "nutrition", plural: "nutrition"},
	&Rule{singular: "obesity", plural: "obesity"},
	&Rule{singular: "oil", plural: "oil"},
	&Rule{singular: "old age", plural: "old age"},
	&Rule{singular: "oxygen", plural: "oxygen"},
	&Rule{singular: "paper", plural: "paper"},
	&Rule{singular: "patience", plural: "patience"},
	&Rule{singular: "permission", plural: "permission"},
	&Rule{singular: "pollution", plural: "pollution"},
	&Rule{singular: "poverty", plural: "poverty"},
	&Rule{singular: "power", plural: "power"},
	&Rule{singular: "pride", plural: "pride"},
	&Rule{singular: "production", plural: "production"},
	&Rule{singular: "progress", plural: "progress"},
	&Rule{singular: "pronunciation", plural: "pronunciation"},
	&Rule{singular: "publicity", plural: "publicity"},
	&Rule{singular: "punctuation", plural: "punctuation"},
	&Rule{singular: "quality", plural: "quality"},
	&Rule{singular: "quantity", plural: "quantity"},
	&Rule{singular: "racism", plural: "racism"},
	&Rule{singular: "rain", plural: "rain"},
	&Rule{singular: "relaxation", plural: "relaxation"},
	&Rule{singular: "research", plural: "research"},
	&Rule{singular: "respect", plural: "respect"},
	&Rule{singular: "rice", plural: "rice"},
	&Rule{singular: "room", plural: "room"},
	&Rule{singular: "rubbish", plural: "rubbish"},
	&Rule{singular: "safety", plural: "safety"},
	&Rule{singular: "salt", plural: "salt"},
	&Rule{singular: "sand", plural: "sand"},
	&Rule{singular: "seafood", plural: "seafood"},
	&Rule{singular: "shopping", plural: "shopping"},
	&Rule{singular: "silence", plural: "silence"},
	&Rule{singular: "smoke", plural: "smoke"},
	&Rule{singular: "snow", plural: "snow"},
	&Rule{singular: "software", plural: "software"},
	&Rule{singular: "soup", plural: "soup"},
	&Rule{singular: "speed", plural: "speed"},
	&Rule{singular: "spelling", plural: "spelling"},
	&Rule{singular: "stress", plural: "stress"},
	&Rule{singular: "sugar", plural: "sugar"},
	&Rule{singular: "sunshine", plural: "sunshine"},
	&Rule{singular: "tea", plural: "tea"},
	&Rule{singular: "tennis", plural: "tennis"},
	&Rule{singular: "time", plural: "time"},
	&Rule{singular: "tolerance", plural: "tolerance"},
	&Rule{singular: "trade", plural: "trade"},
	&Rule{singular: "traffic", plural: "traffic"},
	&Rule{singular: "transportation", plural: "transportation"},
	&Rule{singular: "travel", plural: "travel"},
	&Rule{singular: "trust", plural: "trust"},
	&Rule{singular: "understanding", plural: "understanding"},
	&Rule{singular: "unemployment", plural: "unemployment"},
	&Rule{singular: "usage", plural: "usage"},
	&Rule{singular: "violence", plural: "violence"},
	&Rule{singular: "vision", plural: "vision"},
	&Rule{singular: "warmth", plural: "warmth"},
	&Rule{singular: "water", plural: "water"},
	&Rule{singular: "wealth", plural: "wealth"},
	&Rule{singular: "weather", plural: "weather"},
	&Rule{singular: "weight", plural: "weight"},
	&Rule{singular: "welfare", plural: "welfare"},
	&Rule{singular: "wheat", plural: "wheat"},
	&Rule{singular: "width", plural: "width"},
	&Rule{singular: "wildlife", plural: "wildlife"},
	&Rule{singular: "wisdom", plural: "wisdom"},
	&Rule{singular: "wood", plural: "wood"},
	&Rule{singular: "work", plural: "work"},
	&Rule{singular: "yoga", plural: "yoga"},
	&Rule{singular: "youth", plural: "youth"},
}

func init() {
	for _, r := range plurals {
		pluralize = append(pluralize, caseInsensitivePluralRule(r))
		pluralize = append(pluralize, upperCaseRule(r))
	}

	for _, r := range singulars {
		singularize = append(singularize, caseInsensitiveSingularRule(r))
		singularize = append(singularize, upperCaseRule(r))
	}

	for _, r := range uncountables {
		pluralize = append(pluralize, caseInsensitivePluralRule(r))
		pluralize = append(pluralize, upperCaseRule(r))
		pluralize = append(pluralize, titleCaseRule(r))

		singularize = append(singularize, caseInsensitiveSingularRule(r))
		singularize = append(singularize, upperCaseRule(r))
		singularize = append(singularize, titleCaseRule(r))
	}

	for _, r := range irregulars {
		pluralize = append(pluralize, caseInsensitivePluralRule(delimitedPluralRule(r)))
		pluralize = append(pluralize, upperCaseRule(delimitedPluralRule(r)))
		pluralize = append(pluralize, titleCaseRule(delimitedPluralRule(r)))

		singularize = append(singularize, caseInsensitiveSingularRule(delimitedSingularRule(r)))
		singularize = append(singularize, upperCaseRule(delimitedSingularRule(r)))
		singularize = append(singularize, titleCaseRule(delimitedSingularRule(r)))
	}

	for _, r := range pluralize {
		if err := r.compile(); err != nil {
			panic(err)
		}
	}

	for _, r := range singularize {
		if err := r.compile(); err != nil {
			panic(err)
		}
	}
}

func upperCaseRule(r *Rule) *Rule {
	return &Rule{singular: strings.ToUpper(r.singular), plural: strings.ToUpper(r.plural)}
}

func titleCaseRule(r *Rule) *Rule {
	return &Rule{singular: strings.Title(r.singular), plural: strings.Title(r.plural)}
}

func caseInsensitivePluralRule(r *Rule) *Rule {
	return &Rule{singular: fmt.Sprintf("(?i)%v", r.singular), plural: r.plural}
}

func caseInsensitiveSingularRule(r *Rule) *Rule {
	return &Rule{singular: r.singular, plural: fmt.Sprintf("(?i)%v", r.plural)}
}

func delimitedPluralRule(r *Rule) *Rule {
	return &Rule{singular: fmt.Sprintf("%v$", r.singular), plural: r.plural}
}

func delimitedSingularRule(r *Rule) *Rule {
	return &Rule{singular: r.singular, plural: fmt.Sprintf("%v$", r.plural)}
}

func Pluralize(noun string) (singular string) {
	singular = noun

	for _, r := range pluralize {
		if r.singularRe.MatchString(noun) {
			singular = r.singularRe.ReplaceAllString(noun, r.plural)
		}
	}

	return singular
}

func Singularize(noun string) (plural string) {
	plural = noun

	for _, r := range singularize {
		if r.pluralRe.MatchString(noun) {
			plural = r.pluralRe.ReplaceAllString(noun, r.singular)
		}
	}

	return plural
}
