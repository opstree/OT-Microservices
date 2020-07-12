# Employee

Employee microservice is also designed in Golang to manage employee's information.

## Database

- [Elasticsearch](../elasticsearch) => Employee application stores information in easticsearch

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DESCRIPTION**|
|------------------------|---------------|
| CONFIG_FILE | Path of configuration file |

## Endpoints

The available endpoints for this application are:-

|**ENDPOINT**|**REQUEST TYPE**|**DESCRIPTION**|
|------------|----------------|---------------|
| `/employee/create` | POST | create endpoint collects the JSON payload of request and write the data to elasticsearch database. |
| `/employee/search` | GET | search endpoint fetch the information from elasticsearch for a particular user by sending query in URL. |
| `/employee/search/all` | GET | search/all endpoint fetch the information for all users from elasticsearch server. |
| `/employee/search/roles` | GET | search/roles endpoint fetch the information for user's role from elasticsearch|
| `/employee/search/location` | GET | search/location endpoint fetch the information for user's location from elasticsearch  |
| `/employee/search/status` | GET | search/status endpoint fetch the information for user's status from elasticsearch. |
| `/employee/healthz` | GET | healthz endpoint checks the DB connectivity and tells that application is ready to serve the requests or not. |

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
