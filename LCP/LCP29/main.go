package main

import (
	"fmt"
	"math"
)

func orchestraLayout(num int, xPos int, yPos int) int {
	layer := int(math.Min(math.Min(float64(xPos), float64(yPos)), math.Min(float64(num-1-xPos), float64(num-1-yPos)))) //当前圈数
	sider := num - 2*layer                                                                                             //小正方形的边长
	index := num*num - sider*sider                                                                                     //外层的数量
	num = sider
	xPos -= layer
	yPos -= layer
	if xPos == 0 {
		index += (yPos + 1)
	} else if yPos == num-1 {
		index += (num + xPos)
	} else if xPos == num-1 {
		index += ((num-1)*2 + num - yPos)
	} else if yPos == 0 {
		index += ((num-1)*3 + num - xPos)
	}
	result := index % 9
	if result == 0 {
		result = 9
	}
	return result
}

func main() {
	num, Xpos, Ypos := 3, 0, 2
	fmt.Println(orchestraLayout(num, Xpos, Ypos))
	num, Xpos, Ypos = 4, 1, 2
	fmt.Println(orchestraLayout(num, Xpos, Ypos))
	num, Xpos, Ypos = 2511, 1504, 2235
	fmt.Println(orchestraLayout(num, Xpos, Ypos))
}
