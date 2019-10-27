package internal

import (
	"SMSApp/internal/model"
	"SMSApp/pkg/errorcode"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRouter returns a router
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/send", SendMessageHandler).Methods("POST").Schemes("http").Host("127.0.0.1:8080")
	return router
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
		http.Error(w, decodeErr.Error(), int(errorcode.Internal))
	}

	fmt.Println("mesgReq:", msgReq)
	generator := NewMessagingProvider()
	sendErr := smsGen.SendSMS(&msgReq)
	if sendErr != nil {
		log.Printf("Error while sending sms message: %v", sendErr)
		http.Error(w, sendErr.Error(), int(errorcode.Internal))
	}
}
