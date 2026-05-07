package types

type Mappings struct {
	Letters map[LetterTwo]string
	Vowels  map[VowelTwo]string
}

func DefaultMappings() Mappings {
	return Mappings{
		Letters: defaultLetterSounds(),
		Vowels:  defaultVowelSounds(),
	}
}

func defaultLetterSounds() map[LetterTwo]string {
	m := map[LetterTwo]string{
		LetterTwoNone:   "",
		LetterTwoAleph:  "a", // TODO: or silent
		LetterTwoBet:    "b",
		LetterTwoVet:    "v",
		LetterTwoGimmel: "g",
		LetterTwoDalet:  "d",
		LetterTwoHei:    "h", // TODO: or silent at end of word
		LetterTwoVav:    "v",
		LetterTwoZayin:  "z",
		LetterTwoChet:   "ch",
		LetterTwoTet:    "t",
		LetterTwoYud:    "y",
		LetterTwoKaf:    "k",
		LetterTwoChaf:   "ch",
		LetterTwoLamed:  "l",
		LetterTwoMem:    "m",
		LetterTwoNun:    "n",
		LetterTwoSamech: "s",
		LetterTwoAyin:   "a", // TODO: Or silent
		LetterTwoPey:    "p",
		LetterTwoPhey:   "ph",
		LetterTwoTzadi:  "tz",
		LetterTwoKuf:    "k",
		LetterTwoResh:   "r",
		LetterTwoShin:   "sh",
		LetterTwoSin:    "s",
		LetterTwoTaf:    "t",
		LetterTwoSaf:    "t",
	}
	return m

}

func defaultVowelSounds() map[VowelTwo]string {
	m := map[VowelTwo]string{
		VowelTwoNone:         "",
		VowelTwoShva:         "'", // TODO: Na vs. Nach
		VowelTwoChatafSegol:  "e",
		VowelTwoChatafPatach: "a",
		VowelTwoChatafKamatz: "o",
		VowelTwoChirik:       "i",
		VowelTwoChirikMale:   "i",
		VowelTwoTsere:        "e",
		VowelTwoSegol:        "e",
		VowelTwoPatach:       "a", // TODO: Furtive
		VowelTwoKamatz:       "a",
		VowelTwoCholom:       "o",
		VowelTwoCholomMale:   "o",
		VowelTwoKubutz:       "u",
		VowelTwoShuruk:       "u",
		VowelTwoKamatzKatan:  "o",
	}
	return m
}

// <spacey> <yud> <yud> <spacey>
// <spacey> <yud> <hey> <vav> <hey> <spacey>

func isYud()

func IsGod(head *LinkedNodeTwo) bool {
	curr := head
	state := 0
	for curr != nil {
		switch state {
		case 0:
			if isSpacey(curr.Node) {
				state = 1
			} else {
				return false
			}
		case 1:
			if isYud(curr.Node) {
				state = 2
			} else {
				return false
			}
		case 2:
			if isYud(curr.Node) {
				state = 3
			} else if true {

			} else {
				return false
			}
		}
	}
}

func isSpacey(node NodeTwo) bool {
	switch node.(type) {
	case *SpaceNodeTwo:
		return true
	case *PuncNodeTwo:
		return true
	default:
		return false
	}
}

func ReplaceGod(head *LinkedNodeTwo) *LinkedNodeTwo {
}
