# Salary

Salary is also a golang based application which creates and manages employee's salary information.

## Database

- [Elasticsearch](../elasticsearch) => Salary application stores information in easticsearch

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DESCRIPTION**|
|------------------------|---------------|
| CONFIG_FILE | Path of configuration file |

## Quickstart

```yaml
---
# elasticsearch connection details
elasticsearch:
  enabled: true
  host: http://elastic:9200
  username: elastic
  password: elastic

# Salary application api port
salary:
  api_port: "8082"
```

```shell
# For compiling code
make build
```

```shell
# For running code locally
export CONFIG_FILE=/path/to/config.yaml
./salary
```

```shell
make image
```
