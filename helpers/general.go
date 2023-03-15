package helpers

func GetMinMaxArray(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
