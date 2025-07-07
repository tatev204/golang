package artadryal

func MinNum(n int) int {
	for i := 2; ; i = i * 2 {
		if i > n {
			return i
		}
	}
}
