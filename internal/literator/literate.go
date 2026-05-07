package literator

import (
	// "fmt"
	"fmt"
	"strings"

	"github.com/onegkg/literator/internal/types"
	"golang.org/x/text/unicode/norm"
)

// Preprocess takes a unicode string and returns a trimmed and NFD normalized version.
// All nikkud are broken out into separate runes and are ordered low to high by code point
func Preprocess(input string) string {
	// is there any other preprocessing we want to do?
	return norm.NFD.String(strings.TrimSpace(input))
}

type config struct {
	Builder  *strings.Builder
	Mappings types.Mappings
}

func Literate(head *types.LinkedNodeTwo) string {
	var sb strings.Builder
	c := config{
		Builder:  &sb,
		Mappings: types.DefaultMappings(),
	}
	curr := head
	for curr != nil {
		switch curr.Node.(type) {
		case *types.GraphemeNodeTwo:
			handleGraphemeNode(&c, curr)
		case *types.PuncNodeTwo:
			handlePuncNode(&c, curr)
		case *types.SpaceNodeTwo:
			handleSpaceNode(&c)
		case *types.GodNodeTwo:
			handleGodNode(&c, curr)
		default:
			panic("Inexhaustive handling of Node types in Literate")
		}
		curr = curr.Next
	}
	return sb.String()
}

func handleGodNode(c *config, node *types.LinkedNodeTwo) {
	godType := node.Node.(*types.GodNodeTwo).Kind
	switch godType {
	case types.YudYud:
		fmt.Fprint(c.Builder, "Adonai")
	case types.YudKeyVavKey:
		fmt.Fprint(c.Builder, "Adonai")
	default:
		panic("Inexhaustive handling of god types in handleGodNode")
	}
}

func handleSpaceNode(c *config) {
	fmt.Fprint(c.Builder, " ")
}

func handlePuncNode(c *config, node *types.LinkedNodeTwo) {
	puncNode := node.Node.(*types.PuncNodeTwo)
	fmt.Fprintf(c.Builder, "%c", puncNode.Punc)
}

func handleGraphemeNode(c *config, node *types.LinkedNodeTwo) {
	graphNode := node.Node.(*types.GraphemeNodeTwo)
	letterSound, ok := c.Mappings.Letters[graphNode.Letter]
	if !ok {
		panic("incomplete letter mapping")
	}
	vowelSound, ok := c.Mappings.Vowels[graphNode.Vowel]
	if !ok {
		panic("incomplete vowel mapping")
	}
	fmt.Fprintf(c.Builder, "%s%s", letterSound, vowelSound)
}
