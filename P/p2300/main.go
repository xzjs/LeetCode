package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	size := len(potions)
	sort.Ints(potions)
	var pairs []int = make([]int, len(spells))
	for k, spell := range spells {
		quotient := float64(success) / float64(spell)
		ans := sort.Search(size, func(i int) bool { return float64(potions[i]) >= quotient })
		pairs[k] = size - ans
	}
	return pairs
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	var success int64 = 7
	fmt.Println(successfulPairs(spells, potions, success))
}
