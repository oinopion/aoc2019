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

	memory, err := aoc.ReadCommaSeparatedInts(file)
	if err != nil {
		panic("Can't parse input")
	}

	input := []int{1}
	output := []int{}
	airconProgram := LoadProgram(memory, input, output)
	airconProgram.Execute()
	fmt.Printf("Part one - output: %d\n", airconProgram.output)

	input = []int{5}
	output = []int{}
	radiatorsProgram := LoadProgram(memory, input, output)
	radiatorsProgram.Execute()
	fmt.Printf("Part two - output: %d\n", radiatorsProgram.output)
}

// Opcodes
const (
	AddOp      = 1
	MulOp      = 2
	InOp       = 3
	OutOp      = 4
	JmpTrueOp  = 5
	JmpFalseOp = 6
	LessOp     = 7
	EqOp       = 8
	StopOp     = 99
)

// Opcode ...
type Opcode int

// Parameter modes
const (
	PositionMode  = 0
	ImmediateMode = 1
)

// Mode ...
type Mode int

// Modes ...
type Modes int

// Param returns mode for a parameter at given position
func (modes Modes) Param(position int) Mode {
	digit := digitAt(int(modes), position)
	return Mode(digit)
}

// Program ...
type Program struct {
	memory, input, output []int
	cursor                int
}

// LoadProgram ...
func LoadProgram(memory, input, output []int) *Program {
	return &Program{
		memory: duplicateMemory(memory),
		input:  input,
		output: output,
		cursor: 0,
	}
}

// Execute ...
func (p *Program) Execute() {
	for {
		stop := p.executeOneInstruction()
		if stop {
			return
		}
	}
}

func (p *Program) executeOneInstruction() (stop bool) {
	opcode, paramModes := p.readOpcode()
	switch opcode {
	case AddOp:
		p.opAdd(paramModes)
		return false
	case MulOp:
		p.opMul(paramModes)
		return false
	case InOp:
		p.opIn(paramModes)
		return false
	case OutOp:
		p.opOut(paramModes)
		return false
	case JmpTrueOp:
		p.opJmpTrue(paramModes)
		return false
	case JmpFalseOp:
		p.opJmpFalse(paramModes)
		return false
	case LessOp:
		p.opLess(paramModes)
		return false
	case EqOp:
		p.opEq(paramModes)
		return false
	case StopOp:
		return true
	default:
		panic("Wrong opcode")
	}
}

// Internals

func (p *Program) readOpcode() (opcode Opcode, paramModes Modes) {
	word := p.memory[p.cursor]
	p.cursor++

	return Opcode(word % 100), Modes(word / 100)
}

func (p *Program) readParam(mode Mode) int {
	word := p.memory[p.cursor]
	p.cursor++

	switch mode {
	case ImmediateMode:
		return word
	case PositionMode:
		return p.memory[word]
	default:
		panic("Unknown parameter mode")
	}
}

func (p *Program) storeParam(value int, mode Mode) {
	word := p.memory[p.cursor]
	p.cursor++

	switch mode {
	case PositionMode:
		p.memory[word] = value
	case ImmediateMode:
		panic("Immediate mode not valid for storing values")
	default:
		panic("Unknown parameter mode")
	}
}

// IO

func (p *Program) readFromInput() int {
	value := p.input[0]
	p.input = p.input[1:]
	return value
}

func (p *Program) writeToOutput(value int) {
	p.output = append(p.output, value)
}

// Opcodes

func (p *Program) opAdd(modes Modes) {
	a := p.readParam(modes.Param(0))
	b := p.readParam(modes.Param(1))
	result := a + b
	p.storeParam(result, modes.Param(2))
}

func (p *Program) opMul(modes Modes) {
	a := p.readParam(modes.Param(0))
	b := p.readParam(modes.Param(1))
	result := a * b
	p.storeParam(result, modes.Param(2))
}

func (p *Program) opIn(modes Modes) {
	value := p.readFromInput()
	p.storeParam(value, modes.Param(0))
}

func (p *Program) opOut(modes Modes) {
	value := p.readParam(modes.Param(0))
	p.writeToOutput(value)
}

func (p *Program) opJmpTrue(modes Modes) {
	value := p.readParam(modes.Param(0))
	pointer := p.readParam(modes.Param(1))
	if value != 0 {
		p.cursor = pointer
	}
}

func (p *Program) opJmpFalse(modes Modes) {
	value := p.readParam(modes.Param(0))
	pointer := p.readParam(modes.Param(1))
	if value == 0 {
		p.cursor = pointer
	}
}

func (p *Program) opLess(modes Modes) {
	a := p.readParam(modes.Param(0))
	b := p.readParam(modes.Param(1))
	result := 0
	if a < b {
		result = 1
	}
	p.storeParam(result, modes.Param(2))
}

func (p *Program) opEq(modes Modes) {
	a := p.readParam(modes.Param(0))
	b := p.readParam(modes.Param(1))
	result := 0
	if a == b {
		result = 1
	}
	p.storeParam(result, modes.Param(2))
}

// Helpers

func digitAt(number, position int) int {
	var shifted int
	for shifted = number; position > 0; position-- {
		shifted = shifted / 10
	}
	return shifted % 10
}

func duplicateMemory(memory []int) []int {
	dupe := make([]int, len(memory))
	copy(dupe, memory)
	return dupe
}
