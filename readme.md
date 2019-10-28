Application that sends SMS
==========================

Components: 
----------

a) Backend component 

Source code location: SMSApp/internal
Build location: SMSApp/internal

Execute instructions:
Credentials are passed via environment variables into the application. 
If containerized, these need to be passed into the container as arguments or via a secret provider such as Hashicorp Vault. 

For this app, APIURL, APIKEY, APISECRET are passed as environment variables.
To invoke backend service: on command prompt, cd to SMSApp folder (location of Makefile).
``` 
    export APIKEY=___<APIKEY>___
    export APISECRET=___<APISECRET>___  
    export APIURL=https://api.transmitsms.com/send-sms.json

    make run 
```

This will build and run service. 

b) Frontend component

Source code location: SMSApp/web/app

Execute instructions:
In order to invoke front end, cd to SMSApp folder (location of makefile).
```
    make npm-run
```
This will download dependencies, build and start app

-------------

The application takes a single mobile number and text message. On click of send, it will send an sms. 
Most data validations are provided on the front end and backend. 

Whats not done and could be done with more time: 
1. Tests
2. More logging
3. Containerization if required
