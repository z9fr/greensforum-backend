package utils

import (
	"context"
	"os"

	"github.com/trycourier/courier-go/v2"
)

func SendEmailWithToken(email, url string) (error, bool) {
	var courier_secret = []byte(os.Getenv("COURIER_TOKEN"))
	client := courier.CreateClient(string(courier_secret), nil)

	requestID, err := client.SendMessage(
		context.Background(),
		courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"to": map[string]string{
					"email": email,
				},
				"template": "V6VQBX78VJ4RN4NF6RPQ5VYZB84A",
				"data": map[string]string{
					"url": url,
				},
			},
		},
	)

	if err != nil {
		LogWarn(err)
		return err, false
	}
	LogInfo(requestID)
	return nil, true
}
