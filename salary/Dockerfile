# Stage 1: Build the application using Maven and Java 17
FROM maven:3.8.8-eclipse-temurin-17 AS builder
LABEL maintainer="Opstree Solutions"

WORKDIR /java
COPY pom.xml .
COPY src ./src

RUN mvn clean package

# Stage 2: Run the application with a minimal runtime image
FROM eclipse-temurin:17-jdk-alpine
LABEL maintainer="Opstree Solutions"

WORKDIR /app

COPY --from=builder /java/target/salary-0.3.0-RELEASE.jar /app/salary.jar
COPY ./config.yaml /app/config/config.yaml
ENV CONFIG_FILE=/app/config/config.yaml

# Expose the app port
EXPOSE 8080

# Run the Spring Boot app
ENTRYPOINT ["java", "-jar", "/app/salary.jar"]