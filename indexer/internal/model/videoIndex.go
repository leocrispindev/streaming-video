package model

import (
	"errors"

	"github.com/leocrispindev/streaming-video/indexer/internal/utils"
)

type Video struct {
	ID         int     `json:"id"`
	Titulo     string  `json:"titulo"`
	Descricao  string  `json:"descricao"`
	Category   int     `json:"category"`
	Duracao    float64 `json:"duration"`
	Indexless  bool    `json:"indexless"`
	Repository string  `json:"repository"`
}

func (v *Video) Validate() []error {
	var errs []error

	if utils.IsEmptyString(v.Repository) {
		errs = append(errs, errors.New("Field [repository] cannot be empty"))
	}

	if utils.IsEmptyString(v.Descricao) {
		errs = append(errs, errors.New("Field [descricao] cannot be empty"))
	}

	if utils.IsEmptyString(v.Titulo) {
		errs = append(errs, errors.New("Field [titulo] cannot be empty"))
	}

	return errs
}
