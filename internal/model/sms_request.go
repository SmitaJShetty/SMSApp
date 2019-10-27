package model

// SMSRequest construct
type SMSRequest struct {
	Format  string `json:"format"`
	Message string `json:"message"`
	To      string `json:"to"`
	List    string `json:"list_id"`
}

// NewSMSRequest creates a new SMSMessage and returns a pointer to the struct
func NewSMSRequest(format string, message string, to string, list *RecipientList) *SMSRequest {
	return &SMSRequest{
		Format:  format,
		Message: message,
		To:      to,
		List:    list.GetCommaSepStr(),
	}
}
