# Notification

Notification is a service which gets used to send mail notifications to employees. This application runs in scheduled based frequency.

## Available Flags

```shell
$ python3 notification_api.py --help
usage: notification_api.py [-h] [-m MODE]

optional arguments:
  -h, --help            show this help message and exit
  -m MODE, --mode MODE  Mode in which application will run, options -
                        scheduled and external
```

With `--mode` flag you can control the behaviour of the application. In case of `scheduled` value it will run at every month beginning, and with `external` it will run once.

## Dependency

Notification service is dependent on SMTP server. SMTP server details needs to be provided at application bootup.

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DEFAULT VALUE**|**DESCRIPTION**|
|------------------------|-----------------|---------------|
| CONFIG_FILE | `/app/config.yaml` | Path of configuration file |
| FROM | - | From is the name of the sender from which mail will be sent |
| SMTP_USERNAME | - | Username for the SMTP server |
| SMTP_PASSWORD | - | Password of the SMTP username |
| SMTP_SERVER | - | Name of the SMTP server |
| SMTP_PORT | - | Port number on which SMTP server is listening |
| ELASTIC_USERNAME | - | Username for the elasticsearch server |
| ELASTIC_PASSWORD | - | Password of the elasticsearch server |
| ELASTIC_HOST | - | DNS name or IP of the elasticsearch server |
| ELASTIC_PORT | - | Port number on which elasticsearch is listening |

## Quickstart

```yaml
---
# SMTP connection details
smtp:
  from: "test@example.com"
  username: "test"
  password: "test"
  smtp_server: "test.smtp.com"
  smtp_port: "25"

# elasticsearch connection details
elasticsearch:
  username: "elastic"
  password: "elastic"
  host: "172.17.0.2"
  port: 9200
```

```shell
# For compiling code
make build
```

```shell
# For running code locally
export CONFIG_FILE=/path/to/config.yaml
python3 notification_api.py
```

```shell
make build-image
```