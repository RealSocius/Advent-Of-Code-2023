package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// open file
	f, err := os.Open("puzzle_2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var digitCheck = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[1-9]`)
	var total int

	for scanner.Scan() {
		total += getLineValue(scanner.Text(), digitCheck)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func getLineValue(line string, digitCheck *regexp.Regexp) int {

	var firstDigit string
	var lastDigit string

	digits := digitCheck.FindAllString(line, -1)

	firstDigit = getDigitValue(digits[0])
	lastDigit = getDigitValue(digits[len(digits)-1])

	result, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		panic(err)
	}

	return result
}

func getDigitValue(digit string) string {
	switch digit {
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
		return digit
	}
}
