import React from 'react';

import '../styles/Login.scss';

function Login() {
  return (
  <div className="Login">
    <div className="LoginContainer">
      <h1 className="LoginHeader">Login</h1>
      <form className="LoginForm">
        <input className="LoginInput" type="text" placeholder="Username" />
        <input className="LoginInput" type="password" placeholder="Password" />
        <button className="LoginButton" type="submit">Login</button>
      </form>
      <button className='SignUpButton'>Sign Up</button>
    </div>
  </div>
  );
}

export default Login;