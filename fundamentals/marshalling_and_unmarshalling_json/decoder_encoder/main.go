package main

import (
	"encoding/json"
	"log"
	"os"
)

type Student struct {
	FirstName string
	LastName  string
	Age       uint16
}

func main() {
	file, err := os.Open("../assets/student.json")
	defer file.Close()

	file2, err := os.Create("../assets/student2.json")
	defer file2.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)

	decoder.DisallowUnknownFields()

	var student Student

	err = decoder.Decode(&student)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(student)

	encoder := json.NewEncoder(file2)

	err = encoder.Encode(student)

	if err != nil {
		log.Fatal(err)
	}
}
