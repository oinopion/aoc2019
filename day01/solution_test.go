package main

import "testing"

func TestMassToFuel(t *testing.T) {
	mass := 12
	got := massToFuel(mass)
	want := 2
	if got != want {
		t.Errorf("Got wrong amount of fuel: %d instead of %d for %d fuel", got, want, mass)
	}
}

func TestCalculateFuelForModuleSimple(t *testing.T) {
	mass := 12
	got := calculateFuelForModule(mass)
	want := 2
	if got != want {
		t.Errorf("Got wrong amount of fuel: %d instead of %d for %d fuel", got, want, mass)
	}
}
