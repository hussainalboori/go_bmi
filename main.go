package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	// Prompt user for name and age
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter your age:")
	var age int
	fmt.Scan(&age)

	// Prompt user for unit system
	fmt.Println("Choose unit system:")
	fmt.Println("1) Metric (kg/m)")
	fmt.Println("2) US Customary (lb/ft/in)")

	var unitSystem int
	fmt.Scan(&unitSystem)

	switch unitSystem {
	case 1:
		// Prompt user for height and weight in metric units
		fmt.Println("Choose height unit:")
		fmt.Println("1) Meters (m)")
		fmt.Println("2) Centimeters (cm)")

		var heightUnit int
		fmt.Scan(&heightUnit)

		switch heightUnit {
		case 1:
			// Prompt user for height in meters
			fmt.Println("Enter your height (in meters):")
			fmt.Scan(&height)
		case 2:
			// Prompt user for height in centimeters and convert to meters
			fmt.Println("Enter your height (in centimeters):")
			var heightCm float64
			fmt.Scan(&heightCm)
			height = heightCm / 100
		default:
			fmt.Println("Invalid height unit")
			return
		}

		// Prompt user for weight in kilograms
		fmt.Println("Enter your weight (in kilograms):")
		fmt.Scan(&weight)

		// Calculate BMI using metric formula
		bmi := weight / (height * height)

		// Output BMI and weight category to console
		fmt.Printf("%s, age %d, your BMI is: %.2f\n", name, age, bmi)
		if bmi < 18.5 {
			fmt.Println("You are underweight")
		} else if bmi < 25 {
			fmt.Println("You are normal weight")
		} else if bmi < 30 {
			fmt.Println("You are overweight")
		} else {
			fmt.Println("You are obese")
		}

	case 2:
		// Prompt user for height and weight in US Customary units
		fmt.Println("Enter your height (in feet and inches):")
		fmt.Println("Feet:")
		var heightFt float64
		fmt.Scan(&heightFt)
		fmt.Println("Inches:")
		var heightIn float64
		fmt.Scan(&heightIn)

		// Convert height from feet and inches to inches
		heightInch := heightFt*12 + heightIn

		// Prompt user for weight in
		fmt.Println("Enter your weight (in pounds):")
		fmt.Scan(&weight)

		// Calculate BMI using US Customary formula
		bmi := (weight / (heightInch * heightInch)) * 703

		// Output BMI and weight category to console
		fmt.Printf("%s, age %d, your BMI is: %.2f\n", name, age, bmi)
		if bmi < 18.5 {
			fmt.Println("You are underweight")
		} else if bmi < 25 {
			fmt.Println("You are normal weight")
		} else if bmi < 30 {
			fmt.Println("You are overweight")
		} else {
			fmt.Println("You are obese")
		}

	default:
		fmt.Println("Invalid unit system")
	}
}