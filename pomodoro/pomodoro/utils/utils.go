package utils

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go/pomodoro/pomodoro/models"
)

func Save(pomodoro *models.Pomodoro) {
	timestamp := time.Now().Local().UnixMilli()

	newFilePath, err := filepath.Abs(fmt.Sprintf("assets/sections/%v.bin", timestamp))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(newFilePath)

	newFile, err := os.Create(newFilePath)
	defer newFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	enc := gob.NewEncoder(newFile)

	err = enc.Encode(pomodoro)

	if err != nil {
		log.Fatal(err)
	}
}
