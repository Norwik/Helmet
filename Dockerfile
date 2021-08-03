FROM golang:1.16.6 as build

RUN mkdir -p /app
RUN apt-get update

WORKDIR /app

COPY ./ ./

RUN go build -v -ldflags="-X 'main.version=v1.0.21'" helmet.go

FROM golang:1.16.6

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN apt-get update

WORKDIR /app

COPY --from=build /app/helmet /app/helmet
COPY --from=build /app/config.dist.yml /app/configs/config.dist.yml

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./helmet version

CMD ["./helmet", "server", "-c", "/app/configs/config.dist.yml"]
