package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	arr := []string{}
	for _, v := range dirs {
		switch v {
		case "":
		case ".":
		case "..":
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}
		default:
			arr = append(arr, v)
		}
	}
	return string(rune('/')) + strings.Join(arr, "/")
}

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}
