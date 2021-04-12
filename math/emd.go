package math

import (
	"log"

	"earth-movers-backend/tools"
)

func emd(dist1 []int, dist2 []int) int {
	length := len(dist1)
	dif := make([]int, length)
	result := 0

	for i := 0; i < length; i++ {
		dif[i] = dist1[i] - dist2[i]
		result += tools.Abs(tools.Sum(dif[:i]))
	}
	return result
}

func DistanceMatrix(distributions [][]int, maxVal float64) [][]float64 {
	distance_matrix := make([][]float64, len(distributions))
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