package arrayhashing

import (
	"strconv"
	"strings"
)

type StringBuilder struct {
	Str strings.Builder
}

func (s *StringBuilder) Encode(data []string) string {
	for _, v := range data {
		wordLength := strconv.Itoa(len(v))
		s.Str.WriteString(wordLength + "!" + v)
	}

	return s.Str.String()
}

func (s *StringBuilder) Decode(encoded string) []string {
	var result []string
	i := 0

	for i < len(encoded) {
		lenEnd := i

		for encoded[lenEnd] != '!' {
			lenEnd++
		}

		length, err := strconv.Atoi(encoded[i:lenEnd])

		if err != nil {
			return result
		}

		start := lenEnd + 1
		
		end := start + length

		if end > len(encoded) {
			break
		}

		result = append(result, encoded[start:end])

		i = end
	}

	return result
}
