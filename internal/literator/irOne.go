package literator

import (
	"unicode"

	"github.com/onegkg/literator/internal/types"
)

// ConvertToIROne Takes NFD normalized input and converts to slice of IR1 Nodes
// logs invalid characters as punctuation
func ConvertToIROne(input string) *types.LinkedNodeOne {
	var ir []types.NodeOne
	runes := []rune(input)
	max := len(runes)
	i := 0

	for i < max {
		if unicode.IsSpace(runes[i]) {
			ir = append(ir, &types.SpaceNodeOne{})
			if !incOrReturn(&i, max) {
				break
			}
			continue
		}

		var letter types.Letter = types.LetterNone
		dagesh := false
		var vowel types.Vowel = types.VowelNone
		var dot types.ShinSinDot = types.DotNone
		chirikMale := false

		// Skip invalid characters (non-hebrew characters or misplaced vowels)
		if !types.IsHebrewLetter(runes[i]) {
			// TODO: Skip create a punctuation list and only select from those, log the rest
			node := types.PuncNodeOne{
				Punc: runes[i],
			}
			ir = append(ir, &node)
			i++
			continue
		}

		letter = types.Letter(runes[i])
		if !incOrReturn(&i, max) {
			break
		}

		if types.IsHebrewVowel(runes[i]) {
			vowel = types.Vowel(runes[i])
			if !incOrReturn(&i, max) {
				break
			}
			if vowel == types.Vowel(types.CHIRIK) {
				if runes[i] == rune(types.YUD) {
					chirikMale = true
					if !incOrReturn(&i, max) {
						break
					}
				}
			}
		}

		if runes[i] == types.DAGESH {
			dagesh = true
			if !incOrReturn(&i, max) {
				break
			}
		}

		if types.IsShinDot(runes[i]) {
			dot = types.DotShin
			if !incOrReturn(&i, max) {
				break
			}
		} else if types.IsSinDot(runes[i]) {
			dot = types.DotSin
			if !incOrReturn(&i, max) {
				break
			}
		}

		node := types.GraphemeNodeOne{
			Letter:        letter,
			Dagesh:        dagesh,
			Vowel:         vowel,
			ShinSin:       dot,
			IsChirikMaleh: chirikMale,
		}
		ir = append(ir, &node)
	}
	return linkOne(ir)
}

// Turns a slice of NodeOnes into a linked list of LinkedNodeOnes. Returns the head of the linked list
func linkOne(slice []types.NodeOne) *types.LinkedNodeOne {
	var head *types.LinkedNodeOne
	var prev *types.LinkedNodeOne
	for i, node := range slice {
		if i == 0 {
			head = &types.LinkedNodeOne{
				Node: node,
				Prev: nil,
			}
			prev = head
		} else {
			newNode := &types.LinkedNodeOne{
				Node: node,
				Prev: prev,
			}
			prev.Next = newNode
			// fmt.Printf("prev.Next is nil: %v\n", prev.Next == nil)
			// fmt.Printf("prev: %v\n", prev.DebugPrint())
			// fmt.Printf("newNode: %v\n", newNode.DebugPrint())
			prev = newNode
		}
	}
	return head
}

// incOrReturn increases i by one, returns false if it then exceeds max
// used to safely increment an index without going over the length of the slice
func incOrReturn(i *int, max int) bool {
	*i++
	if *i >= max {
		return false
	}
	return true
}
