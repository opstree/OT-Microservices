#!/bin/bash

BASE_COMMAND="python3 notification_api.py"
CONFIG_FILE=${CONFIG_FILE:-"/app/config.yaml"}

if [ -f "${CONFIG_FILE}" ]; then
    echo "Reading properties from config file"
else
echo """---
smtp:
  from: \"${FROM}\"
  username: \"${SMTP_USERNAME}\"
  password: \"${SMTP_PASSWORD}\"
  smtp_server: \"${SMTP_SERVER}\"
  smtp_port: \"${SMTP_PORT}\"

elasticsearch:
  username: \"${ELASTIC_USERNAME}\"
  password: \"${ELASTIC_PASSWORD}\"
  host: \"${ELASTIC_HOST}\"
  port: \"${ELASTIC_PORT}\"
""" >> "${CONFIG_FILE}"
fi

exec env CONFIG_FILE=${CONFIG_FILE} ${BASE_COMMAND}
