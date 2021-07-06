package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	menu := map[string]int{}
	desks := make([]map[string]int, 501)
	for _, v := range orders {
		menu[v[2]] = 1
		no, _ := strconv.Atoi(v[1])
		if desks[no] != nil {
			desks[no][v[2]]++
		} else {
			desks[no] = map[string]int{
				v[2]: 1,
			}
		}
	}

	keys := make([]string, len(menu))
	index := 0
	for k := range menu {
		keys[index] = k
		index++
	}

	sort.Strings(keys)
	var result [][]string = [][]string{append([]string{"Table"}, keys...)}
	for i, v := range desks {
		if v != nil {
			s := strconv.Itoa(i)
			item := []string{s}
			for _, key := range keys {
				if num, ok := v[key]; ok {
					item = append(item, strconv.Itoa(num))
				} else {
					item = append(item, "0")
				}
			}
			result = append(result, item)
		}
	}
	return result
}

func print(result [][]string) {
	for _, v := range result {
		for _, v2 := range v {
			fmt.Print(v2 + " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	var orders [][]string = [][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}
	print(displayTable(orders))
	orders = [][]string{{"James", "12", "Fried Chicken"}, {"Ratesh", "12", "Fried Chicken"}, {"Amadeus", "12", "Fried Chicken"}, {"Adam", "1", "Canadian Waffles"}, {"Brianna", "1", "Canadian Waffles"}}
	print(displayTable(orders))
	orders = [][]string{{"Laura", "2", "Bean Burrito"}, {"Jhon", "2", "Beef Burrito"}, {"Melissa", "2", "Soda"}}
	print(displayTable(orders))
}
