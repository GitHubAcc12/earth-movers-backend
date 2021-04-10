package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
	"log"

	"earth-movers-backend/math"
	"earth-movers-backend/data"
)



func main() {
	r := gin.Default()
	r.POST("/ping", ping)
	r.Run()
}


func ping(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var request data.Request
	c.BindJSON(&request)

	n, err1 := strconv.Atoi(request.N)
	k, err2 := strconv.Atoi(request.K)

	if err1 =! nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Print("Computing weak compositions of " + strconv.Itoa(n) + " into " + strconv.Itoa(k) + " parts.")

	comps := math.Compositions(n, k)

	jsonResult, err := json.Marshal(comps)

	if err != nil {
		log.Fatal(err)
	}



	c.String(200, string(jsonResult))
}
