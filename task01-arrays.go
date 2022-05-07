package homework

func average(input [15]float32) (result float32) {
	var res float32 = 0

	for i := 0; i < len(input); i++ {
		res += input[i]
	}

	return res / float32(len(input))
}
