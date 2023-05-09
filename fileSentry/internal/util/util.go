package util

import (
	"errors"
	"log"
	"os"
)

func CreatePath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)

		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil

}
