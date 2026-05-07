package literator

import (
	"fmt"

	"github.com/onegkg/literator/internal/types"
)

var ec_debug bool = true

func ConvertToIRTwo(irOne *types.LinkedNodeOne) (*types.LinkedNodeTwo, error) {
	var slice []types.NodeTwo
	currNodeOne := irOne
	for currNodeOne != nil {
		node := currNodeOne.Node
		switch typedNode := node.(type) {
		case *types.SpaceNodeOne:
			currNode := types.SpaceNodeTwo{}
			slice = append(slice, &currNode)
		case *types.PuncNodeOne:
			currNode := types.PuncNodeTwo{Punc: typedNode.Punc}
			slice = append(slice, &currNode)
		case *types.GraphemeNodeOne:
			currNode, skip, err := newGraphemeNodeTwo(currNodeOne)
			if err != nil {
				return nil, err
			}

			if skip {
				currNodeOne = currNodeOne.Next
			}

			slice = append(slice, &currNode)
		default:
			panic("unexpected NodeOne type in ConvertToIRTwo")
		}
		currNodeOne = currNodeOne.Next
	}
	linked := linkTwo(slice)
	godReplaced := replaceGod(linked)
	return godReplaced, nil
}

func newGraphemeNodeTwo(node *types.LinkedNodeOne) (types.GraphemeNodeTwo, bool, error) {
	internalGrapheme, ok := node.Node.(*types.GraphemeNodeOne)
	if !ok {
		return types.GraphemeNodeTwo{}, false, fmt.Errorf("NOT A GRAPHEME NODE!")
	}

	letter, err := getLetter(*internalGrapheme)
	if err != nil {
		return types.GraphemeNodeTwo{}, false, err
	}

	vowel, skip := getVowel(node)

	ec, err := getEdgeCase(node)
	if err != nil {
		return types.GraphemeNodeTwo{}, false, err
	}

	return types.GraphemeNodeTwo{
		Letter:   letter,
		Dagesh:   internalGrapheme.Dagesh,
		Vowel:    vowel,
		EdgeCase: ec,
	}, skip, nil
}

func linkTwo(slice []types.NodeTwo) *types.LinkedNodeTwo {
	var head *types.LinkedNodeTwo
	var prev *types.LinkedNodeTwo

	for i, node := range slice {
		if i == 0 {
			head = &types.LinkedNodeTwo{
				Node: node,
				Prev: nil,
			}
			prev = head
		} else {
			newNode := &types.LinkedNodeTwo{
				Node: node,
				Prev: prev,
			}
			prev.Next = newNode
			prev = newNode
		}
	}
	// Add a space node to the end of the list
	newNode := &types.LinkedNodeTwo{
		Node: &types.SpaceNodeTwo{},
		Prev: prev,
		Next: nil,
	}
	prev.Next = newNode
	return head
}

func getLetter(node types.GraphemeNodeOne) (types.LetterTwo, error) {
	letter := node.Letter
	dagesh := node.Dagesh
	shinSin := node.ShinSin

	if letter.IsFinal() {
		return getLetterFinal(letter, dagesh), nil
	}

	if letter.IsBegadkefat() {
		return getLetterBegadkefat(letter, dagesh), nil
	}

	if shinSin != types.DotNone {
		if letter != types.SHIN {
			return types.LetterTwoNone, fmt.Errorf("Non shin has a shin/sin dot")
		}
		if shinSin == types.DotShin {
			return types.LetterTwoShin, nil
		} else {
			return types.LetterTwoSin, nil
		}
	}

	return types.LetterTwo(letter), nil
}

func getLetterFinal(letter types.Letter, dagesh bool) types.LetterTwo {
	switch letter {
	case types.KAF_SOFIT:
		if dagesh {
			return types.LetterTwoKaf
		} else {
			return types.LetterTwoChaf
		}
	case types.MEM_SOFIT:
		return types.LetterTwoMem
	case types.NUN_SOFIT:
		return types.LetterTwoNun
	case types.PEY_SOFIT:
		if dagesh {
			return types.LetterTwoPey
		} else {
			return types.LetterTwoPhey
		}
	case types.TZADI_SOFIT:
		return types.LetterTwoTzadi
	default:
		panic("getLetterFinal() failed")
	}
}

func getLetterBegadkefat(letter types.Letter, dagesh bool) types.LetterTwo {
	switch letter {
	case types.BET:
		if dagesh {
			return types.LetterTwoBet
		} else {
			return types.LetterTwoVet
		}
	case types.KAF:
		if dagesh {
			return types.LetterTwoKaf
		} else {
			return types.LetterTwoChaf
		}
	case types.PEY:
		if dagesh {
			return types.LetterTwoPey
		} else {
			return types.LetterTwoPhey
		}
	case types.TAF:
		if dagesh {
			return types.LetterTwoTaf
		} else {
			return types.LetterTwoSaf
		}
	default:
		// Sofiyot are handled in getLetterFinal
		panic("getLetterBegadkefat() failed")
	}
}

func getVowel(linkedNode *types.LinkedNodeOne) (vowel types.VowelTwo, skipNext bool) {
	node := linkedNode.Node.(*types.GraphemeNodeOne)
	chirikMaleh := node.IsChirikMaleh
	oldVowel := node.Vowel
	// This should probably be handled here but I've already implemented it and idc enough to change it now
	if chirikMaleh {
		return types.VowelTwoChirikMale, false
	}

	if oldVowel == types.VowelNone {
		if linkedNode.Next == nil {
			return types.VowelTwoNone, false
		}
		next := linkedNode.Next.Node
		switch typedNext := next.(type) {
		case *types.SpaceNodeOne:
		case *types.PuncNodeOne:
		case *types.GraphemeNodeOne:
			if typedNext.Letter == types.VAV {
				if typedNext.Vowel == types.CHOLOM {
					return types.VowelTwoCholomMale, true
				}
				if typedNext.Dagesh == true && typedNext.Vowel == types.VowelNone {
					return types.VowelTwoShuruk, true
				}
			}
		}
		return types.VowelTwoNone, false
	}
	return types.VowelTwo(oldVowel), false
}

// Currently not implemented, will implement once core functionality is complete
func getEdgeCase(linkedNode *types.LinkedNodeOne) (types.EdgeCase, error) {
	if ec_debug {
		ec_debug = false
		fmt.Println("Edgecase currently unimplemented")
	}
	return types.EdgeCaseNone, nil
}

// Scared of slowdown at higher n's since this is
func replaceGod(head *types.LinkedNodeTwo) *types.LinkedNodeTwo {
	curr := head
	for curr != nil {
		godType := types.IsGod(curr)
		if godType != types.GodTypeNone {
			types.ReplaceGod(curr, godType)
		}
		curr = findNextSpace(curr)
	}
	return head
}

func findNextSpace(head *types.LinkedNodeTwo) *types.LinkedNodeTwo {
	curr := head.Next
	for curr != nil {
		if types.IsSpacey(curr.Node) {
			return curr
		}
		curr = curr.Next
	}
	return nil
}
