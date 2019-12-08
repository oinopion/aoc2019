package main

import "testing"

import "reflect"

func TestNewOrbitMap(t *testing.T) {
	orbits := NewOrbitMap("COM)A\nA)B")
	expected := OrbitMap{
		"A": "COM",
		"B": "A",
	}
	if !reflect.DeepEqual(orbits, expected) {
		t.Error("Bad orbit map")
	}
}
func TestOrbitsList(t *testing.T) {
	orbits := OrbitMap{
		"A": "COM",
		"B": "A",
	}
	actual := orbits.OrbitsList("B")
	expected := []string{"A", "COM"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad orbit list")
	}
	actual = orbits.OrbitsList("A")
	expected = []string{"COM"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad orbit list")
	}
	actual = orbits.OrbitsList("COM")
	expected = []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad orbit list")
	}
}

func TestOrbitDepthMap(t *testing.T) {
	orbits := OrbitMap{
		"A": "COM",
		"B": "A",
	}
	if orbits.OrbitDepth("COM") != 0 {
		t.Error("Bad orbit depth")
	}
	if orbits.OrbitDepth("A") != 1 {
		t.Error("Bad orbit depth")
	}
	if orbits.OrbitDepth("B") != 2 {
		t.Error("Bad orbit depth")
	}
}

func TestCommonOrbit(t *testing.T) {
	orbits := OrbitMap{
		"A": "COM",
		"B": "A",
		"C": "B",
		"X": "COM",
		"Y": "X",
	}
	actual := orbits.CommonOrbit("A", "X")
	expected := "COM"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad common orbit")
	}
	actual = orbits.CommonOrbit("C", "B")
	expected = "A"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad common orbit")
	}
	actual = orbits.CommonOrbit("C", "X")
	expected = "COM"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad common orbit")
	}

}
func TestOrbitTransfers(t *testing.T) {
	orbits := OrbitMap{
		"A": "COM",
		"B": "A",
		"C": "B",
		"X": "COM",
		"Y": "X",
	}
	actual := orbits.OrbitTransfers("A", "X")
	expected := 0
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad transfers number")
	}
	actual = orbits.OrbitTransfers("C", "B")
	expected = 1
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad transfers number")
	}
	actual = orbits.OrbitTransfers("C", "Y")
	expected = 3
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Bad transfers number")
	}
}
