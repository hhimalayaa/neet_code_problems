package arrayhashing

func TwoSum(nums []int, target int) []int {
	numberMapping := make(map[int]int, 0)

	for i, num := range nums {
		j := target - num

		if value, ok := numberMapping[j]; ok {

			return []int{value, i}
		} else {
			numberMapping[num] = i
		}
	}
	
	return []int{}
}
