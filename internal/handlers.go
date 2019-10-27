package internal

import (
	"SMSApp/internal/model"
	"SMSApp/pkg/common"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//PingHandler handler for pong
func PingHandler(w http.ResponseWriter, req *http.Request) {
	common.SendResult(w, req, []byte("Pong"))
}

// SendMessageHandler handles requests for sending sms
func SendMessageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("req:%v \n", req.URL)
	if req.Body == nil {
		log.Printf("request body was empty")
		return
	}

	var msgReq model.SMSRequest
	decodeErr := json.NewDecoder(req.Body).Decode(&msgReq)
	if decodeErr != nil {
		log.Printf("Error while decoding message: %v", decodeErr)
		common.SendErrorResponse(w, req, common.NewAppError(decodeErr.Error(), http.StatusInternalServerError))
		return
	}

	smsGen := NewSMSGenerator()
	resp, sendErr := smsGen.SendSMS(&msgReq)
	if sendErr != nil {
		log.Printf("Error while sending sms message: %v", sendErr)
		common.SendErrorResponse(w, req, common.NewAppError(sendErr.Error(), http.StatusInternalServerError))
		return
	}

	respBody, respErr := json.Marshal(resp)
	if respErr != nil {
		log.Printf("Error while sending sms message: %v", respErr)
		common.SendErrorResponse(w, req, common.NewAppError(respErr.Error(), http.StatusInternalServerError))
		return
	}

	common.SendResult(w, req, respBody)
}
