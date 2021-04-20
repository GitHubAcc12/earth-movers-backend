package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
	"encoding/csv"
	"log"
	"strings"

	"earth-movers-backend/math"
	"earth-movers-backend/data"
)



func main() {
	r := gin.Default()
	r.POST("/compositions", computeEmd)
	r.POST("/analyzeData", analyzeData)
	r.Run()
}

func analyzeData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var request data.DataRequest
	c.BindJSON(&request)

	r := csv.NewReader(strings.NewReader(request.DATA_CSV))

	s_grade_data, err := r.ReadAll()
	i_grade_data := make([][]int, len(s_grade_data))

	if err != nil {
		log.Fatal(err)
	}


	// Convert to ints
	for r_idx, distr := range s_grade_data {
		grades := make([]int, len(s_grade_data[0]))
		for c_idx, s_grade := range distr {
			s_grade_trimmed := strings.Trim(s_grade, " ")
			grade, err := strconv.Atoi(s_grade_trimmed)
			if err != nil {
				log.Fatal(err)
			}
			grades[c_idx] = grade
		}
		i_grade_data[r_idx] = grades
	}

	// Since now the metric will already normalize the result to [0,1],
	// The value we divide each distance by can be 1
	// Have to use this norming version of EMD, because can't assume
	// Each class to have the same number of students when 
	// analyzing own data
	distance_matrix := math.DistanceMatrix(i_grade_data, 1, math.NormedEmd)

	var response data.Response
	response.EMD_Distances = distance_matrix
	
	gpa_distance_matrix := math.DistanceMatrix(i_grade_data, 4., math.GPA)
	response.GPA_Distances = gpa_distance_matrix


	jsonResult, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}



	c.String(200, string(jsonResult))
	log.Print("Response sent.")

}

func computeEmd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var request data.Request
	c.BindJSON(&request)

	n, err1 := strconv.Atoi(request.N)
	k, err2 := strconv.Atoi(request.K)

	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Print("Computing weak compositions of " + strconv.Itoa(n) + " into " + strconv.Itoa(k) + " parts.")

	comps := math.Compositions(n, k)

	distance_matrix := math.DistanceMatrix(comps, float64(n*(k-1)), math.EMD)

	var response data.Response
	response.EMD_Distances = distance_matrix
	response.GPA_Distances = [][]float64{}

	gpa_distance_matrix := math.DistanceMatrix(comps, 4., math.GPA)
	response.GPA_Distances = gpa_distance_matrix


	jsonResult, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}



	c.String(200, string(jsonResult))
	log.Print("Response sent.")
}


func analyzeBigDataset(distance_matrix [][]float64, threshold float64) {
	// Merge close points: anything with distance less than
	// given threshold gets merged
	
}