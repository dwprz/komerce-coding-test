package task

import (
	"bufio"
	"fmt"
	"strings"
)

type CharInfo struct {
	Accumulated string
	Index       int
}

func SortCharacters(reader *bufio.Reader) {
	fmt.Print("Enter a word or sentence: ")
	words, _ := reader.ReadString('\n')
	words = strings.TrimSpace(words)

	result := sortCharactersHelper(words)
	fmt.Println(result)
}

func sortCharactersHelper(words string) string {
	characters := strings.ToLower(strings.ReplaceAll(words, " ", ""))

	vowelMap := make(map[rune]*CharInfo)
	vowelIndex := 0

	consonantMap := make(map[rune]*CharInfo)
	consonantIndex := 0

	for _, char := range characters {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			if vowelMap[char] == nil {
				vowelMap[char] = &CharInfo{
					Index: vowelIndex,
				}
				vowelIndex++
			}
			vowelMap[char].Accumulated += string(char)

		} else {

			if consonantMap[char] == nil {
				consonantMap[char] = &CharInfo{
					Index: consonantIndex,
				}
				consonantIndex++
			}
			consonantMap[char].Accumulated += string(char)
		}
	}

	vowelOutput := make([]string, len(vowelMap))
	for _, info := range vowelMap {
		vowelOutput[info.Index] = info.Accumulated
	}

	consonantOutput := make([]string, len(consonantMap))
	for _, info := range consonantMap {
		consonantOutput[info.Index] = info.Accumulated
	}

	return fmt.Sprintf("Vowel Characters : %s \nConsonant Characters : %s", strings.Join(vowelOutput, ""), strings.Join(consonantOutput, ""))
}