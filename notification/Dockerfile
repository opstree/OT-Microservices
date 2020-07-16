FROM python:3.6-alpine

MAINTAINER OpsTree Solutions

COPY ./ /app

WORKDIR /app

RUN apk add --update --no-cache g++ gcc libxslt-dev bash

RUN pip3 install -r requirements.txt

ENTRYPOINT ["./entrypoint.sh"]
