package math

import (
	"earth-movers-backend/tools"
)

type Metric func([]int, []int) float64

func GPA(dist1 []int, dist2 []int) float64 {
	return float64(tools.Weighted_Sum(dist1) - tools.Weighted_Sum(dist2))
}


func EMD(dist1 []int, dist2 []int) float64 {
	length := len(dist1)
	dif := make([]int, length)
	result := 0

	for i := 0; i < length; i++ {
		dif[i] = dist1[i] - dist2[i]
		result += tools.Abs(tools.Sum(dif[:i]))
	}
	return float64(result)
}

func NormedEmd(dist1 []int, dist2 []int) float64 {
	d1_sum := float64(tools.Sum(dist1))
	d2_sum := float64(tools.Sum(dist2))
	dif := make([]float64, len(dist1))

	result := 0.

	for i := 0; i < len(dist1); i++ {
		dif[i] = float64(dist1[i])/d1_sum - float64(dist2[i])/d2_sum

		result += tools.F_Abs(tools.F_Sum(dif[:i]))
	}

	return result/float64(len(dist1)-1)
}

func DistanceMatrix(distributions [][]int, maxVal float64, distance Metric) [][]float64 {
	distance_matrix := make([][]float64, len(distributions))
	for i := range distributions {
		distance_matrix[i] = make([]float64, len(distributions))
	}

	for i := range distributions {
		distance_matrix[i][i] = 0
		for j := i+1; j < len(distributions); j++ {
			dist_val := distance(distributions[i], distributions[j])
			distance_matrix[i][j] = float64(dist_val)/maxVal
			distance_matrix[j][i] = float64(dist_val)/maxVal
		}
	}
	return distance_matrix
}