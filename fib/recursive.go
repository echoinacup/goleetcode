package fib

// Fib sequence 0 1 1 2 3 5 8
func Rec(n int) int {
	if n < 2 {
		return n
	}
	return Rec(n-1) + Rec(n-2)
}

var memo = make(map[int]int)

func Memo(n int) int {
	if _, found := memo[n]; found {
		return memo[n]
	}
	res := 0

	if n < 2 {
		return n
	} else {
		res = Memo(n-1) + Memo(n-2)
	}
	return res
}
