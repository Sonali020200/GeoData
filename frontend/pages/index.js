import React from 'react';
import ReactDOM from 'react-dom';
import App from './App'; 
import FileUpload from '../components/FileUpload';
import Map from '../components/Map';
import Login from './login';
import Profile from './profile';
import Signup from './signup';
ReactDOM.render(
  <React.StrictMode>
    <FileUpload />
    <Map />
    <Login />
    <Profile />
    <Signup />

  </React.StrictMode>,
  document.getElementById('root')
);
