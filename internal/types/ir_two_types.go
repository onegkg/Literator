package types

import (
	"fmt"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

type NodeTwo interface {
	IsNodeTwo()
}

type LinkedNodeTwo struct {
	Node NodeTwo
	Next *LinkedNodeTwo
	Prev *LinkedNodeTwo
}

func (ln2 *LinkedNodeTwo) StringFromHead() string {
	var sb strings.Builder
	curr := ln2
	first := true
	fmt.Fprintf(&sb, "[")
	for curr != nil {
		if first {
			fmt.Fprintf(&sb, "%v", curr)
			first = false
		} else {
			fmt.Fprintf(&sb, ", %v", curr)
		}
		curr = curr.Next
	}
	fmt.Fprint(&sb, "]")
	return sb.String()
}

func (ln2 *LinkedNodeTwo) String() string {
	return fmt.Sprintf("LN2{ %v }", ln2.Node)
}

func (ln2 *LinkedNodeTwo) DebugString() string {
	return fmt.Sprintf("LN2{ Node: %v, Next: %v, Prev: %v}", ln2.Node, ln2.Next, ln2.Prev)
}

type EdgeCase int

const (
	EdgeCaseNone EdgeCase = iota
	EdgeCaseSilent
	EdgeCaseFinal
	EdgeCaseMappikHei
)

func (e EdgeCase) String() string {
	switch e {
	case EdgeCaseNone:
		return "None"
	case EdgeCaseSilent:
		return "Silent"
	case EdgeCaseFinal:
		return "Final"
	case EdgeCaseMappikHei:
		return "Mappik Hei"
	default:
		return fmt.Sprintf("Invalid %d", e)
	}
}

type VowelTwo int

const (
	VowelTwoNone         = VowelTwo(-1)
	VowelTwoShva         = VowelTwo(SHVA)
	VowelTwoChatafSegol  = VowelTwo(CHATAF_SEGOL)
	VowelTwoChatafPatach = VowelTwo(CHATAF_PATACH)
	VowelTwoChatafKamatz = VowelTwo(CHATAF_KAMATZ)
	VowelTwoChirik       = VowelTwo(CHIRIK)
	VowelTwoChirikMale   = VowelTwo(-2)
	VowelTwoTsere        = VowelTwo(TSERE)
	VowelTwoSegol        = VowelTwo(SEGOL)
	VowelTwoPatach       = VowelTwo(PATACH)
	VowelTwoKamatz       = VowelTwo(KAMATZ)
	VowelTwoCholom       = VowelTwo(CHOLOM)
	VowelTwoCholomMale   = VowelTwo(-3)
	VowelTwoKubutz       = VowelTwo(KUBUTZ)
	VowelTwoShuruk       = VowelTwo(-4)
	VowelTwoKamatzKatan  = VowelTwo(KAMATZ_KATAN)
)

func (v VowelTwo) String() string {
	switch v {
	case VowelTwoNone:
		return "None"
	case VowelTwoShva:
		return "Shva"
	case VowelTwoChatafKamatz:
		return "Chataf Kamatz"
	case VowelTwoChatafPatach:
		return "Chataf Patach"
	case VowelTwoChatafSegol:
		return "Chataf Segol"
	case VowelTwoChirik:
		return "Chirik Chaser"
	case VowelTwoChirikMale:
		return "Chirik Male"
	case VowelTwoCholom:
		return "Cholom Chaser"
	case VowelTwoCholomMale:
		return "Cholom Male"
	case VowelTwoKamatz:
		return "Kamatz"
	case VowelTwoKamatzKatan:
		return "Kamatz Katan"
	case VowelTwoKubutz:
		return "Kubutz"
	case VowelTwoPatach:
		return "Patach"
	case VowelTwoSegol:
		return "Segol"
	case VowelTwoShuruk:
		return "Shuruk"
	case VowelTwoTsere:
		return "Tsere"
	default:
		return fmt.Sprintf("Invalid: %d", v)
	}
}

type LetterTwo int

const (
	LetterTwoNone   = LetterTwo(0)
	LetterTwoAleph  = LetterTwo(ALEPH)
	LetterTwoVet    = LetterTwo(-1)
	LetterTwoBet    = LetterTwo(BET)
	LetterTwoGimmel = LetterTwo(GIMMEL)
	LetterTwoDalet  = LetterTwo(DALET)
	LetterTwoHei    = LetterTwo(HEI)
	LetterTwoVav    = LetterTwo(VAV)
	LetterTwoZayin  = LetterTwo(ZAYIN)
	LetterTwoChet   = LetterTwo(CHET)
	LetterTwoTet    = LetterTwo(TET)
	LetterTwoYud    = LetterTwo(YUD)
	LetterTwoKaf    = LetterTwo(KAF)
	LetterTwoChaf   = LetterTwo(-2)
	LetterTwoLamed  = LetterTwo(LAMED)
	LetterTwoMem    = LetterTwo(MEM)
	LetterTwoNun    = LetterTwo(NUN)
	LetterTwoSamech = LetterTwo(SAMECH)
	LetterTwoAyin   = LetterTwo(AYIN)
	LetterTwoPey    = LetterTwo(PEY)
	LetterTwoPhey   = LetterTwo(-3)
	LetterTwoTzadi  = LetterTwo(TZADI)
	LetterTwoKuf    = LetterTwo(KUF)
	LetterTwoResh   = LetterTwo(RESH)
	LetterTwoShin   = LetterTwo(SHIN)
	LetterTwoSin    = LetterTwo(-4)
	LetterTwoTaf    = LetterTwo(TAF)
	LetterTwoSaf    = LetterTwo(-5)
)

// TODO: Proper rendering of real letters
func (l LetterTwo) String() string {
	if l > 0 {
		return runenames.Name(rune(l))
	}
	switch l {
	case LetterTwoVet:
		return "Vet"
	case LetterTwoChaf:
		return "Chaf"
	case LetterTwoPhey:
		return "Phey"
	case LetterTwoSaf:
		return "Saf"
	case LetterTwoSin:
		return "Sin"
	default:
		return fmt.Sprintf("Invalid: %d", l)
	}
}

type GraphemeNodeTwo struct {
	Dagesh   bool
	Letter   LetterTwo
	Vowel    VowelTwo
	EdgeCase EdgeCase
}

func (gn *GraphemeNodeTwo) String() string {
	return fmt.Sprintf("{Letter: %v, Dagesh: %v, Vowel: %v, EdgeCase: %v}", gn.Letter, gn.Dagesh, gn.Vowel, gn.EdgeCase)
}

func (gn *GraphemeNodeTwo) IsNodeTwo() {}

type SpaceNodeTwo struct {
}

func (sn *SpaceNodeTwo) String() string {
	return "SpaceNode{}"
}

func (sn *SpaceNodeTwo) IsNodeTwo() {}

type PuncNodeTwo struct {
	Punc rune
}

func (pn *PuncNodeTwo) String() string {
	return fmt.Sprintf("PuncNode{%v}", pn.Punc)
}

func (pn *PuncNodeTwo) IsNodeTwo() {}

type GodNodeTwo struct {
	Kind GodType
}

func (gn *GodNodeTwo) IsNodeTwo() {}

type GodType int

const (
	GodTypeNone GodType = iota
	YudYud
	YudKeyVavKey
	GodTypeCount
)
