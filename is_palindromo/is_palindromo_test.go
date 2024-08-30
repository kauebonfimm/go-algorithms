package is_palindromo_test

import (
	"testing"

	"github.com/kauebonfimm/go-algorithms/is_palindromo"
)

func TestIsPalindromo(t *testing.T) {
	if !(is_palindromo.IsPalindromo(101)) {
		t.Error("invalid test")
	}
	if !(is_palindromo.IsPalindromo(1001)) {
		t.Error("invalid test two sum")
	}
	if is_palindromo.IsPalindromo(3205) {
		t.Error("invalid test two sum")
	}
}
