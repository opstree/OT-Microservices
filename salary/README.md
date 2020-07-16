# Salary

Salary is also a golang based application which creates and manages employee's salary information.

## Database

- [Elasticsearch](../elasticsearch) => Salary application stores information in easticsearch

## Endpoints

The available endpoints for this application are:-

|**ENDPOINT**|**REQUEST TYPE**|**DESCRIPTION**|
|------------|----------------|---------------|
| `/salary/search` | GET | search will fetches the salary of a user based on URL query. |
| `/salary/healthz` | GET | healthz will check the DB connectivity and return the status of application that whether it can serve the requests or not. |
| `/salary/configure/liveness` | POST | configure/liveness healthcheck will slow the response of healthcheck to simulate liveness in Kubernetes |

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DESCRIPTION**|
|------------------------|---------------|
| CONFIG_FILE | Path of configuration file |
| DELAY_TIME | This will delay the bootup time of the application to simulate readiness in Kubernetes |

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
