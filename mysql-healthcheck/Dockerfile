FROM alpine:latest

RUN apk update && \
    apk add mysql-client bash

COPY ./check.sh /app/

WORKDIR /app

ENTRYPOINT ["./check.sh"]
