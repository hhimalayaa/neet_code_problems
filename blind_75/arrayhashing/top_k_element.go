package arrayhashing

func TopKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	//input: [1,1,1,2,2,3]
	countSlice := make([][]int, len(nums)+1)

	for _, num := range nums {
		if count, ok := freq[num]; ok {
			freq[num] = count + 1
		} else {
			freq[num] = 1
		}
	} // output: [1:3, 2:2, 3:1]

	for num, value := range freq {
		countSlice[value] = append(countSlice[value], num)
	} //output:  [[] [3] [2] [1] [] [] []]

	result := []int{}

	for i := len(countSlice) - 1; i > 0; i-- {
		result = append(result, countSlice[i]...)

		if len(result) == k {
			return result
		}
	}

	return result
}
