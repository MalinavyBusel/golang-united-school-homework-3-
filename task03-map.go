package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	res := make([]string, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for ind, key := range keys {
		res[ind] = input[key]
	}
	return res
}
