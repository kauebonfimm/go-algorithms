package two_numbers

func TwoSum(nums []int, target int) []int {
	hashNums := map[int]int{}
	for i, v := range nums {
		if n, ok := hashNums[target-v]; ok {
			return []int{i, n}
		}

		hashNums[v] = i
	}
	return []int{}
}
