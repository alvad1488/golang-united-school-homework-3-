package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var keys []int

	if input != nil {
		keys = make([]int, 0, len(input))
		for i := range input {
			keys = append(keys, i)
		}
		sort.Ints(keys)
		result = make([]string, 0, len(input))
		for j := range keys {
			result = append(result, input[j])
		}
	}
	return
}
