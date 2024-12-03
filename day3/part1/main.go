package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func stringTransformation(content string) int {
	var result int = 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, value := range matches {
		toInt1, _ := strconv.Atoi(value[1])
		toInt2, _ := strconv.Atoi(value[2])
		result += toInt1 * toInt2
	}
	fmt.Println(doRegex)
	fmt.Println(dontRegex)
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
	// str := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	// res := strings.Split(str, "don't()")
	// fmt.Println(res)
	// fmt.Println(res[0])
	// fmt.Println(res[1])
	// // [xmul(2,4)&mul[3,7]!^ |||  _mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))]
}
