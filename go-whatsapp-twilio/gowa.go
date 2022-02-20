package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.Default()
	router.POST("/WA", waMSG)

	fmt.Println("Server started...!")
	router.Run("localhost:8080")
}

func waMSG(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	dataString := string(jsonData)
	if err != nil {
		return
	}

	sendWhatsAppMsg(dataString)
	fmt.Printf("%T\n", dataString)
	fmt.Println(dataString)

}

func sendWhatsAppMsg(d string) {
	from := "whatsapp:+14155238886" // Your Twilio WhatsApp sandbox number
	to := "whatsapp:+XXXXX" // Receiver's WA number
	accountSid := env.TWACID //Twilio account id
	authToken := env.TWAUTH // Twilio Account Auth
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(d)

	resp, err := client.ApiV2010.CreateMessage(params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

}
