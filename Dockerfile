FROM golang:1.16.6

ARG WALNUT_VERSION=0.1.0

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN mkdir -p /app/var/storage
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Walnut/releases/download/v${WALNUT_VERSION}/walrus_${WALNUT_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./walnut version

CMD ["./walnut", "server", "-c", "/app/configs/config.dist.yml"]