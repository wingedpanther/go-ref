// Golang REST API with API versioning and Basic Auth
// TO-DO: Encryption
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "API authentication successful. The id is: " + IdValue
	c.JSON(http.StatusOK, message)

	ReqPayload := make([]byte, 1024)
	ReqPayload, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Request Payload Data: ", string(ReqPayload))
}

func main() {
	router := gin.Default()

	// Create Sub Router for  customised API version
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
		"super": "easypassword",
	}))

	subRouterAuthenticated.GET("/:IdValue", GetMethod)

	listenPort := "1357"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}

// Test: curl -u super:easypassword -i -X GET http://localhost:1357/api/v1/PersonId/Id456
