package artadryal

func PrimeNum(n int) string {

	result := 0
	p := true

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}

	if result == 2 {
		p = true
	} else {
		p = false
	}
	if p {
		return "prime"
	}
	return "not prime"

}
