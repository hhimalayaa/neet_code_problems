package arrayhashing

func GroupAnagrams(strs []string) [][]string {
	anagram := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		anagram[count] = append(anagram[count], s)
	}

	result := make([][]string, 0)

	for _, value := range anagram {
		result = append(result, value)
	}

	return result

}
