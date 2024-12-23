package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("Book added successfully.")
}

// BorrowBook marks a book as borrowed
func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			fmt.Println("Book is already borrowed.")
			return
		}
		book.IsBorrowed = true
		l.Books[id] = book
		fmt.Println("Book borrowed successfully.")
	} else {
		fmt.Println("Book with the given ID does not exist.")
	}
}

// ReturnBook marks a book as returned
func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			fmt.Println("Book is not currently borrowed.")
			return
		}
		book.IsBorrowed = false
		l.Books[id] = book
		fmt.Println("Book returned successfully.")
	} else {
		fmt.Println("Book with the given ID does not exist.")
	}
}

// ListBooks prints all books that are not currently borrowed
func (l *Library) ListBooks() {
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := Library{Books: make(map[string]Book)}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			id := scanner.Text()

			fmt.Print("Enter Book Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Book Author: ")
			scanner.Scan()
			author := scanner.Text()

			library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})

		case "2":
			fmt.Print("Enter Book ID to borrow: ")
			scanner.Scan()
			id := scanner.Text()
			library.BorrowBook(id)

		case "3":
			fmt.Print("Enter Book ID to return: ")
			scanner.Scan()
			id := scanner.Text()
			library.ReturnBook(id)

		case "4":
			fmt.Println("Available books:")
			library.ListBooks()

		case "5":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
