package permutation

func Permutation(result []string, data string, std []string, n int) []string {
	for _, s := range std {
		var d = data
		d += s

		if len(d) == n {
			result = append(result, d)
		} else {
			result = Permutation(result, d, std, n)
		}
	}

	return result
}
