package reader

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Reader struct {
	ID int `json:"id"`

	Name             string `json:"name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	RegistrationDate string `json:"registration_date"`
}

var Readers = []Reader{}

func AddReader(b *Reader) {
	Readers = append(Readers, *b)
}

func CreateReader(c *gin.Context) {
    var newReader Reader
    if err := c.ShouldBindJSON(&newReader); err != nil {
        c.JSON(400, "Invalid reader data")
        return
    }
    newReader.ID = len(Readers) + 1  
    AddReader(&newReader)
    c.JSON(201, newReader)
}
func GetAllReaders(c *gin.Context) {
	c.JSON(200, Readers)
}

func UpdateReader(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid reader's id")
		return
	}

	var updatedReader Reader

	if err := c.ShouldBindJSON(&updatedReader); err != nil {
		c.JSON(400, "Invalid operation")
		return
	}

	for i, reader := range Readers {
		if reader.ID == id {
			updatedReader.ID = id
			Readers[i] = updatedReader
			c.JSON(200, updatedReader)
			return
		}
	}
	c.JSON(404, "Reader is not found")
}


func GetReaderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid reader's id")
		return
	}

	for _, reader := range Readers {
		if reader.ID == id {
			c.JSON(200, reader)
			return
		}
	}
	c.JSON(404, "Reader is not found")
}

func DeleteReaderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Invalid reader's id")
		return
	}

	for i, reader := range Readers {
		if reader.ID == id {
			Readers = append(Readers[:i], Readers[i+1:]...)
			c.JSON(200, "Deleted Successfully")
			return
		}
	}
	c.JSON(404, "Reader is not found")
}
