package model

// APIResponse construct for burst sms response
type APIResponse struct {
	ID             int            `json:"message_id"`
	SentAt         string         `json:"send_at"`
	RecipientCount int            `json:"recipients"`
	Cost           float32        `json:"cost"`
	SMS            int            `json:"sms"`
	DelStats       *DeliveryStats `json:"delivery_stats"`
	Error          *ResponseErr   `json:"error"`
}

// DeliveryStats construct for delivery stats
type DeliveryStats struct {
	Delivered int `json:"delivered"`
	Pending   int `json:"pending"`
	Bounced   int `json:"bounced"`
	Responses int `json:"responses"`
	OptOuts   int `json:"optouts"`
}

// ResponseErr construct for api error respnose
type ResponseErr struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
