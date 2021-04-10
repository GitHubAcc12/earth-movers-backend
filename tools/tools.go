package tools

func Sum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}