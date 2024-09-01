package arrayhashing

func ProductExceptSelf(nums []int) []int {

	result := make([]int, len(nums))

	prefix := 1

	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	postfix := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
