package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

var person = Person {
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

func getMetrics(c *gin.Context) {
    // Read the JSON file
    fileBytes, err := ioutil.ReadFile("data.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
        return
    }

    // Parse the JSON data
    var data interface{}
    err = json.Unmarshal(fileBytes, &data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
        return
    }

    fmt.Println(data)

    // Send the JSON data as the response
    c.JSON(http.StatusOK, data)
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
    router.GET("/metrics", getMetrics)
    router.Run("localhost:9000")
}