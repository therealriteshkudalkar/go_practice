package problem5

import (
	"container/heap"
	"math"
)

type nodeWeight struct {
	node   int
	weight int
}

type nodeDistance struct {
	node     int
	distance int
}

type priorityQueue []*nodeDistance

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].distance > (*pq)[j].distance
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(*nodeDistance))
}

func (pq *priorityQueue) Peek() any {
	n := pq.Len()
	return (*pq)[n-1]
}

func (pq *priorityQueue) Pop() any {
	n := pq.Len()
	x := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return x
}

func NetworkDelayTime(times [][]int, n int, k int) int {

	// Form an adjacency list for traversing nodes
	adjList := make([][]nodeWeight, n)
	for _, time := range times {
		adjList[time[0]-1] = append(adjList[time[0]-1], nodeWeight{time[1] - 1, time[2]})
	}

	// Create a visited array and a distance array
	distances := make([]int, n)

	for i := range n {
		distances[i] = math.MaxInt
	}

	// Create a new empty queue
	queue := priorityQueue{}
	heap.Init(&queue)
	heap.Push(&queue, &nodeDistance{k - 1, 0})
	distances[k-1] = 0

	for len(queue) != 0 {
		// Remove the item from the queue
		currentNode := heap.Pop(&queue).(*nodeDistance)

		// Explore the neighbours of this item
		for _, neighbour := range adjList[currentNode.node] {
			if distances[neighbour.node] > currentNode.distance+neighbour.weight {
				distances[neighbour.node] = currentNode.distance + neighbour.weight
				heap.Push(&queue, &nodeDistance{neighbour.node, distances[neighbour.node]})
			}
		}
	}

	max := 0
	for index, distance := range distances {
		if distance == math.MaxInt {
			return -1
		}

		if distance > max && index != k-1 {
			max = distance
		}
	}

	return max
}
