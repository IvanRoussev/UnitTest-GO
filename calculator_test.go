package main

import (
	"testing"
)

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discount int
		purchaseAmount int
		expectedAmountint int	
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmountint: 130},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, expectedAmountint: 160},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, expectedAmountint: 290},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, expectedAmountint: 50},
	}
	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)

	if amount != tc.expectedAmountint {
		t.Errorf("Expected 130, Got %v", amount)
	}
	}
}

// func TestDiscountApplied(t *testing.T) {

// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculate(150)

// 	if amount != 130 {
// 		t.Errorf("Expected 130, Got %v", amount)
// 	}
// }

// func TestDiscountMultipliedByTwoApplied(t *testing.T) {

// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculate(200)

// 	if amount != 160 {
// 		t.Errorf("Expected 160, Got %v", amount)
// 	}
// }

// func TestDiscountMultipliedByThreeApplied(t *testing.T) {

// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculate(350)

// 	if amount != 290 {
// 		t.Errorf("Expected 290, Got %v", amount)
// 	}
// }

// func TestDiscountNotApplied(t *testing.T) {

// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculate(60)

// 	if amount != 60 {
// 		t.Errorf("Expected 50, got %v. Failed because the discount was not expected to be applied", amount)
// 	}
// } 