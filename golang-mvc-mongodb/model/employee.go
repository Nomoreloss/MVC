package model

type Employee struct {
	Id     string  `json:"empid" form:"empid" query:"empid"`
	Name   string  `json:"name" form:"name" query:"name"`
	Email  string  `json:"email" form:"email" query:"email"`
	Salary float64 `json:"salary" form:"salary" query:"salary"`
}
