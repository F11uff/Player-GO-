// LoginForm.jsx
import { Link } from "react-router-dom";

import React, { useState } from "react";
import { FaKey } from "react-icons/fa";
import { PiUserFill } from "react-icons/pi";


const LoginForm = ({ onRegisterLinkClick }) => {
  
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [rememberME, setRememberMe] = useState('false');

  const handleSubmit = async (event) => {
    event.preventDefault();

      let form = document.forms["loginForm"];
      let object = {};
      let data = new FormData(form);

      data.forEach(function(value, key){
        if(form[key].type === 'checkbox'){
          if(form[key].checked)
          object[key] = true;
        else object[key] = false;
        } else{
          object[key] = value; 
        }
      });

      object['userRemember'] = rememberME;
      //чтоб значение чекубокса отобрадалось даже, если на него не жмали

      var json = JSON.stringify(object);
      console.log(json);

      const request = await fetch("http://localhost:8080/", {
        // переделай на другой
        // https://learn.javascript.ru/fetch
        method: "POST",
        body: json,
        headers: {
          "Content-type": "application/json; charset=UTF-8"
        }
      })
  };

    return (
    <div className="form-box login">
      <form name="loginForm" className="loginForm" action=""
        onSubmit={handleSubmit} 
      >
        <h1 id="h1">Login</h1>

        <div className="inputBox">
          <input name="userLogin" 
                  type="text"
                  placeholder=" login" 
                  required 
                  autoComplete="username"
                  value={login} onChange={(e) => setLogin(e.target.value)}
          />
          <PiUserFill className="icon" />
        </div>

        <div className="inputBox">
          <input name="userPassword" 
                  type="password" 
                  placeholder=" password" 
                  autoComplete="current-password" 
                  value={password} onChange={(e) => setPassword(e.target.value)}
          />
          <FaKey className="icon" />
        </div>

        <div className="rememberME">
          <input name="userRemember" 
                type="checkbox" 
                value={rememberME} onChange={(e) => setRememberMe(e.target.checked)}

                />
          <label htmlFor="remember">Remember me</label><br></br>
          {/* <a href="#">Forgot password?</a> */}
          <button className="forgot_pass">Forgot password? </button>
        </div>

        <input name="submitButton" id="submitButton1" type="submit" value={'Login'} />
        {/* <p>{login}</p>
        <p>{password}</p> */}

        <div className="registrationLink">
          Don't have an account?
        <Link to="/registration" onClick={onRegisterLinkClick}>Register</Link>
        </div> 
      </form>
    </div>
  );
};

export default LoginForm;