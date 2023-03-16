# Salary

Salary is also a java based application which creates and manages employedde's salary information.

## Database

- [Elasticsearch]](../elasticsearch) => Salary application stores information in easticsearch

## Endpoints

The available endpoints for this application are:---

| **ENDPOINT***dd 222               | **REQUEST TYPE** | **DESCRIPTION**                                                                                                            |
|------------------------------|------------------|----------------------------------------------------------------------------------------------------------------------------|
| `/salary/search/all`         | GET              | search will fetches the salary of all users.                                                                               |
| `/salary/healthz`            | GET              | healthz will check the DB connectivity and return the status of application that whether it can serve the requests or not. |

## Environment Variable

| **ENVIRONMENT VARIABLE**            | **DESCRIPTION**                                  |
|-------------------------------------|--------------------------------------------------|
| ELASTIC_APM_SERVICE_NAME            | Service name of elastic APM configuration        |
| ELASTIC_APM_SERVER_URL              | APM service URL for sending metrics and insights |
| SPRING_ELASTICSEARCH_REST_URIS      | URL for elasticsearch database for interaction   |
| SPRING_ELASTICSEARCH_REST_USERNAME  | Username for elasticsearch database              |
| SPRING_ELASTICSEARCH_REST_PASSWORD  | Password for elasticsearch database              |

## GitOps
