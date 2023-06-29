package model

type Response struct {
	Uuid    string  `json:"uuid"`
	Docs    []Video `json:"docs"`
	Page    int     `json:"page"`
	Message string  `json:"message"`
}
