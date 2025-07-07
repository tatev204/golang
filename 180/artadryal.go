package artadryal

func Bool(N int) bool {
	for i := 3; i <= N; i *= 3 {
		if N == i {
			return true
		}
	}
	return false
}
