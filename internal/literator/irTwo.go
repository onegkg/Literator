package literator

import (
	"fmt"

	"github.com/onegkg/literator/internal/types"
)

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
			currNode, err := newGraphemeNodeTwo(currNodeOne)
			if err != nil {
				return nil, err
			}
			slice = append(slice, &currNode)
		default:
			panic("unexpected NodeOne type in ConvertToIRTwo")
		}
		currNodeOne = currNodeOne.Next
	}
	linked := linkTwo(slice)
	return linked, nil
}

func newGraphemeNodeTwo(node *types.LinkedNodeOne) (types.GraphemeNodeTwo, error) {
	internalGrapheme, ok := node.Node.(*types.GraphemeNodeOne)
	if !ok {
		return types.GraphemeNodeTwo{}, fmt.Errorf("NOT A GRAPHEME NODE!")
	}

	letter, err := getLetter(*internalGrapheme)
	if err != nil {
		return types.GraphemeNodeTwo{}, err
	}
	letter = letter
	panic("")
}

func linkTwo(slice []types.NodeTwo) *types.LinkedNodeTwo {
	panic("Not yet implemented")
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
		panic("getLetterBegadkefat() failed")
	}
}

// func getVowel(node types.GraphemeNodeOne) (types.VowelTwo, error) {
// 	chirikMaleh := node.IsChirikMaleh
// 	vowel := node.Vowel
// 	if chirikMaleh {
// 		return types.VowelTwoChirikMale, nil
// 	}
// 	switch vowel {
// 		case types.CHOLOM:
// 			if
// 	}

// }
