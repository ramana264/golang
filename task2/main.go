package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"task2/grading"
)

func readIntInRange(reader *bufio.Reader, prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		val, err := strconv.Atoi(line)
		if err != nil || val < min || val > max {
			fmt.Printf("Please enter an integer between %d and %d.\n", min, max)
			continue
		}
		return val
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var originals [10]int
	fmt.Println("Enter marks for 10 students (0–100):")
	for i := 0; i < len(originals); i++ {
		originals[i] = readIntInRange(reader, fmt.Sprintf("Student %d: ", i+1), 0, 100)
	}

	reexam := make(map[int]int)
	fmt.Println("\nEnter re-exam marks for students who FAILED initially (<50):")
	for i, m := range originals {
		if m < 50 {
			reexam[i] = readIntInRange(reader,
				fmt.Sprintf("Student %d failed with %d. Re-exam mark: ", i+1, m),
				0, 100)
		}
	}

	results, passCount, failCount, reexamCollected := grading.ProcessMarks(originals[:], reexam)

	fmt.Println("\nFinal Results:")
	fmt.Println("Student\tOriginal\tRe-exam\tFinal\tGrade\tStatus")
	for _, r := range results {
		fmt.Printf("%d\t%d\t\t%d\t%d\t%s\t%s\n",
			r.Index, r.OriginalMark, r.ReexamMark, r.FinalMark, r.Grade, r.Status)
	}

	fmt.Printf("\nTotal Passed: %d\n", passCount)
	fmt.Printf("Total Failed: %d\n", failCount)
	fmt.Printf("Re-exam marks collected: %v\n", reexamCollected)
}
