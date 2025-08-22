package grading

type StudentResult struct {
	Index        int
	OriginalMark int
	ReexamMark   int
	FinalMark    int
	Grade        string
	Status       string
}

func Grade(m int) string {
	switch {
	case m >= 90:
		return "A"
	case m >= 70:
		return "B"
	case m >= 50:
		return "C"
	default:
		return "F"
	}
}

func Status(orig, final int) string {
	if orig >= 50 {
		return "pass"
	}
	if final >= 50 {
		return "passed"
	}
	return "fail"
}

func ProcessMarks(originals []int, reexam map[int]int) ([]StudentResult, int, int, []int) {
	results := []StudentResult{}
	reexamCollected := []int{}
	passCount, failCount := 0, 0

	for i, orig := range originals {
		rex := reexam[i] // 0 if not in map
		final := orig
		if rex > final {
			final = rex
		}

		grade := Grade(final)
		status := Status(orig, final)

		if orig < 50 {
			reexamCollected = append(reexamCollected, rex)
		}

		if status == "fail" {
			failCount++
		} else {
			passCount++
		}

		results = append(results, StudentResult{
			Index:        i + 1,
			OriginalMark: orig,
			ReexamMark:   rex,
			FinalMark:    final,
			Grade:        grade,
			Status:       status,
		})
	}

	return results, passCount, failCount, reexamCollected
}
