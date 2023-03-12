package model
 
type Employee struct {
    Id   string	`json:"id"`
    Name string	`json:"name"`
    Experience int `json:"experience"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employee
}

type ErrorResponse struct {
	Code    int
	Message string
}