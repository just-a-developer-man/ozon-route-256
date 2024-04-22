package main

import (
	"bufio"
	"fmt"
	"os"
)

type ssStat struct {
	redCount      int
	blueCount     int
	inappropriate bool
	black         bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var datasets int
	fmt.Fscan(in, &datasets)

	for datasets != 0 {
		var wordsCount, blueWordsCount, redWordsCount, blackWordIndex int
		fmt.Fscan(in, &wordsCount, &blueWordsCount, &redWordsCount, &blackWordIndex)

		substringsStats := make(map[string]ssStat)

		words := make([]string, wordsCount)
		for index := range words {
			var word string
			fmt.Fscan(in, &word)
			words[index] = word
		}

		// add white strings to map and them as inappropriate
		GenerateWhiteSubstrings(words[blueWordsCount+redWordsCount:], substringsStats)

		// add black substrings to map and mark them as black
		blackSubstrings := GenerateBlackSubstrings(words[blackWordIndex-1])
		for _, substring := range blackSubstrings {
			substringsStats[substring] = ssStat{black: true}
		}
		var ansPresent bool
		for wordIndex, word := range words[:blueWordsCount+redWordsCount] {
			substrings := GenerateBlueRedSubstrings(word)

			if len(substrings) < 1 {
				continue
			}
			for substring := range substrings {
				if checkSubstringIsBlack(substring, substringsStats) {
					continue
				}
				if checkSubstringAppropriate(substring, substringsStats) {
					ansPresent = true
					if substringStat, ok := substringsStats[substring]; ok {
						if wordIndex < blueWordsCount {
							substringStat.blueCount++
							substringsStats[substring] = substringStat
						} else {
							substringStat.redCount++
							substringsStats[substring] = substringStat
						}
					} else {
						substringStat = ssStat{}
						if wordIndex < blueWordsCount {
							substringStat.blueCount++
							substringsStats[substring] = substringStat
						} else {
							substringStat.redCount++
							substringsStats[substring] = substringStat
						}
					}
				}
			}
		}
		if ansPresent {
			ansSubstr, ansDiff := findMaxDiff(substringsStats)
			fmt.Fprintln(out, ansSubstr, ansDiff)
		} else {
			fmt.Fprintln(out, "lkjfslkjf", 0)
		}

		datasets--
	}

}

func GenerateBlueRedSubstrings(in string) map[string]struct{} {
	if len(in) <= 1 {
		return nil
	}
	substrings := make(map[string]struct{})
	for index := 0; index < len(in); index++ {
		for jindex := index; jindex < len(in); jindex++ {
			if jindex == len(in)-1 && index == 0 {
				continue
			}
			substrings[in[index:jindex+1]] = struct{}{}
		}
	}
	return substrings
}

func GenerateBlackSubstrings(in string) []string {
	substrings := make([]string, 0, calculateSubstrings(in)+1)
	for index := 0; index < len(in); index++ {
		for jindex := index; jindex < len(in); jindex++ {
			substrings = append(substrings, in[index:jindex+1])
		}
	}
	return substrings
}

func GenerateWhiteSubstrings(in []string, substringsStats map[string]ssStat) {
	for _, substring := range in {
		substringsStats[substring] = ssStat{inappropriate: true}
	}
}

func calculateSubstrings(in string) int {
	var out int
	for count := len(in); count >= 1; count-- {
		out += count
	}
	return out - 1
}

func checkSubstringAppropriate(substring string, substringsStats map[string]ssStat) bool {
	if _, ok := substringsStats[substring]; ok {
		if substringsStats[substring].inappropriate == true {
			return false
		}
	}
	return true
}

func checkSubstringIsBlack(substring string, substringsStats map[string]ssStat) bool {
	if substringStat, ok := substringsStats[substring]; ok {
		if substringStat.black == true {
			return true
		}
	}
	return false
}

func findMaxDiff(substringsStats map[string]ssStat) (string, int) {
	ansSubstr := ""
	ansDiff := 0
	fmt.Println(substringsStats)
	for substr, stat := range substringsStats {
		if stat.inappropriate || stat.black {
			continue
		}
		tempDiff := stat.blueCount - stat.redCount
		if tempDiff >= ansDiff {
			ansDiff = tempDiff
			ansSubstr = substr
		}

	}
	return ansSubstr, ansDiff
}
