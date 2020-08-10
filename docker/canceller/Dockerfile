FROM golang:1.14.4 AS build

COPY . /src

WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod vendor -a -installsuffix cgo -o /bin/canceller ./cmd/canceller/main.go

# I can't do slim image because I should wait postgres in docker compose with bash...:(
# FROM scratch
# COPY --from=build /bin/canceller /app/canceller
# COPY /src/internal/infrastructure/config/config.yml /app/config.yml

RUN mkdir /app &&\
    cp /bin/canceller /app/canceller &&\
    cp /src/internal/infrastructure/config/config.yml /app/config.yml

ENV CONFIG_FILE=/app/config.yml

WORKDIR /app

EXPOSE 80

ENTRYPOINT ["./canceller"]
