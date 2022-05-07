package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var tmp = make([]int, len(input))

	var idx int32 = 0
	for key := range input {
		tmp[idx] = key
		idx++
	}

	sort.Ints(tmp)

	result = make([]string, len(tmp))
	for i, k := range tmp {
		result[i] = input[k]
	}

	return result
}
