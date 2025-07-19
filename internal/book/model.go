package book

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID   int `json:"id"`
	Year int `json:"year"`

	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`

	Available bool `json:"available"`
}

var Books = []Book{}

func AddBookToLibrary(b *Book) {
	Books = append(Books, *b)
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, "Invalid book data")
		return
	}

	newBook.ID = len(Books) + 1
	newBook.Available = true

	AddBookToLibrary(&newBook)
	c.JSON(201, newBook)
}

func GetAllBooks(c *gin.Context) {
	c.JSON(200, Books)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "Invalid book's id")
		return
	}
	for _, book := range Books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}
	c.JSON(404, "Book is not found")
}

func UpdateBookInLibrary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid book's id")
		return
	}

	var updatedBook Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(400, "Invalid operation")
		return
	}

	for i, book := range Books {
		if book.ID == id {
			updatedBook.ID = id
			Books[i] = updatedBook
			c.JSON(200, updatedBook)
			return
		}
	}
	c.JSON(404, "Book is not found")
}

func DeleteBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid book's id")
		return
	}

	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			c.JSON(200, "Deleted Successfully")
			return
		}
	}
	c.JSON(404, "Book is not found")
}
