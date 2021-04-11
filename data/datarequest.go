package data


type DataRequest struct {
	DATA_CSV		string 	`json:"data" binding:"required"`
}