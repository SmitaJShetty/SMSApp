import { BASEURL } from "../constants";
import {Message, ErrorResponse} from './send-message';

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
    })
    .catch((errorResponse:ErrorResponse)=>{
        console.log("error response:",errorResponse);
        return Promise.reject(errorResponse);
    })

    return Promise.reject();
}