package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFilename := os.Args[1]
	moduleMasses := readInputMasses(inputFilename)

	naiveTotalFuel := calculateNaiveTotalFuel(moduleMasses)
	fmt.Printf("Part one - total fuel: %d\n", naiveTotalFuel)

	totalFuel := calculateTotalFuel(moduleMasses)
	fmt.Printf("Part two - total fuel: %d\n", totalFuel)
}

func readInputMasses(inputFilename string) []int {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic("Can't open the file")
	}
	scanner := bufio.NewScanner(file)
	var moduleMasses []int
	for scanner.Scan() {
		massText := scanner.Text()
		mass, err := strconv.Atoi(massText)
		if err != nil {
			panic("Problem converting string to int")
		}
		moduleMasses = append(moduleMasses, mass)
	}
	return moduleMasses
}

func calculateNaiveTotalFuel(moduleMasses []int) (totalFuel int) {
	for _, mass := range moduleMasses {
		totalFuel += massToFuel(mass)
	}
	return
}

func calculateTotalFuel(moduleMasses []int) (totalFuel int) {
	for _, mass := range moduleMasses {
		totalFuel += calculateFuelForModule(mass)
	}
	return
}

func calculateFuelForModule(mass int) int {
	additionalFuel := massToFuel(mass)
	totalFuel := additionalFuel
	for additionalFuel > 0 {
		additionalFuel = massToFuel(additionalFuel)
		totalFuel += additionalFuel
	}
	return totalFuel
}

func massToFuel(mass int) int {
	// Note: integer division rounds down, which is wanted
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}
