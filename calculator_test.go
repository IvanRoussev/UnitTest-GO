package main

import "testing"

func TestDiscountApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Errorf("Expected 130, Got %v", amount)
	}
}

func TestDiscountMultipliedByTwoApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(200)

	if amount != 160 {
		t.Errorf("Expected 160, Got %v", amount)
	}
}

func TestDiscountMultipliedByThreeApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(350)

	if amount != 290 {
		t.Errorf("Expected 290, Got %v", amount)
	}
}

func TestDiscountNotApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(60)

	if amount != 60 {
		t.Errorf("Expected 50, got %v. Failed because the discount was not expected to be applied", amount)
	}
} 