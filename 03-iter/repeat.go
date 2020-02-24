package iteration

func Repeat(char string, n_reps int) (result string) {
	for i := 1; i <= n_reps; i++ {
		result += char
	}

	return
}
