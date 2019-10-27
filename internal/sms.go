package internal

import (
	"SMSApp/internal/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// SMSGenerator construct that sends sms
type SMSGenerator struct {
	MesgProvider MessagingProvider
}

// SendSMS sends sms
func (s *SMSGenerator) SendSMS(req *model.SMSRequest) error {
	validateErr := req.validateRequest()
	if validateErr != nil {
		return validateErr
	}

	sendErr := s.MesgProvider.SendSMS(req)
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

func (s *SMSGenerator) validateRequest(req *model.SMSRequest) error {
	if req == nil {
		return fmt.Errorf("empty request")
	}

	if strings.Trim(req.Format, " ") == "" {
		req.Format = "json"
	}

	if strings.Trim(req.To, " ") == "" && strings.Trim(req.List, " ") {
		return fmt.Errorf("To field and List field cannot be both empty")
	}

	if strings.Trim(req.Message, " ") == "" {
		return fmt.Errorf("message cannot be empty")
	}

	return nil
}
