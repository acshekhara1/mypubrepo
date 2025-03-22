package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// NewTrieNode() is a contructor function
func NewTrieNode() *TrieNode {
	return &TrieNode{children : make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

// Constructor to create a trie node.
func constructor() *Trie {
	t := new(Trie)
    t.root = NewTrieNode()
    return t
}

// A function to insert a word in trie.
func (this *Trie) insert(word string) {
	node := this.root
	
	for _, chr := range word {
		if _, ok := node.children[chr]; !ok {
			node.children[chr] = NewTrieNode()
		} 
		node = node.children[chr] 
	}

	node.isWord = true 
}

// A function to search for a word in the trie.
func (this *Trie) search(word string) bool {
	node := this.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		} 
		node = node.children[c] 
	}
	return node.isWord 
}

// A function to search for a prefix of a word in the trie.
func (this *Trie) searchPrefix(prefix string) bool {
	
	node := this.root
	
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		} 
		node = node.children[c] 
	}
	return true
}


// Driver Code
func main() {
	keys := []string{"the", "a", "there", "answer"}
	trieForKeys := new(Trie)
	trieForKeys.root = NewTrieNode()
	num := 1

	for _, x := range keys {
		fmt.Printf("%d.\tInserting key: \"%s\"", num, x)
		trieForKeys.insert(x)
		num += 1
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}

	searching := []string{"a", "answer", "xyz", "an"}
	for _, y := range searching {
		fmt.Printf("%d.\tSearching key: \"%s\"", num, y)
		fmt.Printf("\n\tKey found? %v\n", trieForKeys.search(y))
		num += 1
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}

	searchingPrefix := []string{"b", "an"}
	for _, z := range searchingPrefix {
		fmt.Printf("%d.\tSearching prefix: \"%s\"", num, z)
		fmt.Printf("\n\tPrefix Found? %v\n", trieForKeys.searchPrefix(z))
		num += 1
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
