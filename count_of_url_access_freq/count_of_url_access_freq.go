package main

import (
	"fmt"
	"strconv"
)

func mapFunCountOfURLAccessFreq(key, value string) map[string]string {
	var ans = make(map[string]string)

	ans[key] = "1"

	return ans
}

func reduceFunCountOfURLAccessFreq(key string, values []string) []string {
	var ans int = 0

	for _, value := range values {
		intVal, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		ans += intVal
	}

	return []string{strconv.Itoa(ans)}
}

func main() {
	// Sample access logs (URLs accessed by different users or at different times)
	urls := []string{
		"/home",
		"/about",
		"/home",
		"/contact",
		"/home",
		"/about",
		"/home",
	}

	// Map step: process each URL access log entry
	var mapResults []map[string]string
	for _, url := range urls {
		mapResult := mapFunCountOfURLAccessFreq(url, url)
		mapResults = append(mapResults, mapResult)
	}

	// Intermediate step: Aggregate results by URL
	intermediate := make(map[string][]string)
	for _, result := range mapResults {
		for url, count := range result {
			intermediate[url] = append(intermediate[url], count)
		}
	}

	// Reduce step: Summing counts for each URL
	for url, counts := range intermediate {
		finalCount := reduceFunCountOfURLAccessFreq(url, counts)
		fmt.Printf("URL: %s, Count: %s\n", url, finalCount[0])
	}
}
