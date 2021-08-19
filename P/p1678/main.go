package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	return strings.ReplaceAll(command, "(al)", "al")
}

func main() {
	command := "G()(al)"
	fmt.Println(interpret(command))
	command = "G()()()()(al)"
	fmt.Println(interpret(command))
	command = "(al)G(al)()()G"
	fmt.Println(interpret(command))
}
