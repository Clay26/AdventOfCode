package main

import (
	"AoC2023/utils"
	"regexp"
	"strconv"
)

func main() {
	pt1Ans := Part1("input_1.txt")
	pt2Ans := Part2("input_test_2.txt")
	utils.PrintAnswer(1, 1, pt1Ans)
}

func Part1(inputFileName string) int {
	input, err := utils.ReadFile(inputFileName)
	if err != nil {
		return -1
	}

	sum := 0
	for _, inputLine := range input {
		sum += convertFoundStrToNum(findAllNums(inputLine))
		if err != nil {
			return -1
		}
	}

	return sum
}

func Part2(inputFileName string) int {
	return -1
}

func findAllNums(inputLine string) []string {
	numRe := regexp.MustCompile(`(\d)`)
	return numRe.FindAllString(inputLine, -1)
}

func convertFoundStrToNum(inputStr []string) int {
	firstDigitStr := inputStr[0]
	secondDigitStr := inputStr[len(inputStr)-1]
	foundNum, err := strconv.Atoi(firstDigitStr + secondDigitStr)
	if err != nil {
		return -1
	}
	return foundNum
}
