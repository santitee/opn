package assignment4

func CalStudentGrade(score int) string {
	switch {
	case score > 100:
		return "can not calculate please check score value"
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
	default:
		return "F"
	}
}
