package student

func RecursivePower(nb int, power int) int {
	num := nb
	for i := 0; i < power-1; i++ {
		nb = nb * num
	}
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	return nb
}
