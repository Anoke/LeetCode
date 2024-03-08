package main

func dfs(current int, color []byte, graph [][]int, notSafe []bool) {
	color[current] = 'G'
	for _, ints := range graph[current] {
		if color[ints] == 'W' {
			dfs(ints, color, graph, notSafe)
		} else if color[ints] == 'G' {
			notSafe[ints] = true
		}
		notSafe[current] = notSafe[current] || notSafe[ints]
	}
	color[current] = 'B'
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	color := make([]byte, n)
	for i, _ := range graph {
		color[i] = 'W'
	}
	notSafe := make([]bool, n)
	var ans []int
	for i := 0; i < n; i++ {
		if color[i] == 'W' {
			dfs(i, color, graph, notSafe)
		}
	}
	for i := 0; i < n; i++ {
		if !notSafe[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
