package User

import "FK_Machine_Coading/Book"

type User struct {
	Id int
	Name string
	Email string
	ContactNumber string
	BookedSlots []*Book.Book
}
