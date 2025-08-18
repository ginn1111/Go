package main

import "fmt"

type Student struct {
	id        string
	firstname string
	lastname  string
	dob       string
}

type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	// Never panic here, even if ss is nil.
	_, present := ss[key]
	return present
}

type Age int

func (age *Age) IsNil() bool {
	return age == nil
}
func (age *Age) Increase() {
	*age++ // If age is a nil pointer, then
	// dereferencing it will panic.
}

type Book struct {
	pages int
}

type Books []Book

type AliasBooks = Books

func (ab AliasBooks) Pages() int {
	total := 0

	for _, book := range ab {
		total += book.pages
	}

	return total
}

func (books Books) Modify() {
	books = append(books, Book{789})
	books[0].pages = 500
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) Pages2() int {
	return (*b).Pages()
}

func valueEvaluation() {
	var b = Book{pages: 123}
	var p = &b
	var f1 = b.Pages  // value
	var f2 = p.Pages  // reference
	var g1 = p.Pages2 // reference
	var g2 = b.Pages2 // reference

	b.pages = 789

	fmt.Println(f1()) // 123
	fmt.Println(f2()) // 123
	fmt.Println(g1()) // 789
	fmt.Println(g2()) // 789
}

func main() {
	_ = (StringSet(nil)).Has   // will not panic
	_ = ((*Age)(nil)).IsNil    // will not panic
	_ = ((*Age)(nil)).Increase // will not panic

	_ = (StringSet(nil)).Has("key") // will not panic
	_ = ((*Age)(nil)).IsNil()       // will not panic

	// This following line will panic. But the
	// panic is not caused by invoking the method.
	// It is caused by the nil pointer dereference
	// within the method body.
	// ((*Age)(nil)).Increase()

	var books = make(Books, 2, 3)

	copy(books, Books{{123}, {456}})

	fmt.Println(books)

	books.Modify()
	fmt.Println(books)

	var aliasBooks AliasBooks = books

	fmt.Println(aliasBooks.Pages(), books.Pages())

	valueEvaluation()

	Modify()
}
