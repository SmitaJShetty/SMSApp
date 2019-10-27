package internal

import (
	"SMSApp/internal/model"
	"os"
)

//MessagingProvider interface for all messaging providers
type MessagingProvider interface {
	SendSMS(msg *model.SMSRequest) (*model.APIResponse, error)
}

//NewMessagingProvider returns new messaging provider
func NewMessagingProvider() MessagingProvider {
	return NewBurstSMSProvider(&model.APICredentials{
		APIKey:    os.Getenv("APIKEY"),
		APISecret: os.Getenv("APISECRET"),
		APIURL:    os.Getenv("APIURL"),
	})
}
