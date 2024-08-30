package is_palindromo

func IsPalindromo(x int) bool {
	var reverse int
	for copy := x; copy > 0; copy /= 10 {
		reverse = reverse*10 + copy%10
	}

	return x == reverse

}
