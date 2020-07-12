# Frontend

Frontend is written in ReactJS and gets served using nginx proxy, config can be found [here](../webserver)

## Endpoints

The available endpoints for this application are:-

|**ENDPOINT**|**REQUEST TYPE**|**DESCRIPTION**|
|------------|----------------|---------------|
| `/employee-add` | POST | employee-add endpoint collects the data from FORM for employee and sends to gateway |
| `/employee-list` | GET | employee-list endpoint fetches the information of all employees and shows on dashboard. |
| `/attendance-add` | POST | attendance-add endpoint collects the data from FORM for attendance and sends to gateway |
| `/attendance-list` | GET | employee-list endpoint fetches the information of all employee's attendance and shows on dashboard. |

## Quickstart

```shell
# For compiling code
make build
```

```shell
To run the code
serve -s build
```

```shell
make image
```
