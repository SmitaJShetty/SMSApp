package model

// APIResponse construct for burst sms response
type APIResponse struct {
	ID             int            `json:"message_id"`
	SentAt         string         `json:"send_at"`
	RecipientCount int            `json:"recipients"`
	Cost           float32        `json:"cost"`
	SMS            int            `json:"sms"`
	DelStats       *DeliveryStats `json:"delivery_stats"`
	Error          *ErrorCode     `json:"error"`
}

// DeliveryStats construct for delivery stats
type DeliveryStats struct {
	Delivered int `json:"delivered"`
	Pending   int `json:"pending"`
	Bounced   int `json:"bounced"`
	Responses int `json:"responses"`
	OptOuts   int `json:"optouts"`
}

// ErrorCode construct for api error respnose
type ErrorCode struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

//ResponseErr construct for api error
type ResponseErr struct {
	Error *ErrorDetails `json:"error"`
}

//ErrorDetails construct for api error
type ErrorDetails struct {
	Code        string            `json:"code"`
	Description *ErrorDescription `json:"description"`
}

// ErrorDescription constrcut for error description
type ErrorDescription struct {
	Fails   []int32  `json:"fails"`
	OptOuts []string `json:"optouts"`
	Reason  string   `json:"reason"`
}
