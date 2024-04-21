package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer out.Flush()
	Solution(in, out)
}

func Solution(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)

	var inputJsonsCount int
	sc.Scan()
	inputJsonsCount, _ = strconv.Atoi(sc.Text())

	for inputJsonsCount != 0 {
		var inputStringsCount int
		sc.Scan()
		inputStringsCount, _ = strconv.Atoi(sc.Text())

		hackCounter := 0
		builder := strings.Builder{}
		unmarshalled := make(map[string]any)

		for inputStringsCount != 0 {
			sc.Scan()
			substring := sc.Text()

			builder.WriteString(substring)
			inputStringsCount--
		}
		err := json.Unmarshal([]byte(builder.String()), &unmarshalled)
		if err != nil {
			panic(err)
		}

		hackCounter = ParseMap(unmarshalled)
		fmt.Fprintln(out, hackCounter)

		inputJsonsCount--
	}
}

func TotalFiles(m map[string]any) (totalFiles int) {
	for key, value := range m {
		switch key {
		case "files":
			totalFiles += len(value.([]any))
		case "folders":
			for _, subDir := range value.([]any) {
				totalFiles += TotalFiles(subDir.(map[string]any))
			}
		}
	}
	return
}

func IsInjectedFilesInDirectory(files []any) bool {
	for _, value := range files {
		if isInjected(value.(string)) {
			return true
		}
	}
	return false
}

func ParseMap(m map[string]any) (injected int) {
	if files, ok := m["files"]; ok {
		if IsInjectedFilesInDirectory(files.([]any)) {
			injected += len(files.([]any))
		}
	}
	if injected > 0 {
		if folders, ok := m["folders"]; ok {
			for _, subDir := range folders.([]any) {
				injected += TotalFiles(subDir.(map[string]any))
			}
		}
	} else {
		if folders, ok := m["folders"]; ok {
			for _, subDir := range folders.([]any) {
				injected += ParseMap(subDir.(map[string]any))
			}
		}
	}
	return injected
}

func isInjected(s string) bool {
	if strings.HasSuffix(s, ".hack") {
		return true
	}
	return false
}
