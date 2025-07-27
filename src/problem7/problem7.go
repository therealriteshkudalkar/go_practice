package problem7

func CanVisitAllRooms(rooms [][]int) bool {
	// Create a queue
	queue := make([]int, 0)

	// Create a set of visited nodes
	visited := make(map[int]bool)

	// Visit room zero
	queue = append(queue, 0)
	visited[0] = true

	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, roomKey := range rooms[currentNode] {
			if visitVal, ok := visited[roomKey]; !visitVal || !ok {
				visited[roomKey] = true
				queue = append(queue, roomKey)
			}
		}
	}

	return len(visited) == len(rooms)
}
