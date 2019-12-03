package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFilename := os.Args[1]
	lineOne, lineTwo := readInputFile(inputFilename)

	instructionsOne := ParseInstructionsLine(lineOne)
	pointsOne := TranslateInstructionsToPoints(instructionsOne)

	instructionsTwo := ParseInstructionsLine(lineTwo)
	pointsTwo := TranslateInstructionsToPoints(instructionsTwo)

	intersection := IntersectionPoints(pointsOne, pointsTwo)

	dist := FindClosesDistance(intersection)
	fmt.Printf("Part one - closest distance: %d\n", dist)

	delay := FindMinimalDelay(intersection)
	fmt.Printf("Part two - minimal delay: %d\n", delay)
}

func readInputFile(inputFileName string) (string, string) {
	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic("Can't open the input file")
	}
	lines := strings.Split(string(bytes), "\n")
	if len(lines) != 2 {
		panic(fmt.Sprintf("Input has %d lines", len(lines)))
	}
	return lines[0], lines[1]
}

// Direction type
type Direction string

// All directions
const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

// Instruction struct
type Instruction struct {
	Dir  Direction
	Dist int
}

// Point struct
type Point struct {
	X    int
	Y    int
	Step int
}

// Less ...
func (p *Point) Less(o *Point) bool {
	if p.X == o.X {
		return p.Y < o.Y
	}
	return p.X < o.X

}

// Equal ...
func (p *Point) Equal(o *Point) bool {
	return p.X == o.X && p.Y == o.Y
}

// DistToCenter ...
func (p *Point) DistToCenter() int {
	return abs(p.X) + abs(p.Y)
}

// Intersection ...
type Intersection struct {
	A Point
	B Point
}

// Delay ...
func (i *Intersection) Delay() int {
	return i.A.Step + i.B.Step
}

// ParseInstruction ...
func ParseInstruction(s string) Instruction {
	dir := Direction(s[0:1])
	dist, err := strconv.Atoi(s[1:])
	if err != nil {
		panic("Can't create instruction")
	}
	return Instruction{Dir: dir, Dist: dist}
}

// ParseInstructionsLine ...
func ParseInstructionsLine(s string) []Instruction {
	instructions := make([]Instruction, 0)
	instructionStrings := strings.Split(s, ",")
	for _, value := range instructionStrings {
		instructions = append(instructions, ParseInstruction(value))
	}
	return instructions
}

// AdvanceOne ...
func AdvanceOne(p Point, dir Direction) Point {
	switch dir {
	case Up:
		return Point{p.X, p.Y + 1, p.Step + 1}
	case Down:
		return Point{p.X, p.Y - 1, p.Step + 1}
	case Left:
		return Point{p.X - 1, p.Y, p.Step + 1}
	case Right:
		return Point{p.X + 1, p.Y, p.Step + 1}
	default:
		panic("Wrong direction")
	}
}

// AdvanceAll ...
func AdvanceAll(start Point, inst Instruction) (end Point, inBetween []Point) {
	points := make([]Point, 0)
	current := start
	for i := 0; i < inst.Dist; i++ {
		current = AdvanceOne(current, inst.Dir)
		points = append(points, current)
	}
	return current, points
}

// TranslateInstructionsToPoints ..
func TranslateInstructionsToPoints(instructions []Instruction) []Point {
	current := CentralPort()
	points := []Point{}
	inBetween := []Point{}
	for _, inst := range instructions {
		current, inBetween = AdvanceAll(current, inst)
		points = append(points, inBetween...)
	}
	return points
}

// IntersectionPoints ...
func IntersectionPoints(a, b []Point) []Intersection {
	centralPort := CentralPort()
	intersections := []Intersection{}
	SortPoints(a)
	SortPoints(b)
	cursorA := 0
	cursorB := 0
	for cursorA < len(a) && cursorB < len(b) {
		pointA := a[cursorA]
		pointB := b[cursorB]
		switch {
		case pointA.Equal(&pointB):
			if !pointA.Equal(&centralPort) {
				intersections = append(intersections, Intersection{pointA, pointB})
			}
			cursorA++
			cursorB++
		case pointA.Less(&pointB):
			cursorA++
		default:
			cursorB++
		}
	}
	return intersections
}

// FindClosesDistance ...
func FindClosesDistance(intersections []Intersection) int {
	closest := math.MaxInt32
	for _, intersection := range intersections {
		distance := intersection.A.DistToCenter()
		if distance == 0 {
			continue
		}
		if distance < closest {
			closest = distance
		}
	}

	if closest == math.MaxInt32 {
		return -1
	}
	return closest
}

// FindMinimalDelay ...
func FindMinimalDelay(intersections []Intersection) int {
	minimalDelay := math.MaxInt32
	for _, intersection := range intersections {
		delay := intersection.Delay()
		if delay == 0 {
			continue
		}
		if delay < minimalDelay {
			minimalDelay = delay
		}
	}

	if minimalDelay == math.MaxInt32 {
		return -1
	}
	return minimalDelay
}

// CentralPort ...
func CentralPort() Point {
	return Point{0, 0, 0}
}

// P ...
func P(x, y int) Point {
	return Point{x, y, 0}
}

// SortPoints ...
func SortPoints(points []Point) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Less(&points[j])
	})
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
