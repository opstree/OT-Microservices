FROM node:lts-stretch-slim

MAINTAINER Opstree Solutions

COPY . /app

WORKDIR /app/

RUN apt-get update -y && \
    apt-get install git -y

RUN npm install

RUN npm run build

RUN yarn global add serve

EXPOSE 5000

ENTRYPOINT ["serve", "-s", "build"]
