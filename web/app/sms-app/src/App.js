import React from 'react';
import './App.css';
import SendMessage from '../src/sms/send-message.tsx';
import {Layout} from 'antd';

const {Header, Content, Footer} = Layout;

function App() {
  return (
    <div className="App">
     <Layout>
        <Header></Header>
        <Content style={{ padding: '0 500px' }}>
            <div>
              <SendMessage />
           </div>
        </Content>
        <Footer>
        </Footer>
      </Layout>
    </div>
   
  );
}

export default App;
