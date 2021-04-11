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
	"earth-movers-backend/tools"
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
	i_grade_data := [][]int{}

	if err != nil {
		log.Fatal(err)
	}

	// Convert to ints
	for _, i := range s_grade_data {
		grades := make([]int, len(s_grade_data[0]))
		for _, j := range i {
			j_trimmed := strings.Trim(j, " ")
			grade, err := strconv.Atoi(j_trimmed)
			if err != nil {
				log.Fatal(err)
			}
			grades = append(grades, grade)
		}
		i_grade_data = append(i_grade_data, grades)
	}
	n := tools.Sum(i_grade_data[0])
	k := len(i_grade_data[0])

	distance_matrix := math.DistanceMatrix(i_grade_data, float64(n*(k-1)))


	jsonResult, err := json.Marshal(distance_matrix)

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

	distance_matrix := math.DistanceMatrix(comps, float64(n*(k-1)))

	jsonResult, err := json.Marshal(distance_matrix)

	if err != nil {
		log.Fatal(err)
	}



	c.String(200, string(jsonResult))
	log.Print("Response sent.")
}
