FROM maven:latest as builder
MAINTAINER Opstree Solutions
WORKDIR /java/
COPY pom.xml /java/
COPY src /java/src/
RUN mvn clean package

FROM alpine:latest
MAINTAINER Opstree Solutions
USER root
RUN apk update && \
    apk add openjdk17

COPY --from=builder /java/target/salary-0.3.0-RELEASE.jar /app/salary.jar
EXPOSE 8080
ENTRYPOINT ["/usr/bin/java", "-jar", "/app/salary.jar"]
