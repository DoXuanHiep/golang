package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Age  int
}

var person = Person{
	Name: "Hiep",
	Age: 21,
}

var hello = "hello world"

func getPerson(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, person)
}

func getHello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, hello)
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next()
    }
}

func main() {
    router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/person", getPerson)
	router.GET("/hello", getHello)
	router.Run("localhost:9000")
}