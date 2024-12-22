package main

import "fmt"

type People struct {
	age    int
	gender string
}

func main() {
	p1 := People{
		age:    18,
		gender: "female",
	}

	p2 := People{
		age:    24,
		gender: "male",
	}

	if p1.age >= 18 {
		fmt.Println("can drive moto")
	}

	if p3 := p2; p3.age >= 24 {
		fmt.Println("can drive car")
	}

	people := []People{p1, p2}

	for _, p := range people {

		fmt.Println(p)
		switch p.gender {
		case "female":
			if p.age >= 18 {
				fmt.Println("Girl Can marrige")
			}
		case "male":
			if p.age >= 24 {
				fmt.Println("Boy Can marrige")
			}

		default:
			fmt.Println("The third world of gender")
		}
	}
}
