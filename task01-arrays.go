package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var (
		length int     = 0
		sum    float32 = 0
	)

	length = len(input)
	if length > 0 {
		for _, i := range input {
			sum += i
		}
	}
	result = float32(sum) / float32(length)
	return
}
