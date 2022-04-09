package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//* Membuat variabel
	router := gin.Default()

	//* Membuat route URL
	//* Membuat function dengan anonymous function
	// router.GET("/", func(c *gin.Context) {
	// 	// Untuk mengembalikan JSON
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Name" : "Jery Hardianto",
	// 		"Bio" : "Software Engineer",
	// 	})
	// })


	//* Membuat function dengan anonymous function
	// router.GET("/hello", func(c *gin.Context) {
	// 	// Untuk mengembalikan JSON
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"content" : "Hello World",
	// 		"Bio" : "Software Engineer",
	// 	})
	// })

	//* Membuat function dengan named function
	router.GET("/", rootHandler)
	router.GET("/hello", halloHandler)



	//Menjalankan di localhost:9999 / 127.0.0.1:9999
	router.Run(":9999")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"Name" : "Jery Hardianto",
			"Bio" : "Software Engineer",
		})
}
func halloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"content" : "Hello World",
			"Bio" : "Software Engineer",
		})
}
	
