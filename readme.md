** Application that sends SMS **

Components: 
a) Backend component 

Source code location: SMSApp/internal
Build location:SMSApp/internal

Execute instructions:
To invoke backend service: on command prompt, cd to SMSApp folder (location of Makefile).
``` 
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
Tests
More logging
Containerization if required