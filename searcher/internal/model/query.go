package model

type Query struct {
	Fields    []string
	Searchers map[string]interface{}
	Page      int
	Size      int
	From      int
}
