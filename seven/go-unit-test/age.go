package unit

func age(n int) string {
	if n <= 0 {
		return ""
	}
	if n < 7 {
		return "baby"
	}
	if n < 18 {
		return "teenager"
	}
	if n < 60 {
		return "adult"
	}
	return "elderly person"
}
