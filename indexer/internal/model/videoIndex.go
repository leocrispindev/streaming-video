package model

import (
	"errors"

	"github.com/leocrispindev/streaming-video/indexer/internal/utils"
)

type Document struct {
	Action     int        `json:"action"`
	ID         int        `json:"id"`
	Key        string     `json:"key"`
	Repository string     `json:"repository"`
	VideoInfo  *VideoInfo `json:"videoInfo,omitempty"`
}

type VideoInfo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Sinopse   string `json:"sinopse"`
	Category  int    `json:"category"`
	Duration  int    `json:"duration"`
	Indexless bool   `json:"indexless"`
}

func (v *VideoInfo) Validate() []error {
	var errs []error

	if utils.IsEmptyString(v.Sinopse) {
		errs = append(errs, errors.New("Field [SINOPSE] cannot be empty"))
	}

	if utils.IsEmptyString(v.Title) {
		errs = append(errs, errors.New("Field [TITLE] cannot be empty"))
	}

	return errs
}

func (doc *Document) Validate() []error {
	var errs []error

	if utils.IsEmptyString(doc.Repository) {
		errs = append(errs, errors.New("Field [REPOSITORY] cannot be empty"))
	}

	if utils.IsEmptyString(doc.Key) {
		errs = append(errs, errors.New("Field [KEY] cannot be empty"))
	}

	return errs
}
