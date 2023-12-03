package main

import (
	"AoC2023/fileutils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	pt1Ans := Part1("input_1.txt")
	fmt.Println(pt1Ans)
}

func Part1(inputFileName string) int {
	input, err := fileutils.ReadFile(inputFileName)
	if err != nil {
		return -1
	}

	sum := 0
	for _, inputLine := range input {
		sum += ConvertFoundStrToNum(FindAllNums(inputLine))
		if err != nil {
			return -1
		}
	}

	return sum
}

func FindAllNums(inputLine string) []string {
	numRe := regexp.MustCompile(`(\d)`)
	return numRe.FindAllString(inputLine, -1)
}

func ConvertFoundStrToNum(inputStr []string) int {
	firstDigitStr := inputStr[0]
	secondDigitStr := inputStr[len(inputStr)-1]
	foundNum, err := strconv.Atoi(firstDigitStr + secondDigitStr)
	if err != nil {
		return -1
	}
	return foundNum
}
