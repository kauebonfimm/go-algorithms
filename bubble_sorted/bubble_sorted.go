package bubble_sorted

func BubbleSorted(nums []int) []int {
	for index := range nums {
		for subIndex := 0; subIndex < len(nums)-index-1; subIndex++ {
			if nums[subIndex] > nums[subIndex+1] {
				nums[subIndex], nums[subIndex+1] = nums[subIndex+1], nums[subIndex]
			}
		}
	}
	return nums
}
