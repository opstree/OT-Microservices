import React from 'react';
import ReactDOM from 'react-dom';
import App from './navbar';
import employeeCount from './employeeCount.js';

ReactDOM.render(<App />, document.querySelector('#app'));
ReactDOM.render(<employeeCount />, document.querySelector('#cards'));
