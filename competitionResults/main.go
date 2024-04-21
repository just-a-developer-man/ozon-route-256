package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Solution(in, out)
}

func Solution(in io.Reader, out io.Writer) {
	var inputDatasetsCount int
	fmt.Fscan(in, &inputDatasetsCount)

	for inputDatasetsCount != 0 {
		var sportsmenCount int
		fmt.Fscan(in, &sportsmenCount)

		if sportsmenCount == 1 {
			var result int
			fmt.Fscan(in, &result)
			fmt.Fprintln(out, 1)
			inputDatasetsCount--
			continue
		}

		results := make([]int, sportsmenCount)
		for counter := 0; counter < sportsmenCount; counter++ {
			result := 0
			fmt.Fscan(in, &result)
			results[counter] = result
		}

		newResults := make([]int, sportsmenCount)
		copy(newResults, results)

		places := NewFind(newResults)

		for _, value := range results {
			fmt.Fprint(out, places[value])
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)

		inputDatasetsCount--
	}
}

func NewFind(results []int) map[int]int {
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	places := make(map[int]int)

	place := 1
	nextPlace := 2

	for index, value := range results {
		if index == 0 {
			places[value] = place
			continue
		}
		if value-results[index-1] == 1 || value == results[index-1] {
			places[value] = place
			nextPlace += 1
		} else {
			place = nextPlace
			nextPlace += 1
			places[value] = place
		}
	}

	return places
}
