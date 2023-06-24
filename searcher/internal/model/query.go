package model

type Query struct {
	Fields    []string
	Searchers map[string]interface{}
}
