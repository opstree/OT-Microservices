# OT Go-App

The main goal of creating this sample Golang application is to provide an environment and idea of Golang Application's **build**, **test** and **deploy** phase.

## Requirments

- **[Golang](https://golang.org/)** ---> For development environment
- **[Docker](https://www.docker.com/)** ---> For dockerizing the application
- **[Dep](https://github.com/golang/dep)** ---> For Golang dependency management

## Overview

OT Go-App is a CRUD application which provides a Web UI Interface for Employee Management. As a functionality wise it provides:- 

- Web UI interface for employee management
- It stores all the data in MySQL database
- It provides functionality of auto-reconnection of database
- Generates log file for access log and error in */var/log*
    - Access logs will lies in `/var/log/ot-go-webapp.access.log`
    - Error logs will lies in `/var/log/ot-go-webapp.error.log`
- We can pass the database credentials via properties file or environment variables
- For properties file we have to store `database.properties` at this location `/etc/conf.d/ot-go-webapp/database.properties` and the content should be something like this :-

```properties
DB_USER = root
DB_PASSWORD = password
DB_URL = 172.17.0.3
DB_PORT = 3306
```

- For environment variables we have to set these environment variables
    - **DB_USER** ---> Name of the database user
    - **DB_PASSWORD** ---> Password of the database user
    - **DB_URL** ---> URL of the database server
    - **DB_PORT** ---> Port on which database is running

**Important:- In MySQL database should exist with name employeedb**

## To Do
- [X] Implement logging
- [X] Property file 
- [X] Add more fields for employee
- [X] Write unit tests
- [X] Fix code if there is any mess
- [X] Integrate dependency management
- [ ] Fill README with more information
- [X] Make application more attractive
- [X] Add healthcheck API
- [X] Logging of acccess and error log
- [ ] Provide file uploading functionality
