package internal

import (
	"SMSApp/internal/model"
	"fmt"
	"strings"
)

// SMSGenerator construct that sends sms
type SMSGenerator struct {
	MesgProvider MessagingProvider
}

//NewSMSGenerator creates SMSGenerator
func NewSMSGenerator() *SMSGenerator {
	return &SMSGenerator{
		MesgProvider: NewMessagingProvider(),
	}
}

// SendSMS sends sms
func (s *SMSGenerator) SendSMS(msg *model.SMSRequest) (*model.APIResponse, error) {
	reqBodyErr := s.validateRequest(msg)
	if reqBodyErr != nil {
		return nil, reqBodyErr
	}

	resp, sendErr := s.MesgProvider.SendSMS(msg)
	if sendErr != nil {
		return nil, fmt.Errorf("error while sending sms, error:%v", sendErr)
	}

	return resp, nil
}

func (s *SMSGenerator) validateRequest(req *model.SMSRequest) error {
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
