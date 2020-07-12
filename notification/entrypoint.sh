#!/bin/bash

BASE_COMMAND="notification"
CONFIG_FILE=${CONFIG_FILE:-"/app/config.yaml"}


if [ -f "${CONFIG_FILE}" ]; then
    echo "Reading properties from config file"
else
echo """---
notification:
  api_port: \"${API_PORT}\"

smtp:
  from: \"${FROM}\"
  username: \"${SMTP_USERNAME}\"
  password: \"${SMTP_PASSWORD}\"
  smtp_server: \"${SMTP_SERVER}\"
  smtp_port: \"${SMTP_PORT}\"
""" >> "${CONFIG_FILE}"
fi

exec env CONFIG_FILE=${CONFIG_FILE} ${BASE_COMMAND}
