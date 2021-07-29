# -----------------------------------------------------
# STAGE 1 : CACHE DEPENDENCY
# -----------------------------------------------------

FROM maven:3.8.1-jdk-8-slim as build

ARG service_name
ENV SERVICE_NAME=$service_name

ARG deploy_env
ENV DEPLOY_ENV=$deploy_env

WORKDIR /source/${SERVICE_NAME}
COPY . .

RUN mvn -DskipTests clean package

# -----------------------------------------------------
# STAGE 2 : COPY BINAR
# -----------------------------------------------------

FROM java:8


ARG service_name
ENV SERVICE_NAME=$service_name

ARG deploy_env
ENV DEPLOY_ENV=$deploy_env

WORKDIR /service/${SERVICE_NAME}/bin

COPY --from=build /source/${SERVICE_NAME}/target/data-1.0-SNAPSHOT.jar /source/${SERVICE_NAME}/env.dev ./
# TODO temp until the component services are not deployed

EXPOSE 8080

CMD /bin/bash -c "source env.dev && java -jar data-1.0-SNAPSHOT.jar"
