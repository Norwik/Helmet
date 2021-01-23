FROM golang:1.16.6

ARG DRIFTER_VERSION=0.1.0

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN mkdir -p /app/var/storage
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Spacemanio/Helmet/releases/download/v${DRIFTER_VERSION}/walrus_${DRIFTER_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./helmet version

CMD ["./helmet", "server", "-c", "/app/configs/config.dist.yml"]