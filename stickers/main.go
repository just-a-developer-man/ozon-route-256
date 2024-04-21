package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var inputString string
	fmt.Fscanf(in, "%s", &inputString)

	var counter int
	fmt.Fscan(in, &counter)

	for counter != 0 {
		var (
			borderLeft, borderRight int
			replace                 string
		)
		fmt.Fscan(in, &borderLeft, &borderRight, &replace)
		inputString = inputString[:borderLeft-1] + replace + inputString[borderRight:]
		counter--
	}
	fmt.Println(inputString)
}
