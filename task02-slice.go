package homework

func reverse(input []int64) (result []int64) {
	var res = input
	i := 0
	j := len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
