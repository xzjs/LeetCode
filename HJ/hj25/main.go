package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var INum int
	// if _, err := fmt.Scan(&INum); err != nil {
	// 	return
	// }
	INum = 15
	I := make([]string, INum)
	// for i := 0; i < INum; i++ {
	// 	if _, err := fmt.Scan(&I[i]); err != nil {
	// 		return
	// 	}
	// }
	I = []string{"123", "456", "786", "453", "46", "7", "5", "3", "665", "453456", "745", "456", "786", "453", "123"}
	// var RNum int
	// if _, err := fmt.Scan(&RNum); err != nil {
	// 	return
	// }
	// RNum = 5
	RMap := make(map[int]bool)
	// for i := 0; i < RNum; i++ {
	// 	var r int
	// 	if _, err := fmt.Scan(&r); err != nil {
	// 		return
	// 	}
	// 	RMap[r] = true //去重
	// }
	RMap = map[int]bool{6: true, 3: true, 0: true}
	R := make([]int, len(RMap))
	index := 0
	for k := range RMap {
		R[index] = k
		index++
	}
	sort.Ints(R) //排序

	res := make(map[int][]int)
	for _, v := range R {
		level := []int{}
		for i := 0; i < INum; i++ {
			if strings.Contains(I[i], strconv.Itoa(v)) {
				level = append(level, i)
			}
		}
		if len(level) > 0 {
			res[v] = level
		}
	}

	//处理输出
	total := len(res) * 2
	for _, v := range res {
		total += len(v) * 2
	}

	//输出
	fmt.Printf("%d ", total)
	for _, r := range R {
		if _, ok := res[r]; ok {
			fmt.Printf("%d %d ", r, len(res[r]))
			for _, v := range res[r] {
				fmt.Printf("%d %s ", v, I[v])
			}
		}
	}
}
