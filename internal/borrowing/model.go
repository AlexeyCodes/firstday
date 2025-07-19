package borrowing

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
)
type Borrowing struct {
	ID       int `json:"id"`
	BookID   int `json:"book_id"`
	ReaderID int `json:"reader_id"`

	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`

	Returned bool `json:"returned"`
}

var Borrowings = []Borrowing{}

func AddBorrowing(b * Borrowing) {
	Borrowings = append(Borrowings, *b)
}

func CreateBorrowing(c *gin.Context) {
	var newBorrowing Borrowing

	if err := c.ShouldBindJSON(&newBorrowing); err != nil {
		c.JSON(400, "Invalid borrowing data")
		return
	}

	newBorrowing.ID = len(Borrowings) + 1
	newBorrowing.Returned = false

	newBorrowing.BorrowDate = time.Now().Format("2006-01-02 15:04:05")
	
	AddBorrowing(&newBorrowing)

	c.JSON(201, newBorrowing)
}

func GetAllBorrowings(c *gin.Context) {
    c.JSON(200, Borrowings)
}

func GetBorrowingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid borrowing's id")
		return
	}

	for _, borrowing := range Borrowings {
		if borrowing.ID == id {
			c.JSON(200, borrowing)
			return
		}
	}
	c.JSON(404, "Borrowing is not found")
}

func UpdateBorrowing(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid borrowing's id")
		return 
	}

	var updatedBorrowing Borrowing

	if err := c.ShouldBindJSON(&updatedBorrowing); err != nil {
		c.JSON(400, "Invalid operation")
		return 
	}

	for i, borrowing := range Borrowings {
		if borrowing.ID == id {
			updatedBorrowing.ID = id
			Borrowings[i] = updatedBorrowing
			c.JSON(200, updatedBorrowing)
			return 
		}
	}
	c.JSON(404, "Borrowing is not found")
}

func DeleteBorrowingByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid borrowing's id")
		return 
	}

	for i, borrowing := range Borrowings {
		if borrowing.ID == id {
			Borrowings = append(Borrowings[:i], Borrowings[i+1:]...)
			c.JSON(200, "Deleted Successfully")
			return 
		}
	}
	c.JSON(404, "Borrowing is not found")
}