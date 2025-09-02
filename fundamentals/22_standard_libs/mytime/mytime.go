package mytime

import (
	"fmt"
	"log"
	"time"
)

func createDate() time.Time {
	currentTime := time.Date(2025, 9, 2, 1, 13, 0, 100, time.UTC)

	fmt.Println(currentTime.Format("02/01 15:04:05 2006 7"))
	fmt.Println(currentTime.Format("03:04:05pm 02/01/2006"))
	fmt.Println(currentTime.Format(time.RFC3339))

	return currentTime
}

func parseTimeAndUTC() {
	dateFrDB := "2025-09-02T22:35:44+07:00"

	america, _ := time.LoadLocation("America/New_York")

	parsedTime, err := time.Parse(time.RFC3339, dateFrDB)

	if err != nil {
		log.Fatal(err)
	}

	utc := parsedTime.UTC()

	fmt.Println(utc.Local(), time.Local)
	fmt.Println(utc.In(america))
}

const FMT_VN = "03:04:05pm 02/01/2006"
const FMT_US = "2006/01/02 15:04:05"

func comparing() {
	firstTime := time.Date(2025, 8, 1, 0, 0, 0, 100, time.Local)
	secondTime := time.Date(2024, 8, 1, 0, 0, 0, 0, time.Local)

	fmt.Println("IS BEFORE", firstTime.Before(secondTime))
	fmt.Println("IS AFTER", firstTime.After(secondTime))
	fmt.Println("DURATION", firstTime.Sub(secondTime))
}

func manipulate() {
	currTime := time.Now()

	const ADD_HOUR = time.Hour
	const SUB_HOUR = -1 * time.Hour

	fmt.Println(currTime.Add(ADD_HOUR))
	fmt.Println(currTime.Add(SUB_HOUR))

}

func MyTime() {

	manipulate()
}

