package main

import "fmt"

type Student struct {
	name  string
	score int
	grade string
}

func CalStudentGrade(score int) string {
	switch {
	case score > 100:
		return "check score value"
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score <= 89:
		return "B"
	case score >= 70 && score <= 79:
		return "C"
	case score >= 60 && score <= 69:
		return "D"
	case score >= 50 && score <= 59:
		return "E"
	case score > 0 && score <= 49:
		return "F"
	}
	return "check score value"
}

func main() {

	fmt.Println("Grading Students scores:")

	// testing ...
	s := []Student{
		Student{name: "A", score: 90},
		Student{name: "B", score: 48},
		Student{name: "C", score: 63},
		Student{name: "D", score: 75},
		Student{name: "E", score: 70},
	}

	for _, score := range s {
		grade := CalStudentGrade(score.score)
		fmt.Printf("Student: %v  has score: %v = grade %s\n", score.name, score.score, grade)
	}
}
