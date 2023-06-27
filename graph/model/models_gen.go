// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Department struct {
	ID        int         `json:"ID"`
	Name      string      `json:"Name"`
	Employees []*Employee `json:"Employees"`
}

type Employee struct {
	ID           int         `json:"ID"`
	FirstName    string      `json:"FirstName"`
	LastName     string      `json:"LastName"`
	Username     string      `json:"Username"`
	Password     string      `json:"Password"`
	Email        string      `json:"Email"`
	Dob          string      `json:"DOB"`
	DepartmentID int         `json:"DepartmentID"`
	Position     string      `json:"Position"`
	Department   *Department `json:"Department"`
}