package trie

// TODO:
// Handle all kinds of possible alphabets
// Currently only the English alphabet is handled
const CharactersCount int = 26

type Node struct {
	Children        [CharactersCount]*Node
	IsLastCharacter bool
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{Root: &Node{}}
}

func (t *Trie) Add(word string) {
	currentNode := t.Root

	for _, c := range word {
		// TODO:
		// Handle differentiating between different alphabets
		// and get the proper character index
		charIdx := c - 'a'

		if currentNode.Children[charIdx] == nil {
			currentNode.Children[charIdx] = &Node{}
		}

		currentNode = currentNode.Children[charIdx]
	}

	currentNode.IsLastCharacter = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.Root

	for _, c := range word {
		// TODO:
		// Handle differentiating between different alphabets
		// and get the proper character index
		charIdx := c - 'a'

		if currentNode.Children[charIdx] == nil {
			return false
		}

		currentNode = currentNode.Children[charIdx]
	}

	return currentNode.IsLastCharacter
}
