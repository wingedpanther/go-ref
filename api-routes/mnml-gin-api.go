package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DATA
type contact struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

var contacts = []contact{
	{ID: "1", Name: "Superman", Email: "superman@heros.meta", Mobile: "00234-2323"},
	{ID: "2", Name: "Batman", Email: "baty@heros.meta", Mobile: "00234-5666"},
	{ID: "3", Name: "Spider-man", Email: "spidey@heros.meta", Mobile: "00234-6788"},
}


func main() {
	router := gin.Default()
  //routes

	//[GET]
	router.GET("/contacts", getContacts)
	//[POST]
	router.POST("/contact", addContact)
	//[POST/id]
	router.GET("/contact/:id", getContact)

	fmt.Println("Server started...!")
	router.Run("localhost:8080")
}

func getContacts(c *gin.Context) { 	//[GET]
	c.IndentedJSON(http.StatusOK, contacts)
}

func addContact(c *gin.Context) { //[POST]
	var newcontact contact
		
	if err := c.BindJSON(&newcontact); err != nil {
		return
	}
	contacts = append(contacts, newcontact)
	c.IndentedJSON(http.StatusCreated, newcontact)
}

func getContact(c *gin.Context) { 	//[POST/id]
	id := c.Param("id")
	for _, a := range contacts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found"})
}
