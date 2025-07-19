package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"firstday/internal/book"
	"firstday/internal/borrowing"
	"firstday/internal/reader"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")
	
	// Serve HTML files from the "templates" directory
	r.StaticFS("/public", http.Dir("./public"))
	
	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Serve main HTML page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Library Management System",
		})
	})

	// Alternative route to serve JSON API info (useful for API testing)
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome to Library Management API",
			"version":  "v1.1.0",
			"endpoints": gin.H{
				"books":      "/api/v1/books",
				"readers":    "/api/v1/readers",
				"borrowings": "/api/v1/borrowings",
				"health":     "/api/v1/health",
				"stats":      "/api/v1/stats",
			},
		})
	})

	// API group for versioning
	api := r.Group("/api/v1")

	// ============ BOOK ROUTES ============
	bookRoutes := api.Group("/books")
	{
		bookRoutes.POST("/", book.CreateBook)        // Create book
		bookRoutes.GET("/", book.GetAllBooks)        // Get all books (supports ?available=true/false)
		bookRoutes.GET("/:id", book.GetBookByID)     // Get book by ID
		bookRoutes.PUT("/:id", book.UpdateBookInLibrary) // Update book
		bookRoutes.DELETE("/:id", book.DeleteBookByID)   // Delete book
	}

	// ============ READER ROUTES ============
	readerRoutes := api.Group("/readers")
	{
		readerRoutes.POST("/", reader.CreateReader)      // Create reader
		readerRoutes.GET("/", reader.GetAllReaders)      // Get all readers
		readerRoutes.GET("/:id", reader.GetReaderByID)   // Get reader by ID
		readerRoutes.PUT("/:id", reader.UpdateReader)    // Update reader
		readerRoutes.DELETE("/:id", reader.DeleteReaderByID) // Delete reader
	}

	// ============ BORROWING ROUTES ============
	borrowingRoutes := api.Group("/borrowings")
	{
		borrowingRoutes.POST("/", borrowing.CreateBorrowing)        // Create borrowing
		borrowingRoutes.GET("/", borrowing.GetAllBorrowings)        // Get all borrowings (supports ?returned=true/false)
		borrowingRoutes.GET("/:id", borrowing.GetBorrowingByID)     // Get borrowing by ID
		borrowingRoutes.PUT("/:id", borrowing.UpdateBorrowing)      // Update borrowing
		borrowingRoutes.DELETE("/:id", borrowing.DeleteBorrowingByID) // Delete borrowing
	}

	// ============ ADDITIONAL USEFUL ROUTES ============
	// API health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Library Management API is running",
			"version": "v1.1.0",
		})
	})

	// Statistics endpoint
	api.GET("/stats", func(c *gin.Context) {
		totalBooks := len(book.Books)
		availableBooks := 0
		for _, b := range book.Books {
			if b.Available {
				availableBooks++
			}
		}

		totalReaders := len(reader.Readers)
		activeBorrowings := 0
		for _, b := range borrowing.Borrowings {
			if !b.Returned {
				activeBorrowings++
			}
		}

		c.JSON(200, gin.H{
			"total_books":      totalBooks,
			"available_books":  availableBooks,
			"borrowed_books":   totalBooks - availableBooks,
			"total_readers":    totalReaders,
			"active_borrowings": activeBorrowings,
			"total_borrowings": len(borrowing.Borrowings),
		})
	})

	// Start server on port 8080
	r.Run(":8080")
}