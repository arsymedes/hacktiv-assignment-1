package utils

import (
	"errors"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetStudent(code string, participants []participant) (participant, error) {
	for _, el := range participants {
		if el.Code == code {
			return el, nil
		}
	}

	return participant{}, errors.New("student not found")
}
