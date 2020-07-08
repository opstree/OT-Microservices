# Gateway

Gateway is a springboot based API gateway which manages the routing between applications.

## Quickstart

Before compilation [application.yml](./src/resources/application.yml)

```yaml
# Application port on which application will listen
spring:
  profiles: default

zuul:
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