import React from "react";
import Overview from "./overview";
import UserForm from "./form";

const routes = [
  {
    path: "/",
    exact: true,
    name: "Overview",
    toolbar: () => <p>Overview</p>,
    main: () => <Overview />
  },
  {
    path: "/create",
    name: "Add Employee",
    toolbar: () => <p>Add Employee</p>,
    main: () => <UserForm />
  },
  {
    path: "/list",
    name: "List Employee",
    toolbar: () => <p>List Employee</p>,
    main: () => <UserForm />
  }
];

export default routes;
