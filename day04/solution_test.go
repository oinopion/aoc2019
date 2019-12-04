package main

import "testing"

func TestHasRepeatingDigit(t *testing.T) {
	if !HasRepeatingDigit(123466) {
		t.Errorf("Should have double")
	}
	if !HasRepeatingDigit(112345) {
		t.Errorf("Should have double")
	}
	if !HasRepeatingDigit(123345) {
		t.Errorf("Should have double")
	}
	if HasRepeatingDigit(123456) {
		t.Errorf("Shouldn't have double")
	}
}

func TestHasIncreasingDigits(t *testing.T) {
	if !HasIncreasingDigits(123456) {
		t.Errorf("Should have increasing digits")
	}
	if !HasIncreasingDigits(111111) {
		t.Errorf("Should have increasing digits")
	}
	if HasIncreasingDigits(111110) {
		t.Errorf("Shouldn't have increasing digits")
	}
	if HasIncreasingDigits(211111) {
		t.Errorf("Shouldn't have increasing digits")
	}
}

func TestHasDouble(t *testing.T) {
	if !HasDouble(123466) {
		t.Errorf("Should have double")
	}
	if !HasDouble(112345) {
		t.Errorf("Should have double")
	}
	if !HasDouble(112233) {
		t.Errorf("Should have double")
	}
	if !HasDouble(111122) {
		t.Errorf("Should have double")
	}
	if !HasDouble(11122) {
		t.Errorf("Should have double")
	}
	if HasDouble(123444) {
		t.Errorf("Shouldn't have double")
	}
	if HasDouble(111111) {
		t.Errorf("Shouldn't have double")
	}
}
