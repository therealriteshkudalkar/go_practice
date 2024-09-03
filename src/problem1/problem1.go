package problem1

// TwoSum method solves the popular two sum problem
func TwoSum(nums []int, target int) []int {
	memo := map[int]int{}
	for index, num := range nums {
		currentComplementIndex, ok := memo[target-num]
		if ok {
			return []int{currentComplementIndex, index}
		}
		memo[num] = index
	}
	return []int{}
}
