package homework

func reverse(input []int64) (result []int64) {
	for i := 0; i < len(input)/2; i++{
	x:=input[i]
	input[i] = input[len(input)-i-1]
	input[len(input)-i-1] =x	
	}
	return input
}
