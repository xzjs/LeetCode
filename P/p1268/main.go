package main

import "fmt"

var result []string

func suggestedProducts(products []string, searchWord string) [][]string {
	results := [][]string{}
	trie := Trie{}
	for _, product := range products {
		trie.Insert(product)
	}
	key := ""
	for _, word := range searchWord {
		key += string(word)
		result = []string{}
		node := trie.SearchPrefix(key)
		if node == nil {
			results = append(results, result)
			continue
		}
		node.DF(key)
		results = append(results, result)
	}
	return results
}

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(word string) *Trie {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) DF(current string) {
	if t.isEnd {
		if len(current) > 0 {
			result = append(result, current)
		}
		if len(t.children) == 0 {
			return
		}
	}
	for i := 0; i < 26 && len(result) < 3; i++ {
		if t.children[i] != nil {
			node := t.children[i]
			current1 := current + string(rune('a'+i))
			node.DF(current1)
		}
	}

}

func main() {
	// products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	// searchWord := "mouse"
	products := []string{"havana"}
	searchWord := "tatiana"
	results := suggestedProducts(products, searchWord)
	fmt.Println(results)
}
