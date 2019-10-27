package internal

import (
	"SMSApp/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//NewBurstSMSProvider returns a burstsms provider
func NewBurstSMSProvider(creds *model.APICredentials) MessagingProvider {
	return &BurstSMSProvider{
		creds: creds,
	}
}

//BurstSMSProvider type represents a message delivery provider for burst sms
type BurstSMSProvider struct {
	creds *model.APICredentials
}

//SendSMS sends sms
func (b *BurstSMSProvider) SendSMS(msg *model.SMSRequest) (*model.APIResponse, error) {
	reqBody, reqBodyErr := b.getRequestBody(msg)
	if reqBodyErr != nil {
		return nil, reqBodyErr
	}

	validateErr := b.validateCreds()
	if validateErr != nil {
		return nil, fmt.Errorf("SendSMS: creds not set")
	}

	newReq, reqErr := http.NewRequest("GET", b.creds.APIURL, bytes.NewBuffer(reqBody))
	if reqErr != nil {
		return nil, reqErr
	}

	newReq.URL.Query().Add("message", msg.Message)
	newReq.URL.Query().Add("format", msg.Format)
	newReq.URL.Query().Add("to", msg.To)

	apiKey := b.creds.APIKey
	apiURL := b.creds.APIURL
	apiSecret := b.creds.APISecret

	newReq.SetBasicAuth(apiKey, apiSecret)
	newReq.Header.Set("Accept", "application/json")
	c := &http.Client{}
	resp, respErr := c.Do(newReq)
	if respErr != nil {
		return nil, respErr
	}

	if resp == nil {
		return nil, fmt.Errorf("SendSMS: response was empty")
	}

	apiResponse, apiErr := b.getResponse(resp)
	if apiErr != nil {
		return nil, apiErr
	}

	if resp.StatusCode == http.StatusOK {
		return apiResponse, nil
	}

	if apiResponse == nil || apiResponse.DelStats == nil {
		return nil, fmt.Errorf("empty api response")
	}

	return nil, fmt.Errorf("api responded with error: %v", apiResponse.DelStats.Error)
}

func (b *BurstSMSProvider) getResponse(resp *http.Response) (*model.APIResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("getResponse: response was received empty")
	}

	var apiResponse model.APIResponse
	apiRespErr := json.NewDecoder(resp.Body).Decode(&apiResponse)
	if apiRespErr != nil {
		return nil, fmt.Errorf("error %v occurred during unmarshall of response", apiRespErr)
	}

	return &apiResponse, nil
}

func (b *BurstSMSProvider) validateCreds() error {
	if b.creds == nil {
		return fmt.Errorf("error creds not created")
	}

	if b.creds.APIKey == "" || b.creds.APISecret == "" || b.creds.APIURL == "" {
		return fmt.Errorf("api credentials not set: api url: %v", b.creds.APIURL)
	}

	return nil
}

func (b *BurstSMSProvider) validateRequest(req *model.SMSRequest) error {
	if req == nil {
		return fmt.Errorf("empty request")
	}

	if strings.Trim(req.Format, " ") == "" {
		req.Format = "json"
	}

	if strings.Trim(req.To, " ") == "" && strings.Trim(req.List, " ") == "" {
		return fmt.Errorf("To field and List field cannot be both empty")
	}

	if strings.Trim(req.Message, " ") == "" {
		return fmt.Errorf("message cannot be empty")
	}

	return nil
}

func (b *BurstSMSProvider) getRequestBody(req *model.SMSRequest) ([]byte, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
