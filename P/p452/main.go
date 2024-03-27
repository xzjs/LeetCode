package main

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] == points[j][0] {
			if points[i][1] < points[j][1] {
				return true
			}
		}
		return false
	})
	newPoints := [][]int{points[0]}
	r := points[0][1]
	for _, v := range points[1:] {
		if v[0] >= r {
			r = v[1]
			newPoints = append(newPoints, v)
		} else if v[1] < r { //有重叠
			r = v[1]
		}
	}
	return len(newPoints)
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	println(findMinArrowShots(points))
}
