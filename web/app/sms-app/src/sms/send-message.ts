
export interface Message {
    to: string,
    from: string,
    message: string,
}

export interface APIResponse {
    MessageID:string,
    SendAt: string, 
    Recipients: Int16Array,
    Cost: Number,
    SMS: Number,
    DeliveryStats: DeliveryStatsResponse, 
    Error: ErrorResponse
    }

export interface DeliveryStatsResponse {
    Delivered: Number,
    Pending: Number,
    Bounced: Number,
    Responses: Number,
    OptOuts: Number,       
}

export interface ErrorResponse {
    Code: string,
    Description: string
}