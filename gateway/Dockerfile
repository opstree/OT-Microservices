FROM maven:latest as builder
MAINTAINER Opstree Solutions
COPY . /java/
WORKDIR /java/
RUN mvn package

FROM alpine:latest
MAINTAINER Opstree Solutions
USER root
RUN apk update && \
    apk fetch openjdk8 && \
    apk add openjdk8 curl
COPY --from=builder /java/target/gateway-service.jar /app/gateway.jar
EXPOSE 8080
ENTRYPOINT ["/usr/bin/java", "-jar", "/app/gateway.jar"]
