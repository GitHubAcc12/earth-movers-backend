package math

import (

	"earth-movers-backend/tools"
)


func emd(dist1 []int, dist2 []int) int {
	// Only support same length distributions for now
	/*
	length := 0
	if len(dist1) > len(dist2) { 
		length = len(dist1)
	} else {
		length = len(dist2)
	}
	*/

	length := len(dist1)
	dif := make([]int, length)
	result := 0

	i := 0
	for i < length {
		dif[i] = dist1[i] - dist2[i]
		j := 0
		for j < i {
			result += tools.Abs(tools.Sum(dif))
			j++
		}
		i++
	}
	return result
}

func DistanceMatrix(distributions [][]int, maxVal float64) [][]float64 {
	distance_matrix := make([][]float64, len(distributions))

	// Have to make one run through it here for the sym. init.
	// of the lower triangular matrix
	// Should prob be optimized
	for i := range distributions {
		distance_matrix[i] = make([]float64, len(distributions))
	}

	for i := range distributions {
		distance_matrix[i][i] = 0
		for j := i+1; j < len(distributions); j++ {
			dist_val := emd(distributions[i], distributions[j])
			distance_matrix[i][j] = float64(dist_val)/maxVal
			distance_matrix[j][i] = float64(dist_val)/maxVal
		}
	}
	return distance_matrix
}