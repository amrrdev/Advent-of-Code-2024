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

func SolutionPart1(left, right []int64) int {
	result := 0
	slices.Sort(left)
	slices.Sort(right)
	for index := 0; index < len(left); index++ {
		result += int(math.Abs(float64(left[index]) - float64(right[index])))
	}
	fmt.Println(result)
	return result
}

func CountNumber(rigth []int64, target int64) int64 {
	var result int64 = 0

	for i := 0; i < len(rigth); i++ {
		if rigth[i] == target {
			result++
		}
	}
	return result
}

func SolutionPart2(left, right []int64) int64 {
	var result int64 = 0
	for i := 0; i < len(left); i++ {
		result += CountNumber(right, left[i]) * left[i]
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
	SolutionPart2(left, right)
}
