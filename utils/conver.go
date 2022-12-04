package utils

import "strconv"

func ConvertToInts(input []string) []int {
	out := make([]int, len(input))

	for i, v := range input {
		con, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out[i] = con
	}

	return out
}
