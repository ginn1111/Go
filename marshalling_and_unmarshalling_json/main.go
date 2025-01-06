package main

import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"
)

type IdType struct {
	Id string
}

type Class struct {
	Id   string `json:"Mã lớp,omitempty"`
	Name string `json:"Tên lớp,omitempty"`
}

// just demo how to custom UnmarshalJSON
// this function not run well

// func (cl *Class) UnmarshalJSON(data []byte) error {
// cl.<fields> = parse(data)
// }
// func (cl *Class) MarshalJSON() (data []byte, error) {
// return []byte(format(cl.<fields>)), error here
// }

type Student struct {
	Id        string    `json:"Mã SV,omitempty"`
	FirstName string    `json:"Họ,omitempty"`
	LastName  string    `json:"Tên,omitempty"`
	Dob       time.Time `json:"Ngày sinh,omitzero"`
	Class     *Class    `json:"Lớp"`
	Password  string    `json:"-"`
}

func main() {
	myJsonStudent := `{"id": "ec71eac6-c3fe-48cf-ae39-949fde7946fb","firstName": "Pham","lastName": "Thuan"}`
	myJsonStudentList := `
  [
    {
      "Mã SV": "1d7cea58-7a0f-41d6-ba74-1e6ae5a9af76",
      "Họ": "Pham",
      "Tên": "Thuan",
    "Ngày sinh": "2001-09-13T00:00:00.000Z",
      "Lớp": {
        "Mã lớp": "D19CQCNPM01"
      }
    },
    {
      "Mã SV": "1d7cea58-7a0f-41d6-ba74-1e6ae5a9af72",
      "Họ": "Nguyen",
      "Tên": "Hoa",
    "Ngày sinh": "2001-01-24T00:00:00.000Z",
      "Lớp": {
        "Mã lớp": "D19CQHTPM01"
      }
    }
  ]`

	intJson := "8"
	floatJson := "1.23"
	dateJson, _ := time.Now().MarshalJSON()

	var student Student
	var studentList []Student
	var intVar int
	var floatVar float64
	var date time.Time

	json.Unmarshal([]byte(myJsonStudent), &student)
	err := json.Unmarshal([]byte(myJsonStudentList), &studentList)
	json.Unmarshal([]byte(intJson), &intVar)
	json.Unmarshal([]byte(floatJson), &floatVar)
	json.Unmarshal([]byte(dateJson), &date)

	fmt.Println(err)

	fmt.Println(student)
	fmt.Println(studentList)
	fmt.Println(intVar)
	fmt.Println(floatVar)
	fmt.Println(date)

	var unstructedStudentList []map[string]any
	err = json.Unmarshal([]byte(myJsonStudentList), &unstructedStudentList)

	fmt.Println(err)

	for key, val := range unstructedStudentList {
		fmt.Println(key, val)
	}

	if !json.Valid((unsafe.Slice(unsafe.StringData(`{"name": "Thuan"}`), 16))) {
		fmt.Println("Invalid JSON")
	}

	st1 := &Student{
		FirstName: "Pham Van",
		LastName:  "Thuan",
		Password:  "Thuandz",
		Class:     nil,
	}

	st1Json, err := json.Marshal(st1)

	fmt.Println(string(st1Json), err)

	if err != nil {
		fmt.Println(err)
	}

	stList, err := json.MarshalIndent(studentList, "-", " ")

	fmt.Println(string(stList), err)

	unStructureStList, err := json.Marshal(unstructedStudentList)
	fmt.Println(string(unStructureStList), err)
}
