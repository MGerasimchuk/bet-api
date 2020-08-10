# Bet API

## Requirements

* [Docker Engine](https://docs.docker.com/engine/) (tested on 19.03.8)
* [Docker Compose](https://docs.docker.com/compose/) (tested on 1.25.5)
* [Bash](https://www.gnu.org/software/bash/) (tested on 5.0.17)
* [Make](https://www.gnu.org/software/make/) (tested on 3.81)
* Available 80 and 5432 ports

## How to use

**Run this command:**
```
make good
```

After, you can send requests to Service HTTP API like this:
```
curl --location --request POST 'localhost/account/1bb63387-a1df-4158-96ec-8c6077f689ec/bet' \
--header 'Source-Type: game' \
--header 'Content-Type: application/json' \
--data-raw '{"state":"win","amount":"200","transactionId":"e5467485-3bdf-4c59-9e40-563d147a0e2b"}'
```

or like this:
```
curl --location --request POST 'localhost/account/1bb63387-a1df-4158-96ec-8c6077f689ec/bet' \
--header 'Source-Type: game' \
--header 'Content-Type: application/json' \
--data-raw "{\"state\":\"win\",\"amount\":\"200\",\"transactionId\":\"$(uuidgen)\"}"
```


or this:
```
while true; do curl --location --request POST 'localhost/account/1bb63387-a1df-4158-96ec-8c6077f689ec/bet' \
--header 'Source-Type: game' \
--header 'Content-Type: application/json' \
--data-raw "{\"state\":\"win\",\"amount\":\"200\",\"transactionId\":\"$(uuidgen)\"}"; sleep 1; done
```

and view data in database like this:
```
docker-compose run migrations-up bash -c 'psql -P pager=off -c "SELECT * FROM transactions ORDER BY created_at DESC"' ;\
docker-compose run migrations-up bash -c 'psql -c "SELECT * FROM accounts"'
```

## Tests

Run tests with generating html coverage report:
```
docker run -v ${PWD}:/src golang:1.14.4 /bin/bash -c 'cd /src && GO111MODULE=on go test -mod vendor -covermode=count -coverprofile=cover.out -v ./... && go tool cover -html=cover.out -o=cover.html'
``` 

After executing, you can open generated [cover.html](https://htmlpreview.github.io/?https://github.com/mgerasimchuk/bet-api/blob/master/cover.html) file in browser for check coverage 

## Configuration

See config file: [internal/infrastructure/config/config.yml](internal/infrastructure/config/config.yml)

This config file, mounts to docker containers automatically(without rebuild), when you run `make good`

## Rebuild

You can make any changes with source code or with images, but after this, for rebuild and restart app, you should run: 
```
make rebuild
```

## Details of realisation

As an architectural approach was used - [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Some things for the project layout was taken from - [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Original Technical Task Description

The main goal of this test task is to develop the application for processing the incoming requests from the 3d-party providers.
The application must have an HTTP URL to receive incoming POST requests.
To receive the incoming POST requests the application must have an HTTP URL endpoint.

Technologies: Golang + Postgres.

Requirements:
1. Processing and saving incoming requests.

Imagine that we have a user with the account balance.

Example of the POST request:
```
POST /your_url HTTP/1.1
Source-Type: client
Content-Length: 34
Host: 127.0.0.1
Content-Type: application/json
{"state": "win", "amount": "10.15", "transactionId": "some generated identificator"}
```

Header "Source-Type" could be in 3 types (game, server, payment). This type probably can be extended in the future.

Possible states (win, lost):
1. Win requests must increase the user balance
2. Lost requests must decrease user balance.
Each request (with the same transaction id) must be processed only once.

The decision regarding database architecture and table structure is made to you.

You should know that account balance can't be in a negative value.
The application must be competitive ability.

2. Post-processing
Every N minutes 10 latest odd records must be canceled and balance should be corrected by the application.
Cancelled records shouldn't be processed twice.

3. The application should be prepared for running via docker containers.
