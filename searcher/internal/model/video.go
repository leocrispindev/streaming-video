package model

type Video struct {
	ID        int    `json:"id"`
	Titulo    string `json:"title"`
	Sinopse   string `json:"synopsis"`
	Categoria int    `json:"category"`
	Duracao   int    `json:"duration"`
	SemIndice bool   `json:"indexless"`
	Extension string `json:"extension"`
}
