
import React, { useState, useEffect } from "react";
import './LogReg.css';
import LoginForm from './LoginForm';
import RegistrationForm from './RegistrationForm';
// import { Router } from "react-router-dom";

// import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

import { Routes, Route, Link, useNavigate, useLocation } from "react-router-dom";


const LogReg = () => {

  const navigate = useNavigate();
  const location = useLocation();
  const [action, setAction] = useState('');
  

  useEffect(() => {
    if (location.pathname === "/registration") {
      setAction(' active');
    } else {
      setAction('');
    }
  }, [location.pathname]);

  const registreLink = () => {
    navigate("/registration");
  };

  const logLink = () => {
    navigate("/");
  };

  return (
    // <Router>
    <div className={`wrapper${action}`}>
      <Routes>
          <Route path="/"
                  element={<LoginForm onRegisterLinkClick={registreLink} />} />
          <Route path="/registration" 
                  element={<RegistrationForm onLoginLinkClick={logLink} />} />
      </Routes>

      {/* <LoginForm onRegistrLinkClick={registrLink} />
      <RegistrationForm onLoginLinkClick={logLink} /> */}
     </div>
    
 //   </Router>
  );
};

export default LogReg;