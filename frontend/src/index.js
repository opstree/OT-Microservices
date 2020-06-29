import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
// import Dashboard from './navbar';
// import Overview from './overview';
// import employeeCount from './employeeCount.js';
// import { Switch, Route, BrowserRouter } from 'react-router-dom';
// import CssBaseline from "@material-ui/core/CssBaseline";
// import { Component } from "react";
// import SignupForm from './form.js';


import Main from "./Main";
import ResponsiveDrawers from "./Responsivedrawer";

ReactDOM.render(<ResponsiveDrawers />, document.getElementById("app"));

// ReactDOM.render(<App />, document.querySelector('#app')

// );
ReactDOM.render(<employeeCount />, document.querySelector('#cards'));
