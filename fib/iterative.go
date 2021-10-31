package fib

// Fib sequence 0 1 1 2 3 5 8
func Iterative(n int) int {
	if n < 2 {
		return n
	}

	var prev, current = 0, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = current + prev
		prev = current
		current = res
	}
	return res
}
