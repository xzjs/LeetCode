package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ { //从0出发
		g := make([]int, n) //剩余汽油
		for j := i + 1; ; j++ {
			j %= n //防止右溢出
			pre := j - 1
			if pre < 0 { //防止左溢出
				pre += n
			}
			g[j] = g[pre] + gas[pre] - cost[pre]
			if g[j] < 0 {
				break
			}
			if j == i {
				return i
			}
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
