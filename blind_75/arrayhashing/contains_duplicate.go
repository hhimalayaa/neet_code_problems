package arrayhashing

func ContainsDuplicate(value []int) bool {
	numberMapping := make(map[int]int)

	if len(value) > 0 {
		for i := 0; i < len(value); i++ {
			if _, ok := numberMapping[value[i]]; ok {
				return true
			} else {
				numberMapping[value[i]]++
			}
		}
	}

	return false

}
