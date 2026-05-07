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

