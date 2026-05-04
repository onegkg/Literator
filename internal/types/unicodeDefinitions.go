package types

func IsTaam(r rune) bool {
	return r >= 0x0591 && r <= 0x05AF
}

// Vowel represents a Hebrew vowel diacritic.
type Vowel rune

// Vowels
const (
	VowelNone             Vowel = 0
	SHVA                  Vowel = 0x05B0
	CHATAF_SEGOL          Vowel = 0x05B1
	CHATAF_PATACH         Vowel = 0x05B2
	CHATAF_KAMATZ         Vowel = 0x05B3
	CHIRIK                Vowel = 0x05B4
	TSERE                 Vowel = 0x05B5
	SEGOL                 Vowel = 0x05B6
	PATACH                Vowel = 0x05B7
	KAMATZ                Vowel = 0x05B8
	CHOLOM                Vowel = 0x05B9
	CHOLOM_CHASER_FOR_VAV Vowel = 0x05BA
	KUBUTZ                Vowel = 0x05BB
	KAMATZ_KATAN          Vowel = 0x05C7
)

const (
	DAGESH  rune = 0x05BC
	SHINDOT rune = 0x05C1
	SINDOT  rune = 0x05C2
)

func IsShinDot(r rune) bool { return r == SHINDOT }
func IsSinDot(r rune) bool  { return r == SINDOT }

func IsHebrewVowel(r rune) bool {
	return ((r >= 0x05B0 && r <= 0x05BC) || r == 0x05C1 || r == 0x05C2 || r == 0x05C7) && r != DAGESH
}

// Assorted symbols
const METEG = 0x05BD
const MAQAF = 0x05BE
const RAFE = 0x05BF

// Letter represents a Hebrew letter.
type Letter rune

// Letters
const (
	LetterNone  Letter = 0
	ALEPH       Letter = 0x05D0
	BET         Letter = 0x05D1
	GIMMEL      Letter = 0x05D2
	DALET       Letter = 0x05D3
	HEI         Letter = 0x05D4
	VAV         Letter = 0x05D5
	ZAYIN       Letter = 0x05D6
	CHET        Letter = 0x05D7
	TET         Letter = 0x05D8
	YUD         Letter = 0x05D9
	KAF_SOFIT   Letter = 0x05DA
	KAF         Letter = 0x05DB
	LAMED       Letter = 0x05DC
	MEM_SOFIT   Letter = 0x05DD
	MEM         Letter = 0x05DE
	NUN_SOFIT   Letter = 0x05DF
	NUN         Letter = 0x05E0
	SAMECH      Letter = 0x05E1
	AYIN        Letter = 0x05E2
	PEY_SOFIT   Letter = 0x05E3
	PEY         Letter = 0x05E4
	TZADI_SOFIT Letter = 0x05E5
	TZADI       Letter = 0x05E6
	KUF         Letter = 0x05E7
	RESH        Letter = 0x05E8
	SHIN        Letter = 0x05E9
	TAF         Letter = 0x05EA
)

func IsHebrewLetter(r rune) bool {
	return r >= 0x05D0 && r <= 0x05EA
}

func (l Letter) IsFinal() bool {
	return l == KAF_SOFIT ||
		l == MEM_SOFIT ||
		l == NUN_SOFIT ||
		l == PEY_SOFIT ||
		l == TZADI_SOFIT
}

// Sofiyot are handled seperately
func (l Letter) IsBegadkefat() bool {
	return l == BET || l == KAF || l == PEY || l == TAF
}
