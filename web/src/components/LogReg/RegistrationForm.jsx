
import React, { useState } from "react";
import { Link } from "react-router-dom";

import { PiUserFill } from "react-icons/pi";
import { FaKey } from "react-icons/fa";
import { IoIosMail } from "react-icons/io";

const RegistrationForm = ({ onLoginLinkClick }) => {

  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [terms, setTerms] = useState('false');
  
    const handleSubmit = async (event) => {
      event.preventDefault();
      
        let form1 = document.forms["registrationForm"];
        let object = {};
        let data = new FormData(form1);
  
        data.forEach(function(value, key){
          if(form1[key].type === 'checkbox'){
            if(form1[key].checked)
            object[key] = true;
          else object[key] = false;
          } else{
            object[key] = value; 
          }
        });

        object['userTerms'] = terms;
 
        var json = JSON.stringify(object);
        console.log(json);
  
        const request = await fetch("http://localhost:8080/registration", {
          // переделай на их
          // https://learn.javascript.ru/fetch
          method: "POST",
          body: json,
          headers: {
            "Content-type": "application/json; charset=UTF-8"
          }
        });
  
    };


  return (
    <div className="form-box registration">
      <form name="registrationForm" id="registrationForm" action=""
      onSubmit={handleSubmit} >
        <h1 id="h1">Registration</h1>

        <div className="inputBox">
          <input name="userLogin" 
                  type="text" 
                  placeholder="  login" 
                  required 
                  value={login} onChange={(e) => setLogin(e.target.value)}  
          />
          <PiUserFill className="icon" />
        </div>

        <div className="inputBox">
          <input name="userEmail" 
                  type="text" 
                  placeholder="  Email" 
                  required 
                  value={email} onChange={(e) => setEmail(e.target.value)}
          />
          <IoIosMail className="icon" />
        </div>

        <div className="inputBox">
          <input name="userPassword" 
                type="password" 
                placeholder="  password" 
                required 
                value={password} onChange={(e) => setPassword(e.target.value)}
          />
          <FaKey className="icon" />
        </div>

        <div className="rememberME">
          <input name="userTerms" 
                  type="checkbox" 
                  id="terms_chbx" 
                  value={terms} onChange={(e) => setTerms(e.target.checked)} />
          <label htmlFor="termsConditions">I agree to the terms & conditions</label><br></br>
        </div>

        <input id="submitButton2" type="submit" value={'Register'} />

        <div className="loginLink">
          Have an account?
          <Link to="/" onClick={onLoginLinkClick}>Login</Link>
          {/* <a href="#" onClick={onLoginLinkClick}>Login</a> */}
        </div>
      </form>
    </div>
  );
};

export default RegistrationForm;