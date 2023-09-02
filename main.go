package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"arsymedes.github.io/hacktiv-assignment-1/utils"
)

func main() {
	// Read JSON file content
	data, err := os.ReadFile("participants.json")
	utils.CheckError(err)

	// var participants from utils.Participants
	var participants utils.Participants

	// Unmarshal JSON data into participants struct
	utils.CheckError(json.Unmarshal(data, &participants))

	// Check if the argument is given in terminal
	if len(os.Args) < 2 {
		log.Fatal("argument not given")
	}

	// Get a student with the matching student ID
	participant, err := utils.GetStudent(os.Args[1], participants.Participants)
	utils.CheckError(err)

	fmt.Printf("ID: %s\n", participant.Id)
	fmt.Printf("Code: %s\n", participant.Code)
	fmt.Printf("Nama: %s\n", participant.Name)
	fmt.Printf("Alamat: %s\n", participant.Address)
	fmt.Printf("Pekerjaan: %s\n", participant.Occupation)
	fmt.Printf("Alasan: %s\n", participant.Reason)
}
