package smsapp

import (
	"SMSApp/internal/smsapp/model"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// SMSGenerator construct that sends sms
type SMSGenerator struct {
}

// SendSMS sends sms
func (s *SMSGenerator) SendSMS(req *model.SMSRequest) error {
	// call api to send sms
	//get config to generate key
	url := "https://api.transmitsms.com/send-sms.json"
	apiKey := "36b4cb6122e1787a026f331fd3fe66cd"
	apiSecret := "Ubuntu123"

	reqBody, reqBodyErr := s.getRequestBody(req)
	if reqBodyErr != nil {
		return reqBodyErr
	}

	newReq, reqErr := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	if reqErr != nil {
		return reqErr
	}

	newReq.URL.Query().Add("message", req.Message)
	newReq.URL.Query().Add("format", req.Format)
	newReq.URL.Query().Add("to", req.To)

	newReq.Header.Set("Authorization", s.getBase64EncodedKeySecret(apiKey, apiSecret))
	newReq.Header.Set("Accept", "application/json")
	c := &http.Client{}
	resp, respErr := c.Do(newReq)
	if respErr != nil {
		return respErr
	}

	fmt.Println("resp.status:", resp.StatusCode, ", body:")
	return nil
}

func (s *SMSGenerator) getBase64EncodedKeySecret(key string, value string) string {
	toEncodeStr := fmt.Sprintf("%s:%s", key, value)
	encodeStr := base64.StdEncoding.EncodeToString([]byte(toEncodeStr))
	return fmt.Sprintf("Basic %s", encodeStr)
}

func (s *SMSGenerator) getRequestBody(req *model.SMSRequest) ([]byte, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
