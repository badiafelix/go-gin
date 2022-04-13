package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", roothandler)

	router.GET("/hello", hellohandler)
	router.GET("/books/:id/:title", booksHandler) //router model index http://localhost:8888/books/11
	router.GET("/query", queryHandler)            //router model query string http://localhost:8888/query?title=bumi manusia
	router.POST("books", postBooksHandler)
	router.Run(":8888")
}

func roothandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "Badia Felix",
		"bio":  "Software Engineer",
	})
}

func hellohandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"content": "Hello world",
		"bio":     "Belajar GOlang",
	})
}
func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`        //binding untuk menggunakan library validasi
	Price int    `json:"price" binding:"required,number"` //binding untuk menggunakan library validasi
	//SubTitle string `json:"sub_title"`                       //jika nama variabel dan nama key tidak sama maka harus pake directive seperti ini
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		//log.Fatal(err)
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		//"sub_title": bookInput.SubTitle,
	})
}
