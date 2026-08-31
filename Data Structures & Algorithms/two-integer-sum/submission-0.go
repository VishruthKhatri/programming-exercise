func twoSum(nums []int, target int) []int {
    
	output := make(map[int]int)
	for i,n := range nums{
		sum := target-n
		if index, exists := output[sum]; exists{
			return []int{index,i}
		}
		output[n] = i
	}
	return nil
}
