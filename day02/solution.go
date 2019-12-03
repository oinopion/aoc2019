package main

import (
	"fmt"
	"os"

	aoc "hauru.eu/aoc2019"
)

func main() {
	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)
	if err != nil {
		panic("Can't open the input file")
	}

	originalMemory, err := aoc.ReadCommaSeparatedInts(file)
	if err != nil {
		panic("Can't parse input")
	}

	memoryCopy := duplicateMemory(originalMemory)
	fixMemory(memoryCopy, 12, 2)
	executeProgram(memoryCopy)
	fmt.Printf("Part one - position 0: %d\n", memoryCopy[0])

	noun, verb := findInputs(originalMemory, 19690720)
	magicNumber := 100*noun + verb
	fmt.Printf("Part two - noun: %d, verb: %d, magic number: %d\n", noun, verb, magicNumber)
}

func findInputs(memory []int, target int) (noun int, verb int) {
	for noun = 0; noun < 99; noun++ {
		for verb = 0; verb < 99; verb++ {
			memoryCopy := duplicateMemory(memory)
			fixMemory(memoryCopy, noun, verb)
			executeProgram(memoryCopy)
			if memoryCopy[0] == target {
				return noun, verb
			}
		}
	}
	panic("Couldn't find matching pair")
}

func duplicateMemory(memory []int) []int {
	dupe := make([]int, len(memory))
	copy(dupe, memory)
	return dupe
}

func fixMemory(memory []int, noun int, verb int) {
	memory[1] = noun
	memory[2] = verb
}

func executeProgram(memory []int) {
	cursor, op := readOp(0, memory)
	for !op.IsStop() {
		op.Execute(memory)
		cursor, op = readOp(cursor, memory)
	}
}

func readOp(cursor int, memory []int) (newCursor int, op Op) {
	opcode := memory[cursor]
	switch opcode {
	case 1:
		op = Add{
			posA:      memory[cursor+1],
			posB:      memory[cursor+2],
			posResult: memory[cursor+3],
		}
		return cursor + 4, op
	case 2:
		op = Mul{
			posA:      memory[cursor+1],
			posB:      memory[cursor+2],
			posResult: memory[cursor+3],
		}
		return cursor + 4, op
	case 99:
		return cursor + 1, Stop{}
	default:
		panic("Wrong opcode")
	}
}

// Op ...
type Op interface {
	Execute([]int)
	IsStop() bool
}

// Add represents addition
type Add struct {
	posA      int
	posB      int
	posResult int
}

// Execute executes addition
func (op Add) Execute(memory []int) {
	a := memory[op.posA]
	b := memory[op.posB]
	memory[op.posResult] = a + b
}

// IsStop ...
func (op Add) IsStop() bool {
	return false
}

// Mul represents multiplication
type Mul struct {
	posA      int
	posB      int
	posResult int
}

// Execute executes multiplication
func (op Mul) Execute(memory []int) {
	a := memory[op.posA]
	b := memory[op.posB]
	memory[op.posResult] = a * b
}

// IsStop ...
func (op Mul) IsStop() bool {
	return false
}

// Stop represents stop
type Stop struct{}

// Execute for Stop is noop
func (op Stop) Execute(memory []int) {
}

// IsStop ...
func (op Stop) IsStop() bool {
	return true
}
