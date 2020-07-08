import * as React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import HomePage from "./HomePage.react";
import EmployeeForm from "./EmployeeForm";
import EmployeeList from './EmployeeList';
import AttendanceForm from "./AttendanceForm";
import ListAttendance from './AttendanceList';
import SalaryGenerator from './SalaryGenerator';

import "tabler-react/dist/Tabler.css";

type Props = {||};

function App(props: Props): React.Node {
  return (
    <React.Fragment>
      <Router>
        <Switch>
          <Route exact path="/" component={HomePage} />
          <Route exact path="/employee-add" component={EmployeeForm} />
          <Route exact path="/employee-list" component={EmployeeList} />
          <Route exact path="/attendance-add" component={AttendanceForm} />
          <Route exact path="/attendance-list" component={ListAttendance} />
          <Route exact path="/salary" component={SalaryGenerator} />
        </Switch>
      </Router>
    </React.Fragment>
  );
}

export default App;
