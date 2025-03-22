package main

import (
	"fmt"
	"math"
)

// Savings
func savings(grossPay int, taxRate float64, expenses int) int {
	afterTaxPay := math.Floor(float64(grossPay) * (1 - taxRate))
	return int(afterTaxPay) - expenses
}

// Material Waste
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	remainingMaterial := totalMaterial - (numJobs * jobConsumption)
	return fmt.Sprintf("%d%s", remainingMaterial, materialUnits)
}

// Interest
func interest(principal int, rate float64, periods int) int {
	finalAmount := math.Floor(float64(principal) + (float64(principal) * rate * float64(periods)))
	return int(finalAmount)
}

func main() {
	fmt.Println(savings(50000, 0.10, 15000))
	fmt.Println(savings(20000, 0.10, 10000))
	fmt.Println(materialWaste(250, "g", 3, 40))
	fmt.Println(materialWaste(500, "kg", 10, 30))
	fmt.Println(materialWaste(200, "m", 10, 40))
	fmt.Println(interest(200000, 0.00, 15)) 
	fmt.Println(interest(100000, 0.05, 10))
}
