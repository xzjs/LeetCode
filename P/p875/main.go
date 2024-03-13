package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	var sum float64 = 0
	max := 0
	for _, v := range piles {
		sum += float64(v)
		if v > max {
			max = v
		}
	}
	min := int(math.Ceil(sum / float64(h)))
	if max == min {
		return max
	}
	results := []int{}
	k := 0
	for max >= min {
		k = (max + min) / 2
		t := 0
		for _, v := range piles {
			t += int(math.Ceil(float64(v) / float64(k)))
		}
		if t > h {
			min = k + 1
		} else if t < h {
			max = k
		} else {
			max--
			results = append(results, k)
		}
	}

	if len(results) > 0 {
		for _, v := range results {
			if v < k {
				k = v
			}
		}
	}
	return k
}

func main() {
	piles := []int{312884470}
	h := 312884469
	fmt.Println(minEatingSpeed(piles, h))
}
