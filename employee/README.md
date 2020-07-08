# Employee

Employee microservice is also designed in Golang to manage employee's information.

## Database

- [Elasticsearch](../elasticsearch) => Employee application stores information in easticsearch

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DESCRIPTION**|
|------------------------|---------------|
| CONFIG_FILE | Path of configuration file |

## Quickstart

```yaml
---
# Elasticsearch connection details
elasticsearch:
  enabled: true
  host: http://elastic:9200
  username: elastic
  password: elastic

# Employee application port
employee:
  api_port: "8083"
```

```shell
# For compiling code
make build
```

```shell
# For running code locally
export CONFIG_FILE=/path/to/config.yaml
./employee
```

```shell
make build-image
```
