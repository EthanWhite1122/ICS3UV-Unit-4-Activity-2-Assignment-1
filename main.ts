/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-12-08
 * @fileoverview This program will take a start and end input then add up all numbe that the user inputs in the range and out of the range
 */

//Input
const startInput = prompt("Enter the starting integer:");
const endInput = prompt("Enter the ending integer:");

//create variables
const start = Number(startInput);
const end = Number(endInput);

let sumInRange = 0;
let sumOutRange = 0;

let num: number;

//begin loop 
do {
  const input = prompt("Enter an integer (0 to stop):");
  num = Number(input);

  if (num !== 0) {
    if (num >= start && num <= end) {
      sumInRange += num;
    } else {
      sumOutRange += num;
    }
  }
} while (num !== 0);

//print
console.log("Sum of integers in range:", sumInRange);
console.log("Sum of integers out of range:", sumOutRange);

console.log("\nDone.")