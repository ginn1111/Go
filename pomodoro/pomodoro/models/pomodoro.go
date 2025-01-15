package models

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Timer struct {
	Tick  int
	CStop chan bool
}

type Pomodoro struct {
	ID        int
	Minute    int
	StartTime time.Time
	EndTime   time.Time
}

func (t *Timer) Stops() {
	t.CStop <- true
}

func (t *Timer) Ticks() {
	go func() {
		select {
		case <-t.CStop:
			return

		case <-time.After(time.Second):
			isClose := <-t.CStop
			if isClose {
				close(t.CStop)
				return
			}
			t.Tick++
			log.Println(t.Tick)
			t.Ticks()
		}
	}()
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

	for i, dir := range dirs {
		filepath, _ := filepath.Abs("assets/sections/" + dir.Name())
		file, _ := os.Open(filepath)
		dec := gob.NewDecoder(file)

		var po Pomodoro

		dec.Decode(&po)

		log.Printf("Section %d: %s - %s -> total: %dm", i+1, po.StartTime.Format("02/01/2006 15:04"), po.EndTime.Format("02/01/2006 15:04"), int(po.EndTime.Sub(po.StartTime).Seconds()/60))
	}
}
