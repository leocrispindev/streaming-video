package model

type Response struct {
	Uuid    string
	Docs    []map[string]interface{}
	Message string
}
