package aoc

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadCommaSeparatedInts(t *testing.T) {
	reader := strings.NewReader("1")
	expected := []int{1}
	testInts(t, reader, expected)

	reader = strings.NewReader("")
	expected = []int{}
	testInts(t, reader, expected)

	reader = strings.NewReader("1,2")
	expected = []int{1, 2}
	testInts(t, reader, expected)

	reader = strings.NewReader("1,2,")
	expected = []int{1, 2}
	testInts(t, reader, expected)

	reader = strings.NewReader("1, 2")
	expected = []int{1, 2}
	testInts(t, reader, expected)
}

func testInts(t *testing.T, reader io.Reader, expected []int) {
	actual, err := ReadCommaSeparatedInts(reader)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v not equal to %v", actual, expected)
	}
}
