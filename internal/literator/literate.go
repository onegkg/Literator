package literator

import (
	// "fmt"
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
	Builder *strings.Builder
}

func Literate(nodes []types.NodeOne) string {
	var b strings.Builder
	c := config{
		Builder: &b,
	}
	for i, node := range nodes {
		switch n := node.(type) {
		case *types.GraphemeNodeOne:
			handleGraphemeNode(&c, nodes, i)
		case *types.PuncNodeOne:
			b.WriteRune(n.Punc)
		case *types.SpaceNodeOne:
			b.WriteRune(' ')
		}
	}
	return b.String()
}

func handleGraphemeNode(c *config, nodes []types.NodeOne, idx int) {
	// node := nodes[idx].(*types.GraphemeNodeOne)
	// letter := node.Letter
	// fmt.Fprintf(c.Builder, "%v", letterSound)
}
