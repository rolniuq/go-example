package main

import "fmt"

func bfs(adj [][]int, start int) []int {
	res := make([]int, 0)
	visited := make([]bool, len(adj))

	res = append(res, start)
	visited[start] = true

	for _, v := range adj {
		for _, w := range v {
			if !visited[w] {
				res = append(res, w)
				visited[w] = true
			}
		}
	}

	return res
}

func main() {
	adj := [][]int{
		{1, 2}, // 0 -> 1, 2
		{3, 4}, // 1 -> 3, 4
		{},     // 2 -> []
		{},     // 3 -> []
		{},     // 4 -> []
		{6},    // 5 -> 6
		{},     // 6 -> []
	}

	visited := make([]bool, len(adj))
	for i := range visited {
		visited[i] = false
	}

	res := bfs(adj, 0)

	fmt.Println(res)
}
