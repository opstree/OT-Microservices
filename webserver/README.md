# Webserver

Webserver is a nginx based proxy which proxies the frontend application.

## Endpoints

The available endpoints for this application are:-

|**ENDPOINT**|**REQUEST TYPE**|**DESCRIPTION**|
|------------|----------------|---------------|
| `/` | GET | root will serve the landing page of the frontend application |
| `/healthz` | GET | healthz will return the status of nginx service |

## Quickstart

```shell
make image
```
