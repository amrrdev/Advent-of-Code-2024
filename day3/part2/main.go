package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func stringTransformation(content string) int {
	var result int = 0

	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	isScanning := false

	for i := 0; i < len(content); i++ {
		if i+4 <= len(content) && doRegex.MatchString(content[i:i+4]) {
			isScanning = true
			i += 3
			continue
		}
		if i+7 <= len(content) && dontRegex.MatchString(content[i:i+7]) {
			isScanning = false
			i += 6
			continue
		}

		if isScanning && i+5 < len(content) && content[i:i+4] == "mul(" {
			match := mulRegex.FindStringSubmatch(content[i:])
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				fmt.Printf("Found numbers: %d, %d\n", num1, num2)
				result += num1 * num2
			}
			i += len(match[0]) - 1
		}
	}
	return result
}

func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("fuck")
	}
	return string(content)
}

func main() {
	content := readFile("./../input.txt")
	result := stringTransformation(content)
	fmt.Println(result)
}
