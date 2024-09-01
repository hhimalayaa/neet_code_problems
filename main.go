package main

import (
	"fmt"

	"preparation/neet_code_problems/blind_75/arrayhashing"
)

func main() {

	// duplicate value
	duplicateValue := []int{1, 2, 3, 4, 5, 3, 6}
	isDuplicate := arrayhashing.ContainsDuplicate(duplicateValue)
	fmt.Println("result:", isDuplicate)

	//isAnagram
	firstString := "racecar"
	secondString := "carrace"
	IsAnagram := arrayhashing.IsAnagram(firstString, secondString)
	fmt.Println("result:", IsAnagram)

	//groupAnagram
	data := []string{"eat", "tea", "rat", "art", "pip", "hi"}
	groupAnagram := arrayhashing.GroupAnagrams(data)
	fmt.Println("result:", groupAnagram)

	// kfrequent element
	kfrequentElement := []int{1, 1, 1, 2, 2, 3}
	kFrequent := arrayhashing.TopKFrequent(kfrequentElement, 2)
	fmt.Println("result:", kFrequent)

	// encode & decode
	str := &arrayhashing.StringBuilder{}
	words := []string{"neat", "code", "go", "!34abc", "*#@asdp@3!"}
	encodeString := str.Encode(words)
	fmt.Println("encode:", encodeString)
	decode := str.Decode(encodeString)
	fmt.Println("result:", decode)

	selfArray := []int{1, 2, 3, 4}

	result := arrayhashing.ProductExceptSelf(selfArray)
	fmt.Println("result:", result)
}
