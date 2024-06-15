package models

type Object struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}