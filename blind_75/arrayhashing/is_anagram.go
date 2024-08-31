package arrayhashing

import "reflect"

func IsAnagram(s, t string) bool {

	firstWord := make(map[rune]int, len(s))
	secondWord := make(map[rune]int, len(t))

	for _, value := range s {
		firstWord[value]++
	}

	for _, value := range t {
		secondWord[value]++
	}

	return reflect.DeepEqual(firstWord, secondWord)

}
