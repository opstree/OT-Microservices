# Attendance

Attendance is a microservice which is designed in Golang to manage employee's attendance information.

## Database

- [MySQL](../mysql) => Attendance application store information in MySQL database

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DESCRIPTION**|
|------------------------|---------------|
| CONFIG_FILE | Path of configuration file |

## Quickstart

```yaml
# Application port on which application will listen
attendance:
  api_port: "8081"

# Mysql connection details
mysql:
  enabled: true
  db_name: "attendancedb"
  host: "mysql:3306"
  username: root
  password: password
```

```shell
# For compiling code
make build
```

```shell
# For running code locally
export CONFIG_FILE=/path/to/config.yaml
./attendance
```

```shell
make build-image
```