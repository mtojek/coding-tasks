package coding_tasks

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n - 1)
}