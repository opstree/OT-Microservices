# Notification

Notification is a service which gets used to send mail notifications to employees.

## Dependency

Notification service is dependent on SMTP server. SMTP server details needs to be provided at application bootup.

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DEFAULT VALUE**|**DESCRIPTION**|
|------------------------|-----------------|---------------|
| CONFIG_FILE | `/app/config.yaml` | Path of configuration file |
| API_PORT | 8085 | Port number on which application will listen |
| FROM | - | From is the name of the sender from which mail will be sent |
| SMTP_USERNAME | - | Username for the SMTP server |
| SMTP_PASSWORD | - | Password of the SMTP username |
| SMTP_SERVER | - | Name of the SMTP server |
| SMTP_PORT | - | Port number on which SMTP server is listening |

## Quickstart

```yaml
---
# Notification service port
notification:
  api_port: "8085"

# SMTP connection details
smtp:
  from: "test@example.com"
  username: "test"
  password: "test"
  smtp_server: "test.smtp.com"
  smtp_port: "25"

```

```shell
# For compiling code
make build
```

```shell
# For running code locally
export CONFIG_FILE=/path/to/config.yaml
./notification
```

```shell
make build-image
```