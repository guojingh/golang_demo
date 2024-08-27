package draft

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Trie{
		root: &Node{next: make(map[rune]*Node)},
	}

	trie.Insert("AB")
	trie.Insert("ABC")
	trie.Insert("DF")
	trie.Insert("DH")
	trie.Insert("XY")
	Print(trie.root)
}

func Print(node *Node) {
	fmt.Printf("Node{isEnd:%t, next:[", node.isEnd)
	n := len(node.next)
	i := 0
	for k, v := range node.next {
		fmt.Printf("'%c': %p", k, v)
		i++
		if i < n {
			fmt.Printf(",")
		}
	}
	fmt.Println("]}")
	for _, v := range node.next {
		Print(v)
	}
}
