package internal

import (
	"SMSApp/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	validateErr := b.validateCreds()
	if validateErr != nil {
		return nil, fmt.Errorf("SendSMS: creds not set")
	}

	newReq, reqErr := http.NewRequest("GET", b.creds.APIURL, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	q := newReq.URL.Query()

	q.Add("message", msg.Message)
	q.Add("format", msg.Format)
	q.Add("to", msg.To)
	newReq.URL.RawQuery = q.Encode()

	log.Println("query:", newReq.URL.RawQuery)
	newReq.SetBasicAuth(b.creds.APIKey, b.creds.APISecret)
	newReq.Header.Set("Accept", "application/json")

	c := &http.Client{}
	resp, respErr := c.Do(newReq)
	if respErr != nil {
		return nil, respErr
	}

	if resp == nil {
		return nil, fmt.Errorf(" SendSMS: response was empty")
	}

	// resArr, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	//log.Println(string(resArr))
	apiResponse, apiErr := b.getResponse(resp)
	if apiErr != nil {
		return nil, apiErr
	}

	if resp.StatusCode == http.StatusOK {
		return apiResponse, nil
	}

	log.Fatalf("status not ok: %v", resp.StatusCode)
	if apiResponse == nil || apiResponse.DelStats == nil {
		return nil, fmt.Errorf("empty api response")
	}
	return nil, fmt.Errorf("api responded with error: %v", apiResponse.Error)
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
	log.Println("api response: ", apiResponse)
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
