package problem8

import "fmt"

func find(parent map[string]string, item string) string {
	if parent[item] != item {
		parent[item] = find(parent, parent[item])
	}
	return parent[item]
}

func union(parent map[string]string, rank map[string]int, x string, y string) {
	parentX := find(parent, x)
	parentY := find(parent, y)

	if parentX == parentY {
		return
	}

	if rank[parentX] < rank[parentY] {
		parent[parentX] = parentY
	} else if rank[parentX] > rank[parentY] {
		parent[parentY] = parentX
	} else {
		parent[parentX] = parentY
		rank[parentY] += 1
	}
}

func EquationsPossible(equations []string) bool {
	// Store parent and ranks
	parent := make(map[string]string)
	rank := make(map[string]int)

	// Store just inequalities
	inequalities := make([][]string, 0)

	// Find all unique variables
	for _, equation := range equations {
		lhsOperand := equation[0:1]
		rhsOperand := equation[3:]
		equality := equation[1:2]

		if _, ok := parent[lhsOperand]; !ok {
			parent[lhsOperand] = lhsOperand
			rank[lhsOperand] = 0
		}
		if _, ok := parent[rhsOperand]; !ok {
			parent[rhsOperand] = rhsOperand
			rank[rhsOperand] = 0
		}

		// Check if lhsOperand is in the set, if yes, then what is its index
		if equality == "=" {
			union(parent, rank, lhsOperand, rhsOperand)
		} else {
			inequalities = append(inequalities, []string{lhsOperand, rhsOperand})
		}
	}

	fmt.Print(parent, rank)

	for _, inequality := range inequalities {
		lhsOperand := inequality[0]
		rhsOperand := inequality[1]

		if find(parent, lhsOperand) == find(parent, rhsOperand) {
			return false
		}
	}

	return true
}
