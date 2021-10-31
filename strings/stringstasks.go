package strings

// Reverse string (Rec)
func ReverseString(s []byte) string {
	swap(s, 0, len(s)-1)
	return string(s)

}

func swap(s []byte, begin int, end int) {
	if begin < end {
		tmp := s[begin]
		s[begin] = s[end]
		s[end] = tmp
		swap(s, begin+1, end-1)
	}
	return
}
