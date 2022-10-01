import * as React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { ApmRoute } from '@elastic/apm-rum-react'
import { init as initApm } from "@elastic/apm-rum";

import HomePage from "./HomePage.react";
import EmployeeForm from "./EmployeeForm";
import EmployeeList from './EmployeeList';
import AttendanceForm from "./AttendanceForm";
import ListAttendance from './AttendanceList';
import ListSalary from './ListSalary';

import "tabler-react/dist/Tabler.css";

type Props = {||};

initApm({
  serverUrl: "http://apm.monitoring.opstree.com",
  serviceName: "frontend",
});

function App(props: Props): React.Node {
  return (
    <React.Fragment>
      <Router>
        <Switch>
          <ApmRoute exact path="/" component={HomePage} />
          <ApmRoute exact path="/employee-add" component={EmployeeForm} />
          <ApmRoute exact path="/employee-list" component={EmployeeList} />
          <ApmRoute exact path="/attendance-add" component={AttendanceForm} />
          <ApmRoute exact path="/attendance-list" component={ListAttendance} />
          <ApmRoute exact path="/salary" component={ListSalary} />
        </Switch>
      </Router>
    </React.Fragment>
  );
}

export default App;
