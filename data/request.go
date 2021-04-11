package data


type Request struct {
	N		string 	`json:"students" binding:"required"`
	K		string 	`json:"grades" binding:"required"`
}