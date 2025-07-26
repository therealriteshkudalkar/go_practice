package problem5

import "math"

func NetworkDelayTime(times [][]int, n int, k int) int {
	type NodeWeight struct {
		node int
		weight int
	}

	type NodeDistance struct {
		node int
		distance int
	}

	// Form an adjacency list for traversing nodes
	adjList := make([][]NodeWeight, n)
	for _, time := range times {
		adjList[time[0] - 1] = append(adjList[time[0] - 1], NodeWeight{ time[1] - 1, time[2] })
	}

	// Create a visited array and a distance array
	visited := make([]bool, n)
	distances := make([]int, n)

	for i := range n {
		distances[i] = math.MaxInt
	}

	// Create a new empty queue
	queue := make([]NodeDistance, 0)
	queue = append(queue, NodeDistance{k - 1, 0})
	visited[k - 1] = true
	distances[k - 1] = 0

	for len(queue) != 0 {

	}

	min := math.MaxInt
	for index, distance := range distances {
		if distance == math.MaxInt {
			return -1
		}

		if distance < min && index != k - 1 {
			min = distance
		}
	}

	return min
}