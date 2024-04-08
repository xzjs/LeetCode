package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		dict := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if _, ok := dict[board[i][j]]; ok && board[i][j] != '.' {
				return false
			}
			dict[board[i][j]] = true
		}
	}
	for i := 0; i < 9; i++ {
		dict := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if _, ok := dict[board[j][i]]; ok && board[j][i] != '.' {
				return false
			}
			dict[board[i][j]] = true
		}
	}
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			dict := make(map[byte]bool)
			for k := i; k < 3; k++ {
				for l := j; l < 3; l++ {
					if _, ok := dict[board[k][l]]; ok && board[k][l] != '.' {
						return false
					}
					dict[board[k][l]] = true
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}
