package draft

type Node struct {
	isEnd bool           // 当前字符串是否是单词结尾
	next  map[rune]*Node // 子节点存储后续的字符
}

type Trie struct {
	root *Node
}

// 前缀树的插入
func (trie *Trie) Insert(word string) {
	cur := trie.root
	for _, char := range []rune(word) {
		if _, ok := cur.next[char]; !ok {
			cur.next[char] = &Node{next: make(map[rune]*Node)}
		}
		cur = cur.next[char]
	}
	cur.isEnd = true
}

// 前缀树的查找
func (trie *Trie) Search(word string) bool {
	cur := trie.root
	for _, char := range []rune(word) {
		if _, ok := cur.next[char]; !ok {
			return false
		}
		cur = cur.next[char]
	}
	return cur.isEnd
}
