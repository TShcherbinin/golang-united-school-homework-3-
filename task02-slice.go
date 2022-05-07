package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))

	var idx int32 = 0
	for i := len(input) - 1; i >= 0; i-- {
		result[idx] = input[i]
		idx++
	}

	return result
}
