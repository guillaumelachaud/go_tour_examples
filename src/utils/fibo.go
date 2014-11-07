package utils

func Fibonacci() func() float64 {
	sum := float64(0)

	i1 := float64(0)
	i2 := float64(1)

	return func() float64 {
		sum = i1 + i2
		i1 = i2
		i2 = sum
		return sum
	}

}
