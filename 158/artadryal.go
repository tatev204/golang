package artadryal

func Multiply() int {

	product := 1

	for i := 100; i <= 999; i++ {
		if i%2 != 0 && i%3 != 0 {
			product *= i
		}
	}
	return product

}
