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
	f, err := os.Open("puzzle_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
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

	for _, digit := range line {
		if digitCheck.MatchString(string(digit)) {
			firstDigit = string(digit)
			break
		}
	}

	for _, digit := range reverse(line) {
		if digitCheck.MatchString(string(digit)) {
			lastDigit = string(digit)
			break
		}
	}

	if firstDigit == "" || lastDigit == "" {
		panic("No digits found in line")
	}

	result, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		panic(err)
	}

	return result
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
