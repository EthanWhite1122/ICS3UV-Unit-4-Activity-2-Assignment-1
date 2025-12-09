// Author: Ethan White
// version: 1.0.0
// date: 2025-12-08
// Fileoverview: This program will take a start and end input then add up all numbe that the user inputs in the range and out of the range

package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Print("Enter the starting integer: ")
	fmt.Scan(&start)

	fmt.Print("Enter the ending integer: ")
	fmt.Scan(&end)

	sumInRange := 0
	sumOutRange := 0

	var num int

	for {
		fmt.Print("Enter an integer (0 to stop): ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if num >= start && num <= end {
			sumInRange += num
		} else {
			sumOutRange += num
		}
	}

	fmt.Println("Sum of integers in range:", sumInRange)
	fmt.Println("Sum of integers out of range:", sumOutRange)
	fmt.Println("\nDone.")
}
