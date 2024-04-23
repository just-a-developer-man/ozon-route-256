package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counters struct {
	xCount int
	yCount int
	zCount int
	noZ    bool
	noY    bool
	noX    bool
}

func (c Counters) isCorrect() string {
	if c.xCount == 0 {
		if c.yCount == 0 {
			if c.zCount == 0 {
				return "Yes"
			}
		}
	}
	return "No"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var datasets int
	fmt.Fscan(in, &datasets)

	for datasets != 0 {
		var (
			queue    string
			queueLen int
		)
		fmt.Fscan(in, &queueLen, &queue)

		counters, err := GenerateCounters(queue)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Fprintln(out, "No")
			datasets--
			continue
		}

		counters = Handle(queue, counters)

		fmt.Fprintln(out, counters.isCorrect())
		datasets--
	}
}

func Handle(queue string, counters Counters) Counters {
	if counters.noZ {
		counters = HandleWithoutZ(queue, counters)
	} else if counters.noX {
		counters = HandleWithoutX(queue, counters)
	} else if counters.noY {
		counters = HandleWithoutY(queue, counters)
	} else {
		counters = HandleGeneral(queue, counters)
	}
	return counters
}

func HandleGeneral(queue string, counters Counters) Counters {
	if queue[0] == 'Z' {
		return counters
	}

	var yDeleteCount, zDeleteCount int

	for index, letter := range queue {
		if letter == 'X' {
			if counters.xCount == 0 || (counters.zCount == 0 && counters.yCount == 0) {
				return counters
			}
			counters.xCount--
			if counters.yCount > counters.zCount || ((counters.yCount == counters.zCount) && IsNextY(queue, index, zDeleteCount, yDeleteCount) && IsNextNotYZ(queue, index, zDeleteCount, yDeleteCount)) {
				counters.yCount--
				yDeleteCount++
			} else {
				counters.zCount--
				zDeleteCount++
			}
		}
		if letter == 'Y' {
			if yDeleteCount > 0 {
				yDeleteCount--
				continue
			}
			if counters.yCount == 0 || counters.zCount == 0 {
				return counters
			}
			counters.yCount--
			counters.zCount--
			zDeleteCount++
		}
		if letter == 'Z' {
			if zDeleteCount > 0 {
				zDeleteCount--
				continue
			}
			return counters
		}
	}

	return counters
}

func HandleWithoutZ(queue string, counters Counters) Counters {
	if queue[0] == 'Y' {
		return counters
	}
	yDeleteCount := 0
	for _, letter := range queue {
		if letter == 'X' {
			if counters.xCount == 0 || counters.yCount == 0 {
				return counters
			}
			counters.xCount--
			counters.yCount--
			yDeleteCount++
		}
		if letter == 'Y' {
			if yDeleteCount > 0 {
				yDeleteCount--
				continue
			}
			return counters
		}
	}
	return counters
}

func HandleWithoutY(queue string, counters Counters) Counters {
	if queue[0] == 'Z' {
		return counters
	}
	zDeleteCount := 0
	for _, letter := range queue {
		if letter == 'X' {
			if counters.xCount == 0 || counters.zCount == 0 {
				return counters
			}
			counters.xCount--
			counters.zCount--
			zDeleteCount++
		}
		if letter == 'Z' {
			if zDeleteCount > 0 {
				zDeleteCount--
				continue
			}
			return counters
		}
	}
	return counters
}

func HandleWithoutX(queue string, counters Counters) Counters {
	if queue[0] == 'Z' {
		return counters
	}
	zDeleteCount := 0
	for _, letter := range queue {
		if letter == 'Y' {
			if counters.yCount == 0 || counters.zCount == 0 {
				return counters
			}
			counters.yCount--
			counters.zCount--
			zDeleteCount++
		}
		if letter == 'Z' {
			if zDeleteCount > 0 {
				zDeleteCount--
				continue
			}
			return counters
		}
	}
	return counters
}

func GenerateCounters(queue string) (Counters, error) {
	counters := Counters{}
	for _, letter := range queue {
		if letter == 'Z' {
			counters.zCount++
		}
		if letter == 'Y' {
			counters.yCount++
		}
		if letter == 'X' {
			counters.xCount++
		}
	}
	if counters.zCount == 0 {
		counters.noZ = true
	}
	if counters.xCount == 0 {
		counters.noX = true
	}
	if counters.yCount == 0 {
		counters.noY = true
	}
	if (counters.noX && counters.noZ) || (counters.noX && counters.noY) || (counters.noY && counters.noZ) {
		return counters, fmt.Errorf("")
	}
	return counters, nil
}

func IsNextY(queue string, currIndex, zDeleteCounter, yDeleteCounter int) bool {
	for _, letter := range queue[currIndex+1:] {
		if letter == 'Y' {
			if yDeleteCounter == 0 {
				return true
			}
			yDeleteCounter--

		}
		if letter == 'Z' {
			if zDeleteCounter == 0 {
				return false
			}
			zDeleteCounter--
		}
	}
	return false
}

func IsNextNotYZ(queue string, currIndex, zDeleteCounter, yDeleteCounter int) bool {
	for index, letter := range queue[currIndex+1:] {
		if letter == 'Y' {
			if yDeleteCounter == 0 && queue[index+1] == 'Z' {
				return true
			}
			yDeleteCounter--
		}
		if letter == 'Z' {
			if zDeleteCounter == 0 {
				return false
			}
			zDeleteCounter--
		}
	}
	return false
}
