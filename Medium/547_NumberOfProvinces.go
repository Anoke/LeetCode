package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	proviences := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs_1(isConnected, visited, i)
			proviences++
		}
	}

	return proviences
}

func dfs_1(isConnected [][]int, visited []bool, city int) {
	visited[city] = true

	for neighbour := 0; neighbour < len(isConnected); neighbour++ {
		if !visited[neighbour] && isConnected[city][neighbour] == 1 {
			dfs_1(isConnected, visited, neighbour)
		}
	}
}
