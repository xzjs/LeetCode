package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationMap := make(map[string]map[string]float64)
	for key, equation := range equations {
		if _, ok := equationMap[equation[0]]; !ok {
			equationMap[equation[0]] = make(map[string]float64)
		}
		equationMap[equation[0]][equation[1]] = values[key]
		if _, ok := equationMap[equation[1]]; !ok {
			equationMap[equation[1]] = make(map[string]float64)
		}
		equationMap[equation[1]][equation[0]] = 1 / values[key]
	}
	stack := make(map[string]float64)
	isFind := false
	var bfs func(numerator, denominator string)
	bfs = func(numerator, denominator string) {
		if _, ok := equationMap[numerator][denominator]; ok {
			stack[denominator] = equationMap[numerator][denominator]
			isFind = true
		} else {
			for key, value := range equationMap[numerator] {
				if _, ok := stack[key]; ok {
					continue
				}
				stack[numerator] = value
				bfs(key, denominator)
				if !isFind {
					delete(stack, numerator)
				} else {
					return
				}
			}
		}
	}
	answers := []float64{}
	for _, query := range queries {
		stack = make(map[string]float64)
		isFind = false
		_, ok1 := equationMap[query[0]]
		_, ok2 := equationMap[query[1]]
		if !ok1 || !ok2 {
			answers = append(answers, -1.0)
			continue
		}
		if query[0] == query[1] {
			answers = append(answers, 1.0)
			continue
		}
		bfs(query[0], query[1])
		answer := -1.0
		if isFind {
			answer = 1
			for _, v := range stack {
				answer *= v
			}
		}
		answers = append(answers, answer)
	}
	return answers
}

func main() {
	equations := [][]string{{"a", "b"}, {"c", "d"}}
	values := []float64{1.0, 1.0}
	queries := [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}}
	answers := calcEquation(equations, values, queries)
	for _, answer := range answers {
		println(answer)
	}

}
