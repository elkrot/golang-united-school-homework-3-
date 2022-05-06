package homework
import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys = []int{}
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		result = append(result, input[v])
	}
	return result
}
