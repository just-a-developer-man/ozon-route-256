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

	var counter int
	fmt.Fscan(in, &counter)
	for counter != 0 {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
		counter--
	}
}
