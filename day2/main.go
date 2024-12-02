package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseToSliceOfInt(line []string) []int {
	sliceOfInt := []int{}
	for i := 0; i < len(line); i++ {
		asInt, _ := strconv.Atoi(line[i])
		sliceOfInt = append(sliceOfInt, asInt)
	}
	return sliceOfInt
}

func processLine(line []int, countCon int) int {
	if countCon == 0 {
		return 0
	}

	if line[0] < line[1] {
		for i := 0; i < len(line)-1; i++ {
			if line[i+1]-line[i] < 1 || line[i+1]-line[i] > 3 {
				newSlice := append(line[:i], line[i+1:]...)
				fmt.Println(newSlice, "\n-----")
				return processLine(newSlice, countCon-1)
			}
		}
	} else {
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] < 1 || line[i]-line[i+1] > 3 {
				newSlice := append(line[:i], line[i+1:]...)
				fmt.Println(newSlice, "\n-----")
				return processLine(newSlice, countCon-1)
			}
		}
	}
	return 1
}

func ReadFileLineByLine(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Can not open %v", filename))
	}
	defer file.Close()
	countSafe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, " ")
		countSafe += processLine(parseToSliceOfInt(line), 2)
		fmt.Println(countSafe)
	}
	return countSafe

}

func main() {
	result := ReadFileLineByLine("./input.txt")
	fmt.Println(result)
}
