import React from "react";
import DashboardIcon from '@material-ui/icons/Dashboard';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import ListIcon from '@material-ui/icons/List';
import Overview from "./overview";
import UserForm from "./form";
import EmployeeList from './ListEmployee';
import CreateAttendance from './CreateAttendance';
import ListAttendance from './ListAttendance';
import SalaryGenerator from './SalaryGenerator';

const routes = [
  {
    path: "/",
    exact: true,
    name: "Overview",
    icon: <DashboardIcon />,
    toolbar: () => <p>Overview</p>,
    main: () => <Overview />
  },
  {
    path: "/create",
    name: "Add Employee",
    icon: <AddCircleOutlineIcon />,
    toolbar: () => <p>Add Employee</p>,
    main: () => <UserForm />
  },
  {
    path: "/list",
    name: "List Employee",
    icon: <ListIcon />,
    toolbar: () => <p>List Employee</p>,
    main: () => <EmployeeList />
  },
  {
    path: "/attendance/create",
    name: "Add Attendance",
    icon: <AddCircleOutlineIcon />,
    toolbar: () => <p>Add Attendance</p>,
    main: () => <CreateAttendance />
  },
  {
    path: "/attendance/list",
    name: "List Attendance",
    icon: <ListIcon />,
    toolbar: () => <p>List Attendance</p>,
    main: () => <ListAttendance />
  },
  {
    path: "/salary",
    name: "Salary",
    icon: <MonetizationOnIcon />,
    toolbar: () => <p>Salary</p>,
    main: () => <SalaryGenerator />
  }
];

export default routes;
