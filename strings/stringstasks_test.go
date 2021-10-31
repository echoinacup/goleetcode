package strings

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	expected := "rewq"
	actual := ReverseString([]byte("qwer"))
	if actual != expected {
		t.Errorf("Result was incorrect, got: %s, expected: %s.", actual, expected)
	}
}
