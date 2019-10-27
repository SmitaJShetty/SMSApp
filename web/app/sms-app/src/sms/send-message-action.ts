import { BASEURL } from "../constants";
import { Descriptions } from "antd";

interface Message {
    to: string,
    from: string,
    message: string,
}

interface APIResponse {
    MessageID:string,
    SendAt: string, 
    Recipients: Int16Array,
    Cost: Number,
    SMS: Number,
    DeliveryStats: DeliveryStatsResponse, 
    Error: ErrorResponse
    }

interface DeliveryStatsResponse {
    Delivered: Number,
    Pending: Number,
    Bounced: Number,
    Responses: Number,
    OptOuts: Number,       
}

interface ErrorResponse {
    Code: string,
    Description: string
}

export const SendSMS = (message:Message): Promise<any> => {
    const url:string = BASEURL+"send";
    fetch(url, {
        method: "POST",
        body: JSON.stringify(message),
    })
    .then((response:any)=>{
        if (response.ok){
            return Promise.resolve(response.json());
        }
        return Promise.reject("error");
    })
    .catch((errorResponse:APIResponse)=>{
        console.log(errorResponse);
        return Promise.reject(errorResponse);
    })

    return Promise.resolve();
}