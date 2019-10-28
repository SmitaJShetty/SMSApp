
export const isPhoneNumber = (inputtxt:string):boolean => 
{
  var phoneno = /^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\s/0-9]*$/g;
  if(inputtxt.match(phoneno)){
      return inputtxt.length===12;
    }
    else
    {
        alert("message");
        return false;
    }
}

export const isValidSMSLength = (inputText:string):boolean => {
    return inputText.length <= 160 
}