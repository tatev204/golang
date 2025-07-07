package artadryal

func Multiply() int {

	product := 1

	for i := 10; i <= 99; i++ {
		if i%15 == 0 {
			product *= i
		}
	}
	return product

}
