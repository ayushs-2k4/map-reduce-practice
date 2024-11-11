package main

import (
	"fmt"
	"strconv"
)

func getWords(sentence string) []string {
	words := make([]string, 0)

	i := 0
	for j := 0; j < len(sentence); j++ {
		if sentence[j] == ' ' {
			words = append(words, sentence[i:j])
			i = j + 1
		}
	}

	words = append(words, sentence[i:])

	return words
}

// key: document name
// value: document contents
func mapFun(key, value string) map[string]string {

	words := getWords(value)

	ans := make(map[string]string)

	for _, word := range words {
		intVal, _ := strconv.Atoi(ans[word])
		ans[word] = strconv.Itoa(intVal + 1)
	}

	return ans
}

func intermediateStep(mapResult []map[string]string) map[string][]string {
	intermediate := make(map[string][]string)
	for _, result := range mapResult {
		for word, count := range result {
			intermediate[word] = append(intermediate[word], count)
		}
	}

	return intermediate
}

func reduceFun(key string, values []string) []string {
	result := 0
	for _, v := range values {
		intVal, _ := strconv.Atoi(v)
		result += intVal
	}

	return []string{strconv.Itoa(result)}
}

func main() {
	sentences := []string{
		"Hi Hello Hi World",
		"Hello World",
		"Hi there World",
	}

	var allMapResults []map[string]string

	for _, sentence := range sentences {
		// Perform the map step
		mapResult := mapFun("file_name", sentence)
		allMapResults = append(allMapResults, mapResult)
	}

	// Intermediate step to group by word
	reduceData := intermediateStep(allMapResults)

	// Perform the reduce step and print results
	for word, counts := range reduceData {
		finalCount := reduceFun(word, counts)
		fmt.Printf("Word: %s, Count: %s\n", word, finalCount[0])
	}
}
