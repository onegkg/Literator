package types

import (
	"fmt"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

// NodeOne is the interface for all elements in the parsed text.
type NodeOne interface {
	IsNodeOne()
	String() string
}

type LinkedNodeOne struct {
	Node NodeOne
	Next *LinkedNodeOne
	Prev *LinkedNodeOne
}

func (ln1 *LinkedNodeOne) StringFromHead() string {
	var sb strings.Builder
	curr := ln1
	first := true
	fmt.Fprint(&sb, "[")
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

func (ln1 *LinkedNodeOne) String() string {
	return fmt.Sprintf("LN1{ %v }", ln1.Node)
}

func (ln1 *LinkedNodeOne) DebugString() string {
	return fmt.Sprintf("LN1{ Node: %v, Next: %v, Prev: %v}", ln1.Node, ln1.Next, ln1.Prev)
}

// ShinSinDot represents a dot on a Shin or Sin.
type ShinSinDot rune

const (
	DotNone ShinSinDot = 0
	DotShin ShinSinDot = ShinSinDot(SHINDOT)
	DotSin  ShinSinDot = ShinSinDot(SINDOT)
)

// GraphemeNodeOne represents a consonant cluster with its diacritics.
type GraphemeNodeOne struct {
	Letter        Letter
	Dagesh        bool
	Vowel         Vowel
	ShinSin       ShinSinDot
	IsChirikMaleh bool
}

func (gn *GraphemeNodeOne) IsNodeOne() {}

// SpaceNodeOne represents a structural space token.
type SpaceNodeOne struct{}

func (sn *SpaceNodeOne) IsNodeOne() {}

// PuncNodeOne represents punctuation like Maqaf.
type PuncNodeOne struct {
	Punc rune
}

func (pn *PuncNodeOne) IsNodeOne() {}

// --- Specific Letter Checks ---

func (gn *GraphemeNodeOne) IsShin() bool {
	return gn.Letter == SHIN && gn.ShinSin == DotShin
}

func (gn *GraphemeNodeOne) IsSin() bool {
	return gn.Letter == SHIN && gn.ShinSin == DotSin
}

func (gn *GraphemeNodeOne) String() string {
	repr := fmt.Sprintf("{Letter: %v, Dagesh: %v, Vowel: %v}", runenames.Name(rune(gn.Letter)), gn.Dagesh, runenames.Name(rune(gn.Vowel)))
	if gn.ShinSin != 0 {
		repr = fmt.Sprintf("{Letter: %v, Shin/SinDot: %v, Dagesh: %v, Vowel: %v}", runenames.Name(rune(gn.Letter)), runenames.Name(rune(gn.ShinSin)), gn.Dagesh, runenames.Name(rune(gn.Vowel)))
	}
	return repr
}

func (sn *SpaceNodeOne) String() string {
	return "SpaceNode{}"
}

func (pn *PuncNodeOne) String() string {
	return fmt.Sprintf("PunctuationNode{Punct: %c}", pn.Punc)
}
