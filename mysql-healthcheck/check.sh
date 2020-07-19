#!/bin/bash

MYSQL_USERNAME=${MYSQL_USERNAME:-"root"}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-"password"}
MYSQL_HOST=${MYSQL_HOST:-"empms-db"}
MYSQL_PORT=${MYSQL_PORT:-3306}
MYSQL_DATABASE=${MYSQL_DATABASE:-"attendancedb"}
SLEEP_INTERVAL=${SLEEP_INTERVAL:-10}

while true
do
    if ! mysql -h ${MYSQL_HOST} -u ${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -e "use ${MYSQL_DATABASE}"; then
        echo "Waiting for ${MYSQL_DATABASE} to be available"
    else
	sleep ${SLEEP_INTERVAL}
        break
    fi
done
