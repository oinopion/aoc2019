package main

import "testing"

import "reflect"

func TestParseInstruction(t *testing.T) {
	actual := ParseInstruction("U8452")
	expected := Instruction{Dir: "U", Dist: 8452}
	if actual != expected {
		t.Errorf("Got wrong instruction: %v", actual)
	}
}

func TestParseInstructionsLine(t *testing.T) {
	actual := ParseInstructionsLine("U1,D2")
	expected := []Instruction{{Up, 1}, {Down, 2}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got wrong instructions list: %v", actual)
	}
}

func TestTranslateInstructionsToPoints(t *testing.T) {
	instructions := ParseInstructionsLine("L1,U1,R1,D1")
	actual := TranslateInstructionsToPoints(instructions)
	expected := []Point{
		{-1, 0, 1},
		{-1, 1, 2},
		{0, 1, 3},
		{0, 0, 4},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got wrong points: %v", actual)
	}
}

func TestIntesectionPoints(t *testing.T) {
	a := []Point{P(1, 0), P(1, 0)}
	b := []Point{P(1, 0), P(0, 1)}
	actual := IntersectionPoints(a, b)
	expected := []Intersection{{P(1, 0), P(1, 0)}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got wrong points: %v", actual)
	}
}

func TestSortPoints(t *testing.T) {
	actual := []Point{P(1, 0), P(0, 0), P(0, -1)}
	SortPoints(actual)
	expected := []Point{P(0, -1), P(0, 0), P(1, 0)}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got wrong points: %v", actual)
	}
}
