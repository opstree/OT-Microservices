import React from 'react';
import ReactDOM from 'react-dom';
import ResponsiveDrawers from "./Responsivedrawer";

require('dotenv').config()


ReactDOM.render(<ResponsiveDrawers />, document.getElementById("app"));

// ReactDOM.render(<App />, document.querySelector('#app')

// );
ReactDOM.render(<employeeCount />, document.querySelector('#cards'));
