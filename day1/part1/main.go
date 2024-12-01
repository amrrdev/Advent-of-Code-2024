package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(left, right []int64) int {
	result := 0
	slices.Sort(left)
	slices.Sort(right)
	for index := 0; index < len(left); index++ {
		result += int(math.Abs(float64(left[index]) - float64(right[index])))
	}
	fmt.Println(result)
	return result
}

func ReadInputFile(fileName string) ([]int64, []int64) {

	file, err := os.Open(fileName)
	if err != nil {
		panic("waaawaaawaaawawawawawa")
	}
	scanner := bufio.NewScanner(file)
	left := []int64{}
	right := []int64{}
	for scanner.Scan() {
		result := strings.Fields(scanner.Text())
		parseLeft, err1 := strconv.ParseInt(result[0], 10, 64)
		parseRight, err2 := strconv.ParseInt(result[1], 10, 64)
		_ = err1
		_ = err2
		left = append(left, parseLeft)
		right = append(right, parseRight)
	}
	return left, right
}

func main() {
	left, right := ReadInputFile("./input.txt")
	Solve(left, right)
}
