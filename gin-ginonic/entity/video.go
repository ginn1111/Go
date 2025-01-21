package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Url         string `json:"url" binding:"required,url"`
	Description string `json:"description" binding:"max=20"`
	Author      Person `json:"author" binding:"required"`
}
