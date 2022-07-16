package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "Boy",
		LastName:  "Business",
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/", indexHandler)
	router.Run(":5000")
}
