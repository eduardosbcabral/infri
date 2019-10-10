package models

type Node struct {
	Id 		int64 	`json: "id"`
	Ip 		string 	`json:"ip" validate:"required"`
	Name 	string	`json: "name" validate:"required"`
}