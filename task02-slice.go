package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	var length int = 0

	if len(input) > 0 {
		result = make([]int64, len(input))
		copy(result, input)
		length = len(result)
		for i := 0; i < length/2; i++ {
			tmp := result[i]
			result[i] = result[length-1-i]
			result[length-1-i] = tmp
		}
	}
	return
}
