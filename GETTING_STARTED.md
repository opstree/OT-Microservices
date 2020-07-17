## Overview

There are different langauge stack are used to build this microservices. If you are new to the microservice world, we would recommend you to run the application locally first.

For local development environment, we need these tools installed to build/compile the applications:-

- [Golang](https://golang.org/dl/)
- [Python 3](https://www.python.org/download/releases/3.0/)
- [Java 8](https://www.java.com/en/download/faq/java8.xml)
- [Maven](https://maven.apache.org/download.cgi)
- [NodeJs](https://nodejs.org/en/download/)
- [Yarn](https://yarnpkg.com/lang/en/docs/install/)

## Database

As we discussed, in the main [README](.README.md) page that we are using these two database:-

- [Elasticsearch](https://www.elastic.co/downloads/elasticsearch)
- [MySQL](https://www.mysql.com/downloads/)

We need to install these databases well.

## Application's Sequence and Flow

Once the databases are installed you can compile the application and make it ready to serve requests. The application build, configuration, and URL related details can be found after clicking on below applications.

### Attendance

But let's move step by step, first we can try to up attendance service. So the attendance service is dependent on some services so we have to make sure that those dependent services are up.

The dependent service are:-

- [MySQL](./mysql)
- [Gateway](./gateway)

Once the dependent service are up, we can up the attendance service:-

- [Attendance](./attendance)

For validating the attendance service, we have to make sure that our frontend and webserver is also up

- [Frontend](./frontend)
- [Webserver](./webserver)

## Employee

So once the attendance service is up, the next service we can hop on is employee service. So employee service depends on elasticsearch service.

- [Elasticsearch](./elasticsearch)

Once the elastisearch is up, we can up the employee service as well

- [Employee](./employee)

So the frontend is already deployed and we can test the employee management functionality from there.

## Salary

Salary service depends on elasticsearch which is already up, so we can directly up that service without any hesitation.

- [Salary](./salary)

Once the salary service is up, you can try generating the PDF from salary tab.
