package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputFilename := os.Args[1]
	input := readInputFile(inputFilename)

	orbitMap := NewOrbitMap(input)
	// fmt.Println(orbitMap)

	totalOrbits := calcTotalOrbits(orbitMap)
	fmt.Printf("Part one - total orbits: %d\n", totalOrbits)

	totalTransfers := orbitMap.OrbitTransfers("YOU", "SAN")
	fmt.Printf("Part two - total orbits: %d\n", totalTransfers)

}

func readInputFile(inputFileName string) string {
	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic("Can't open the input file")
	}
	return string(bytes)
}

// UCM is Universal Center of Mass, it doesn't orbit anything
const UCM = "COM"

// OrbitMap maps an orbiting body to it's orbit center
type OrbitMap map[string]string

// NewOrbitMap ...
func NewOrbitMap(input string) OrbitMap {
	mapping := OrbitMap{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ")")
		center, orbiter := parts[0], parts[1]
		mapping[orbiter] = center
	}
	return mapping
}

// OrbitsList returns list of all direct and indirect centers
func (om OrbitMap) OrbitsList(body string) []string {
	if body == UCM {
		return []string{}
	}

	orbits := []string{}
	for current := om[body]; current != UCM; current = om[current] {
		orbits = append(orbits, current)
	}
	return append(orbits, UCM)
}

// OrbitDepth calculates depth of the orbit from the center
func (om OrbitMap) OrbitDepth(body string) int {
	return len(om.OrbitsList(body))
}

// CommonOrbit returns deepest common orbit center
func (om OrbitMap) CommonOrbit(a, b string) string {
	orbitsA := om.OrbitsList(a)
	orbitsB := om.OrbitsList(b)

	pA, pB := len(orbitsA)-1, len(orbitsB)-1
	for pA >= 0 && pB >= 0 && orbitsA[pA] == orbitsB[pB] {
		pA--
		pB--
	}
	return orbitsA[pA+1]
}

// OrbitTransfers ...
func (om OrbitMap) OrbitTransfers(a, b string) int {
	transfers := 0
	commonOrbit := om.CommonOrbit(a, b)

	orbitsA := om.OrbitsList(a)
	for i := 0; orbitsA[i] != commonOrbit; i++ {
		transfers++
	}

	orbitsB := om.OrbitsList(b)
	for i := 0; orbitsB[i] != commonOrbit; i++ {
		transfers++
	}
	return transfers
}

func calcTotalOrbits(om OrbitMap) int {
	total := 0
	for key := range om {
		total = total + om.OrbitDepth(key)
	}
	return total
}
