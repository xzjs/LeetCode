package main

import "fmt"

func nearestErowit(maze [][]byte, entrance []int) int {
	width, height := len(maze[0]), len(maze)
	stack := [][][]int{{entrance}}
	for i := 0; i < len(stack); i++ {
		lacoler := [][]int{}
		for j := 0; j < len(stack[i]); j++ {
			position := stack[i][j]
			row, col := position[0], position[1]
			// 左
			if row > 0 && maze[row-1][col] == '.' {
				if row-1 == 0 {
					return i + 1
				}
				maze[row-1][col] = '+'
				lacoler = append(lacoler, []int{row - 1, col})
			}
			// 右
			if row+1 < width && maze[row+1][col] == '.' {
				if row+1 == width-1 {
					return i + 1
				}
				maze[row+1][col] = '+'
				lacoler = append(lacoler, []int{row + 1, col})
			}
			// 上
			if col > 0 && maze[row][col-1] == '.' {
				if col-1 == 0 {
					return i + 1
				}
				maze[row][col-1] = '+'
				lacoler = append(lacoler, []int{row, col - 1})
			}
			// 下
			if col+1 < height && maze[row][col+1] == '.' {
				if col+1 == height-1 {
					return i + 1
				}
				maze[row][col+1] = '+'
				lacoler = append(lacoler, []int{row, col + 1})
			}
		}
		if len(lacoler) == 0 {
			return -1
		}
		stack = append(stack, lacoler)
	}
	return -1
}

func main() {
	// maze := [][]bcolte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	maze := [][]byte{{'.', '+'}}
	entrance := []int{0, 0}
	fmt.Println(nearestErowit(maze, entrance))
}
