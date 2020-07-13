# Gateway

Gateway is a springboot based API gateway which manages the routing between applications.

## Endpoints

The available endpoints for this application are:-

|**ENDPOINT**|**REQUEST TYPE**|**DESCRIPTION**|
|------------|----------------|---------------|
| `/employee/*` | POST and GET | /employee/* will send all the request for employee related stuff to employee microservice |
| `/attendance/*` | POST and GET | /attendance/* will send all the request for attendance related stuff to attendance microservice  |
| `/salary/*` | POST and GET | /salary/* will send all the request for salary related stuff to salary microservice |
| `/health` | GET | /health will return the health complete gateway service |

## Quickstart

Before compilation [application.yml](./src/resources/application.yml)

```yaml
spring:
  profiles: default

eureka:
  client:
    healthcheck:
      enabled: true

zuul:
  debug:
    request: true

  routes:
    employee:
      path: /employee/**
      url: http://employee:8083/employee
      service-id: /employee/**

    attendance:
      path: /attendance/**
      url: http://attendance:8081/attendance
      service-id: /attendance/**

    salary:
      path: /salary/**
      url: http://salary:8082/salary
      service-id: /salary/**
```

```shell
# For compiling code
make build
```

```shell
java -jar target/gateway-service.jar
```

```shell
make image
```