package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	sum := 0
	scoreRecord := []int{}

	for i := range operations {
		switch operations[i] {
		case "+":
			// record a new score that is the sum of the last two previous scores
			score1 := scoreRecord[len(scoreRecord)-1]
			score2 := scoreRecord[len(scoreRecord)-2]
			scoreRecord = append(scoreRecord, score1+score2)
		case "D":
			// record a new score that is double of the previous score
			doublePrevScore := (scoreRecord[len(scoreRecord)-1] * 2)
			scoreRecord = append(scoreRecord, doublePrevScore)

		case "C":
			// delete the previous score from the record
			if len(scoreRecord) != 0 {
				scoreRecord = scoreRecord[:len(scoreRecord)-1]
			}
		default:
			// push elements to scoreRecord
			elem, _ := strconv.Atoi(operations[i])
			scoreRecord = append(scoreRecord, elem)
		}
	}

	for i := range scoreRecord {
		sum += scoreRecord[i]
	}

	return sum
}

func main() {
	ops := []string{"5", "D", "+", "C"}
	fmt.Println(calPoints(ops))
}
