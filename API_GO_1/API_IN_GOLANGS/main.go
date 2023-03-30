package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID string `json:"id"`
	// this is written in Json as we need to serialize our data in Json when we receive our API response
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`

	// Here the first letter of variable names is capital so they can be used outside the struct as well.
}

var books = []book{
	{ID: "1", Title: "In search of lost time", Author: "Rajiv Kapor", Quantity: 2},
	{ID: "2", Title: "In search of time", Author: "Sajiv Kapor", Quantity: 3},
	{ID: "3", Title: "In lost time", Author: "Vednata", Quantity: 4},
}

// Here gin context takes the request and gives the required response.

// NORMAL GET FUNCTION< TO CHECK ALL THE BOOKS.

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// CHECKING OUT BOOK

func checkoutBook(c *gin.Context) {
	// Using Query parameter ?
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messsage": "Missing Query Parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not avialble "})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// REFRENCE TO GET BOOK BY ID

func bookByID(c *gin.Context) {
	id := c.Param("id")
	// The above is Path parameter
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// GETTING BOOK BY ID

func getBookById(id string) (*book, error) {
	//c.IndentedJSON()
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

// CREATING A NEW BOOK ...
// this is done static since this is not link to DB for storing Data.

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// RETURN BOOK FUNCTION -- ADD BOOK IN COLLECTION >

func returnBook(c *gin.Context) {
	// Using Query parameter ?
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messsage": "Missing Query Parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// MAIN FUNCTION .....////

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/returnbook", returnBook)
	router.POST("/books", createBook)
	router.Run("localhost:8080")

}
