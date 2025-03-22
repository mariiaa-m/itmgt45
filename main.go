package main

import (
	"fmt"
	"math"
)

// Savings
func savings(grossPay int, taxRate float64, expenses int) int {
	afterTaxPay := math.Floor(float64(grossPay) * (1 - taxRate)) // Apply tax and round down
	return int(afterTaxPay) - expenses                           // Subtract expenses
}

// Material Waste
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	remainingMaterial := totalMaterial - (numJobs * jobConsumption)
	return fmt.Sprintf("%d%s", remainingMaterial, materialUnits) // Convert to string with unit
}

// Interest
func interest(principal int, rate float64, periods int) int {
	finalAmount := math.Floor(float64(principal) + (float64(principal) * rate * float64(periods)))
	return int(finalAmount) // Convert back to integer
}

// Main function for testing
func main() {
	// Test cases
	fmt.Println(savings(100000, 0.12, 20000))     // Expected output: 68000
	fmt.Println(materialWaste(500, "kg", 10, 30)) // Expected output: "200kg"
	fmt.Println(interest(100000, 0.05, 10))       // Expected output: 150000
}

