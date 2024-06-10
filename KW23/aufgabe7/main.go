package main

import "fmt"

// TrieNode represents a single node in the Trie.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie represents the Trie structure with a root node.
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie.
func (t *Trie) Insert(word string) {

}

// Search checks if a word is in the Trie.
func (t *Trie) Search(word string) bool {

}

// StartsWith checks if any word in the Trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {

}

// GetWordsWithPrefix returns all words in the Trie that start with the given prefix.
func (t *Trie) GetWordsWithPrefix(prefix string) []string {

}

// collectWords is a helper function to perform a DFS and collect words from the given node.
func (t *Trie) collectWords(node *TrieNode, currentWord string, words *[]string) {

}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot", "banana", "band", "bandana", "cat", "car", "cart"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Testing Search and StartsWith
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: true
	fmt.Println(trie.Search("apricot")) // Output: true
	fmt.Println(trie.Search("banana"))  // Output: true
	fmt.Println(trie.Search("bananas")) // Output: false
	fmt.Println(trie.StartsWith("ap"))  // Output: true
	fmt.Println(trie.StartsWith("ba"))  // Output: true
	fmt.Println(trie.StartsWith("ca"))  // Output: true
	fmt.Println(trie.StartsWith("da"))  // Output: false

	// Testing GetWordsWithPrefix
	fmt.Println(trie.GetWordsWithPrefix("ap"))  // Output: [apple app apricot]
	fmt.Println(trie.GetWordsWithPrefix("ba"))  // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("ca"))  // Output: [cat car cart]
	fmt.Println(trie.GetWordsWithPrefix("ban")) // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("xyz")) // Output: []
}

/*
package main

import "fmt"

// TrieNode represents a single node in the Trie.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie represents the Trie structure with a root node.
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		// If the character is not present, create a new node.
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true // Mark the end of the word.
}

// Search checks if a word is in the Trie.
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		// If the character is not found, the word does not exist in the Trie.
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd // Return true only if the word ends here.
}

// StartsWith checks if any word in the Trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		// If the character is not found, no words with the given prefix exist.
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true // If we've traversed the whole prefix, return true.
}

// GetWordsWithPrefix returns all words in the Trie that start with the given prefix.
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	node := t.root
	var words []string
	var currentWord string

	// Navigate to the node at the end of the prefix
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return words // Return empty if prefix not found
		}
		node = node.children[char]
		currentWord += string(char)
	}

	// Collect all words from this node onwards
	t.collectWords(node, currentWord, &words)
	return words
}

// collectWords is a helper function to perform a DFS and collect words from the given node.
func (t *Trie) collectWords(node *TrieNode, currentWord string, words *[]string) {
	if node.isEnd {
		*words = append(*words, currentWord)
	}
	for char, childNode := range node.children {
		t.collectWords(childNode, currentWord+string(char), words)
	}
}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot", "banana", "band", "bandana", "cat", "car", "cart"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Testing Search and StartsWith
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: true
	fmt.Println(trie.Search("apricot")) // Output: true
	fmt.Println(trie.Search("banana"))  // Output: true
	fmt.Println(trie.Search("bananas")) // Output: false
	fmt.Println(trie.StartsWith("ap"))  // Output: true
	fmt.Println(trie.StartsWith("ba"))  // Output: true
	fmt.Println(trie.StartsWith("ca"))  // Output: true
	fmt.Println(trie.StartsWith("da"))  // Output: false

	// Testing GetWordsWithPrefix
	fmt.Println(trie.GetWordsWithPrefix("ap"))  // Output: [apple app apricot]
	fmt.Println(trie.GetWordsWithPrefix("ba"))  // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("ca"))  // Output: [cat car cart]
	fmt.Println(trie.GetWordsWithPrefix("ban")) // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("xyz")) // Output: []
}

*/
