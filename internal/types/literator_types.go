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

// <spacey> <yud> <yud> <spacey>
// <spacey> <yud> <hey> <vav> <hey> <spacey>

func isYud(node NodeTwo) bool {
	graphNode, ok := node.(*GraphemeNodeTwo)
	if !ok {
		return false
	}
	return graphNode.Letter == LetterTwoYud
}

func isHei(node NodeTwo) bool {
	graphNode, ok := node.(*GraphemeNodeTwo)
	if !ok {
		return false
	}
	return graphNode.Letter == LetterTwoHei
}

func isVav(node NodeTwo) bool {
	graphNode, ok := node.(*GraphemeNodeTwo)
	if !ok {
		return false
	}
	return graphNode.Letter == LetterTwoVav
}

// Returns true if head is the space before one of god's name
// TODO: Make it work with prefixes eg. BaHashem, KaHashem
func IsGod(head *LinkedNodeTwo) GodType {
	if GodTypeCount != 3 {
		panic("Inexhaustive handling of god types in IsGod()")
	}
	curr := head
	state := 0
	for curr != nil {
		switch state {
		case 0:
			if IsSpacey(curr.Node) {
				state = 1
			} else {
				return GodTypeNone
			}
		case 1:
			if isYud(curr.Node) {
				state = 2
			} else {
				return GodTypeNone
			}
		case 2:
			if isYud(curr.Node) {
				state = 3
			} else if isHei(curr.Node) {
				state = 4
			} else {
				return GodTypeNone
			}
		case 3:
			if IsSpacey(curr.Node) {
				return YudYud
			} else {
				return GodTypeNone
			}
		case 4:
			if isVav(curr.Node) {
				state = 5
			} else {
				return GodTypeNone
			}
		case 5:
			if isHei(curr.Node) {
				state = 6
			} else {
				return GodTypeNone
			}
		case 6:
			if IsSpacey(curr.Node) {
				return YudKeyVavKey
			} else {
				return GodTypeNone
			}
		}
		curr = curr.Next
	}
	return GodTypeNone
}

func IsSpacey(node NodeTwo) bool {
	switch node.(type) {
	case *SpaceNodeTwo:
		return true
	case *PuncNodeTwo:
		return true
	default:
		return false
	}
}

// Takes the node which is the space before god's name and returns it with the list modified so that god's name is replaced by a GodNode
func ReplaceGod(node *LinkedNodeTwo, kind GodType) *LinkedNodeTwo {
	if !IsSpacey(node.Node) {
		panic("non spacey head passed to ReplaceGod")
	}
	if kind == GodTypeNone {
		panic("Invalid god type passed to ReplaceGod")
	}

	if GodTypeCount != 3 {
		panic("inexhaustive handling of god types in ReplaceGod")
	}
	switch kind {
	case YudYud:
		replaceYudYud(node)
	case YudKeyVavKey:
		replaceYudKeyVavKey(node)
	}
	return node
}

// Accepts the space before an instance of god's YudYud name and replaces it with the appropriate GodNode
func replaceYudYud(head *LinkedNodeTwo) {
	for range 2 {
		removeNext(head)
	}
	godNode := &GodNodeTwo{
		Kind: YudYud,
	}
	insertAfter(head, godNode)
}

// Accepts the space before an instance of the tetragrammaton and replaces it with the appropriate GodNode
func replaceYudKeyVavKey(head *LinkedNodeTwo) {
	for range 4 {
		removeNext(head)
	}
	godNode := &GodNodeTwo{
		Kind: YudKeyVavKey,
	}
	insertAfter(head, godNode)
}

func removeNext(node *LinkedNodeTwo) (ok bool) {
	next := node.Next
	if next == nil {
		return false
	}
	node.Next = next.Next
	if next.Next != nil {
		next.Next.Prev = node
	}
	next.Next = nil
	next.Prev = nil
	return true
}

func insertAfter(node *LinkedNodeTwo, value NodeTwo) {
	next := node.Next
	newNode := &LinkedNodeTwo{
		Node: value,
		Next: next,
		Prev: node,
	}
	node.Next = newNode
	if next != nil {
		next.Prev = newNode
	}
}
