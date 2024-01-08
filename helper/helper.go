package helper

import (
	"errors"
	"fetch/model"
)

func AlbumById(id string) (*model.Album, error) {
	for i, a := range model.Albums {
		if a.ID == id {
			return &model.Albums[i], nil
		}
	}
	return nil, errors.New("Album not found")
}

func ValidateData(id string) (*model.Album, error) {
	album, err := AlbumById(id)

	if err != nil {
		return nil, errors.New("Album not found")
	}
	return album, nil
}

func Add_numbers(num1 int, num2 int) (total int) {
	return num1 + num2
}
