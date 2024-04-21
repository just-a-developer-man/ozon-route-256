package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Solution(in, out)

}

func Solution(in io.Reader, out io.Writer) {
	var userCount, requestCount, notificationNumber int
	fmt.Fscan(in, &userCount, &requestCount)
	var lastAllNotificatedNumber int

	users := make(map[int]int)

	for requestCount != 0 {
		var requestType, id int
		fmt.Fscan(in, &requestType, &id)
		if requestType == 1 {
			notificationNumber += 1
			if id == 0 {
				lastAllNotificatedNumber = notificationNumber
			} else {
				users[id] = notificationNumber
			}
		} else {
			if _, ok := users[id]; ok && users[id] > lastAllNotificatedNumber {
				fmt.Fprintln(out, users[id])
			} else {
				fmt.Fprintln(out, lastAllNotificatedNumber)
			}
		}
		requestCount--
	}
}
