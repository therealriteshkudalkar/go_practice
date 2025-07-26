package problem6

func LoudAndRich(richer [][]int, quiet []int) []int {
	// Create an adjacency list
	adjList := make([][]int, len(quiet))

	// Create an indegree array
	inDegrees := make([]int, len(quiet))

	// Fill the adjacency list and the in-degree array
	for _, rich := range richer {
		adjList[rich[0]] = append(adjList[rich[0]], rich[1])
		inDegrees[rich[1]] += 1
	}

	// Make and fill the answer array
	answer := make([]int, len(quiet))
	for i := range answer {
		answer[i] = i
	}

	// Put the ones with in-degree zero in the queue
	queue := make([]int, 0)
	for node, degree := range inDegrees {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// While queue is not empty
	for len(queue) != 0 {
		// dequeue from the queue
		currentNode := queue[0]
		queue = queue[1:]

		// Visit all it's neighbours
		for _, neighbor := range adjList[currentNode] {
			if quiet[answer[neighbor]] > quiet[answer[currentNode]] {
				// Update the answer
				answer[neighbor] = answer[currentNode]
			}
			inDegrees[neighbor] -= 1
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return answer
}
