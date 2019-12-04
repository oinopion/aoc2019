package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFilename := os.Args[1]
	start, end := readInputFile(inputFilename)

	candidates := CountCandidates(start, end)
	fmt.Printf("Part one - candidates: %d\n", candidates)

	restrictedCandidates := CountRestrictedCandidates(start, end)
	fmt.Printf("Part two - restricted candidates: %d\n", restrictedCandidates)
}

func readInputFile(inputFileName string) (int, int) {
	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic("Can't open the input file")
	}
	numbers := strings.Split(string(bytes), "-")
	if len(numbers) != 2 {
		panic(fmt.Sprintf("Input has %d numbers", len(numbers)))
	}
	start, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(fmt.Sprintf("Can't convert %s to int", numbers[0]))
	}
	end, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(fmt.Sprintf("Can't convert %s to int", numbers[1]))
	}

	return start, end
}

// CountCandidates ...
func CountCandidates(start, end int) int {
	count := 0
	for n := start; n <= end; n++ {
		if HasRepeatingDigit(n) && HasIncreasingDigits(n) {
			count++
		}
	}
	return count
}

// CountRestrictedCandidates ...
func CountRestrictedCandidates(start, end int) int {
	count := 0
	for n := start; n <= end; n++ {
		if HasDouble(n) && HasIncreasingDigits(n) {
			count++
		}
	}
	return count
}

// HasRepeatingDigit ...
func HasRepeatingDigit(number int) bool {
	prev := number % 10
	number = number / 10
	for number > 0 {
		curr := number % 10
		if curr == prev {
			return true
		}
		prev = curr
		number = number / 10
	}
	return false
}

// HasDouble ...
func HasDouble(number int) bool {
	doubleCandidate := -1
	consecutive := 1
	prev, number := number%10, number/10

	for number > 0 {
		curr := number % 10
		if curr != prev {
			if doubleCandidate != -1 && consecutive == 2 {
				return true
			}
			doubleCandidate = -1
			consecutive = 1
		} else {
			doubleCandidate = curr
			consecutive++
		}
		prev, number = curr, number/10
	}
	return doubleCandidate != -1 && consecutive == 2
}

// HasIncreasingDigits ...
func HasIncreasingDigits(number int) bool {
	// We're going backwards in digits, so need to check they
	// are decreasing
	prev := number % 10
	number = number / 10
	for number > 0 {
		curr := number % 10
		if curr > prev {
			return false
		}
		prev = curr
		number = number / 10
	}
	return true
}
