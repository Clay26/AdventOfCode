package dayone

import (
	"AoC2023/utils"
	"regexp"
	"strconv"
)

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
	input, err := utils.ReadFile(inputFileName)
	if err != nil {
		return -1
	}

	sum := 0
	for _, inputLine := range input {
		wordsReplaced := true
		for wordsReplaced {
			inputLine, wordsReplaced = replaceWords(inputLine)
		}
		sum += convertFoundStrToNum(findAllNums(inputLine))
		if err != nil {
			return -1
		}
	}

	return sum
}

func replaceWords(inputLine string) (string, bool) {
	numRe := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine)`)
	if !numRe.MatchString(inputLine) {
		return inputLine, false
	}

	return numRe.ReplaceAllStringFunc(inputLine, func(match string) string {
		switch match {
		case "zero":
			return "0ero"
		case "one":
			return "1ne"
		case "two":
			return "2wo"
		case "three":
			return "3hree"
		case "four":
			return "4our"
		case "five":
			return "5ive"
		case "six":
			return "6ix"
		case "seven":
			return "7even"
		case "eight":
			return "8ight"
		case "nine":
			return "9ine"
		default:
			return match
		}
	}), true
}

func findAllNums(inputLine string) []string {
	numRe := regexp.MustCompile(`(\d)`)
	return numRe.FindAllString(inputLine, -1)
}

func findAllNumsAndWords(inputLine string) []string {
	numRe := regexp.MustCompile(`(\d|zero|one|two|three|four|five|six|seven|eight|nine)`)
	return numRe.FindAllString(inputLine, -1)
}

func convertFoundStrToNum(inputStr []string) int {
	firstDigitStr := convertFoundStrToDigit(inputStr[0])
	secondDigitStr := convertFoundStrToDigit(inputStr[len(inputStr)-1])
	foundNum, err := strconv.Atoi(firstDigitStr + secondDigitStr)
	if err != nil {
		return -1
	}
	return foundNum
}

func convertFoundStrToDigit(inputStr string) string {
	switch inputStr {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return inputStr
	}
}
