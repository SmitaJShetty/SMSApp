import {Card, Row, Col, Icon} from 'antd';
import TextArea from 'antd/lib/input/TextArea';
import { Button } from 'antd/lib/radio';
import React, { Props } from "react";
import 'antd/dist/antd.css';
import '../sms/send-message.css';
import { SendSMS } from './send-message-action'
import {Message} from './send-message' ;

interface SendMessageState{
    To: string
    Message: string
    Sent: boolean
}

export default class SendMessage<SendMessageState> extends React.Component{    
    constructor(props:any){
        super(props);
        this.state={
            To: '',
            Message: '',
            Sent:false,
        }
    }

    sendSMS(){
        const msg:Message = {
            to: 
        }
        SendSMS
    }

    render(){
        return (
            
            <Card title="Send SMS" style={{ width: 500, textAlign: "left"}} >
                <Row  gutter={16} >
                    <Col span={6} >To:</Col>
                    <Col span={18} > <input type="text" />{}</Col>
                </Row>
                <Row >
                   <Col span={6} >Message: </Col> 
                   <Col span={6} ><TextArea /></Col>
                </Row>
                <Row>
                    <Button onClick={()=>{ this.sendSMS()}} >Send
                    <Icon type="message"></Icon></Button>     
                </Row>
            </Card>                       
        )
    }
}