package main

import (
	"bufio"
	"fmt"
	"os"
)

func grade(m int) string { // grade distribution
	switch {
	case m >= 90:
		return "A"
	case m >= 70:
		return "B"
	case m >= 50:
		return "C"
	default:
		return "Fail"
	}
}

func main() {
	const n = 10
	var marks [n]int

	in := bufio.NewReader(os.Stdin)

	// Read first-attempt marks
	fmt.Println("Enter marks for 10 students (0-100):")
	for i := 0; i < n; i++ {
		for {
			fmt.Printf("Student %d: ", i+1)
			_, err := fmt.Fscan(in, &marks[i])
			if err != nil {
				fmt.Println("Invalid input, try again.")
				in.ReadString('\n') // clear bad token
				continue
			}
			if marks[i] < 0 || marks[i] > 100 {
				fmt.Println("Please enter a value between 0 and 100.")
				continue
			}
			break
		}
	}

	// Collect re-exam marks for failed students into a slice
	var failedIdx []int
	for i := 0; i < n; i++ {
		if marks[i] < 50 {
			failedIdx = append(failedIdx, i)
		}
	}

	reExamMarks := make([]int, len(failedIdx)) // slice for re-exam marks
	if len(failedIdx) > 0 {
		fmt.Printf("\nRe-exam for %d failed students:\n", len(failedIdx))
		for j, idx := range failedIdx {
			for {
				fmt.Printf("Re-exam mark for Student %d (first %d): ", idx+1, marks[idx])
				_, err := fmt.Fscan(in, &reExamMarks[j])
				if err != nil {
					fmt.Println("Invalid input, try again.")
					in.ReadString('\n')
					continue
				}
				if reExamMarks[j] < 0 || reExamMarks[j] > 100 {
					fmt.Println("Please enter a value between 0 and 100.")
					continue
				}
				break
			}
		}
	} else {
		fmt.Println("\nNo re-exams needed; everyone passed on first attempt.")
	}

	// Map student index -> re-exam mark
	reMap := make(map[int]int)
	for j, idx := range failedIdx {
		reMap[idx] = reExamMarks[j]
	}

	// Compute final marks, grade, and status
	finalMarks := make([]int, n)
	status := make([]string, n) // "pass", "passed", or "fail"
	totalPass, totalFail := 0, 0

	for i := 0; i < n; i++ {
		first := marks[i]
		re := reMap[i]
		if first >= 50 {
			finalMarks[i] = first
			status[i] = "pass"
			totalPass++
		} else {
			// take maximum of first and re-exam
			if re > first {
				finalMarks[i] = re
			} else {
				finalMarks[i] = first
			}
			if finalMarks[i] >= 50 {
				status[i] = "passed" // passed after re-exam
				totalPass++
			} else {
				status[i] = "fail"
				totalFail++
			}
		}
	}

	// Print table
	fmt.Printf("\n%-8s %-12s %-8s %-8s %-7s %-8s\n", "Student", "FirstAttempt", "Re-exam", "Final", "Grade", "Status")
	for i := 0; i < n; i++ {
		re := "-"
		if v, ok := reMap[i]; ok {
			re = fmt.Sprintf("%d", v)
		}
		fmt.Printf("%-8d %-12d %-8s %-8d %-7s %-8s\n",
			i+1, marks[i], re, finalMarks[i], grade(finalMarks[i]), status[i])
	}

	fmt.Printf("\nTotal passed: %d\nTotal failed: %d\n", totalPass, totalFail)
}
