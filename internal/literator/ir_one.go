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

		letter := types.LetterNone
		dagesh := false
		vowel := types.VowelNone
		dot := types.DotNone
		chirikMale := false

		// Skip invalid characters (non-hebrew characters or misplaced vowels)
		if !types.IsHebrewLetter(runes[i]) {
			// TODO: Create a punctuation list and only select from those, log the rest
			node := types.PuncNodeOne{
				Punc: runes[i],
			}
			ir = append(ir, &node)
			i++
			continue
		}

		letter = types.Letter(runes[i])
		notEmpty := incOrReturn(&i, max)

		if notEmpty && types.IsHebrewVowel(runes[i]) {
			vowel = types.Vowel(runes[i])
			notEmpty = incOrReturn(&i, max)
			if notEmpty && vowel == types.Vowel(types.CHIRIK) {
				if runes[i] == rune(types.YUD) {
					chirikMale = true
					notEmpty = incOrReturn(&i, max)
				}
			}
		}

		if notEmpty && runes[i] == types.DAGESH {
			dagesh = true
			notEmpty = incOrReturn(&i, max)
		}

		if notEmpty {
			if types.IsShinDot(runes[i]) {
				dot = types.DotShin
				notEmpty = incOrReturn(&i, max) // technically not necessary, but done for consistency
			} else if types.IsSinDot(runes[i]) {
				dot = types.DotSin
				notEmpty = incOrReturn(&i, max)
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
