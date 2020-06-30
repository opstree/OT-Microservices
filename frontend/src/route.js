import React from "react";
import DashboardIcon from '@material-ui/icons/Dashboard';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import ListIcon from '@material-ui/icons/List';
import Overview from "./overview";
import UserForm from "./form";
import EmployeeList from './ListEmployee';

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
  }
];

export default routes;
