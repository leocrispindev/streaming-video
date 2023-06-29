package model

type Video struct {
	ID        int    `json:"id,omitempty"`
	Titulo    string `json:"title,omitempty"`
	Sinopse   string `json:"synopsis,omitempty"`
	Categoria int    `json:"category,omitempty"`
	Duracao   int    `json:"duration,omitempty"`
	SemIndice bool   `json:"indexless,omitempty"`
	Extension string `json:"extension,omitempty"`
}
