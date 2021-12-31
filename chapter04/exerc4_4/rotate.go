package rotate

func Rotate(slice []int, times int) {
	size := len(slice)
	buffer := make([]int, size)
	for i, j := 0, times%size; i < size; i, j = i+1, j+1 {
		buffer[i] = slice[j%size]
	}
	for i := range slice {
		slice[i] = buffer[i]
	}
}
