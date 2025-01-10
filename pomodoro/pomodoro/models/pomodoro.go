package models

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Timer struct {
	Tick int
}

type Pomodoro struct {
	ID        int
	Minute    int
	StartTime time.Time
	EndTime   time.Time
}

func Tick(t *Timer) {
	t.Tick++
}

func (p *Pomodoro) Start() {
	p.StartTime = time.Now()
}

func (p *Pomodoro) Stop() {
	p.EndTime = time.Now()
}

func List() {
	dirPath, err := filepath.Abs("assets/sections")

	if err != nil {
		log.Fatal(err)
	}

	dirs, err := os.ReadDir(dirPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		filepath, _ := filepath.Abs("assets/sections/" + dir.Name())
		file, _ := os.Open(filepath)
		dec := gob.NewDecoder(file)

		var po Pomodoro

		dec.Decode(&po)

	}

}
