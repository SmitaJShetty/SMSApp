import {Card, Row, Col, Icon} from 'antd';
import { Button } from 'antd/lib/radio';
import React from "react";
import 'antd/dist/antd.css';
import '../sms/send-message.css';
import { SendSMS } from './send-message-action'
import {Message, APIResponse, ErrorResponse} from './send-message' ;

interface SendMessageState{
    To: string
    Message: string
    Sent: boolean
    Error: string
}

export default class SendMessage extends React.Component<{}, SendMessageState>{    
    constructor(props:any){
        super(props);
        this.state={
            To: '',
            Message: '',
            Sent: false,
            Error:''
        }
    }

    sendSMS(){
        const msg:Message = {
            to: this.state.To,
            format: "json",
            message: this.state.Message,
        }

        if (!this.validateMessage(msg)){
            this.setState({
                Error: 'Error: please enter valid message and to'
            });
            return ;
        }

        SendSMS(msg)
        .then((response:APIResponse)=>{
            console.log('response received:', response);
            this.setState({Sent: true, Error:''});
        })
        .catch((error:ErrorResponse)=>{
            console.log('error received:', error);
            if (error){
                this.setState({
                    Error: error.Message
                });
            }    else {
                this.setState({
                    Error: 'error occurred',
                });
            }        
        })
    }

    validateMessage = (message:Message):Boolean => {
        if (message.to.trim()===""){
            return false;
        }

        if (message.message.trim() === ""){
            return false
        }
        return true
    }

    updateMessage(event:any){
        this.setState({
            Message: event.target.value,
            Error: '',
        });
    }

    updateTo(event:any){
        this.setState({
            To: event.target.value,
            Error: '',
        });
    }

    render(){
        return (
            
            <Card title="Send SMS" style={{ width: 500, textAlign: "left"}} >
                <Row  gutter={16} >
                    <Col span={6} >To:</Col>
                    <Col span={18} > <input type="text" value={this.state.To} onChange={this.updateTo.bind(this)}/></Col>
                </Row>
                <Row >
                   <Col span={6} >Message: </Col> 
                   <Col span={6} ><input value={this.state.Message} multiple={true}  onChange={this.updateMessage.bind(this)}/></Col>
                </Row>
                <Row>
                    <Button onClick={()=>{ this.sendSMS()}} >Send
                    <Icon type="message"></Icon></Button>     
                </Row>
                <Row>
                    <div>
                        <span id="error" >{this.state.Error}</span>
                    </div>
                </Row>
            </Card>                       
        )
    }
}