package utils

func Max(arr []int) int {
	max := arr[0]

	for _, i := range arr {
		if i > max {
			max = i
		}
	}

	return max
}
