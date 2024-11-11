package main

import "fmt"

type Pair struct {
	target string
	source string
}

// key is webpage url
// values are urls found on that webpage
func mapReverseWebLinkGraph(key string, values []string) []Pair {
	ans := make([]Pair, 0)
	for _, value := range values {
		ans = append(ans, Pair{
			source: key,
			target: value,
		})
	}

	return ans
}

func intermediateReverseWebLinkGraph(pairs []Pair) map[string][]string {
	reversedGraph := make(map[string][]string)

	for _, pair := range pairs {
		reversedGraph[pair.target] = append(reversedGraph[pair.target], pair.source)
	}

	return reversedGraph
}

func reduceReverseWebLinkGraph(key string, intermediateData map[string][]string) []string {
	return intermediateData[key]
}

func main() {
	// Example: web page links (key is a webpage, values are URLs found on that webpage)
	webPageLinks := map[string][]string{
		"/home":    {"/about", "/contact"},
		"/about":   {"/home", "/contact"},
		"/contact": {"/home"},
	}

	// Step 1: Map phase - process each webpage and create reverse links
	var mapResults []Pair
	for key, values := range webPageLinks {
		mapResults = append(mapResults, mapReverseWebLinkGraph(key, values)...)
	}

	// Step 2: Intermediate phase - group the reverse links by their target
	intermediateData := intermediateReverseWebLinkGraph(mapResults)

	fmt.Println(intermediateData)

	// Step 3: Reduce phase - aggregate the reversed web link graph and print results
	for target, _ := range intermediateData {
		reducedResults := reduceReverseWebLinkGraph(target, intermediateData)
		fmt.Printf("Target: %s, Sources: %v\n", target, reducedResults)
	}
}
