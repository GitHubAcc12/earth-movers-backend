package tools




func Sum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}

func F_Sum(input []float64) float64 {
	sum := 0.
	for _, i := range input {
		sum += i
	}
	return sum
}

func Weighted_Sum(input []int) int {
	sum := 0
	for i, e := range input {
		sum += i * e
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

func F_Abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

/*
func FindMin(d_list [][]float64) {
	// Expect d_list to be symmetric with 0 on main diagonal
	var min float64
	if len(d_list) > 0 && len(d_list[0]) > 0 {
		min = d_list[0][0]
	} else {
		log.Fatal("Wrong input") // Return error instead?
	}
	for i := 0; i < len(d_list); i++ {
		for j := i+1; j < len(d_list[i]); j++ {
			if min > d_list[i][j] > 0 {
				min = d_list[i][j]
			}
		}
	}
	return min
}*/