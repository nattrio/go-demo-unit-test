package services

func CheckGrade(grade int) string {
	switch {
	case grade >= 80:
		return "A"
	case grade >= 70:
		return "B"
	case grade >= 60:
		return "C"
	case grade >= 50:
		return "D"
	default:
		return "F"
	}
}
