package interfaces

type Book struct {
	Name   string
	Author string
	Pages  int
}

type IBookHandler interface {
	getFullName() string
}

func (book Book) getFullName(b Book) string {

	return b.Name + " _ " + b.Author
}

func PrintInterface() {

	println("***  begin interfaces ***")

	var book Book

	book.Name = "Clean Code"
	book.Author = "Uncle Bob Martin"
	book.Pages = 345

	newBookList := book.getFullName(book)

	println(newBookList)

	println("***  end interfaces ***")
}
