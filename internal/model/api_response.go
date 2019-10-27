package model

// APIResponse construct for burst sms response
type APIResponse struct {
	ID             string         `json:"message_id"`
	SentAt         string         `json:"send_at"`
	RecipientCount int            `json:"recipients"`
	Cost           float32        `json:"cost"`
	SMS            int            `json:"sms"`
	DelStats       *DeliveryStats `json:"delivery_stats"`
}

// DeliveryStats construct for delivery stats
type DeliveryStats struct {
	Delivered bool         `json:"delivered"`
	Pending   bool         `json:"pending"`
	Bounced   bool         `json:"bounced"`
	Responses int          `json:"responses"`
	OptOuts   int          `json:"optouts"`
	Error     *ResponseErr `json:"error"`
}

// ResponseErr construct for api error respnose
type ResponseErr struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
