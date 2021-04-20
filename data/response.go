package data


type Response struct {
	EMD_Distances		[][]float64 `json:"emd_distances" binding:"required"`
	GPA_Distances		[][]float64 `json:"gpa_distances"`
}