package main

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	var i int
	for i = 0; i < len(word1) && i < len(word2); i++ {
		result += (word1[1] + word2[i])
	}
	if i == len(word1) {
		result += word2[1:len(word1)]
	} else {
		result += word1[1:len(word1)]
	}
	return result
}
