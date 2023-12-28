# "Word of Wisdom" TCP server.

## About Project
This project runs on TCP based protocol that is protected from DDOS attack with the Proof of Work mechanism. The message uses delimiter and consist of two parts divided by |:
+ header -  integer
+ payload - string

## Start the project
### Requirement
+ Install [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/linux/) on your local machine to run the project
+  Install [Go 1.19+](https://go.dev/dl/) to start the project.

### Start project with docker-compose
+ use the below syntax to detach the docker from your terminal

```
docker-compose up --force-recreate --build server --build client -d
```

+ use the below syntax to run docker with detaching it from your terminal

```
docker-compose up --abort-on-container-exit --force-recreate --build server --build client
```

### You can run server and client individually:
+ use the below syntax to run server:
```
go run cmd/server/main.go
```
+ use the below syntax to run client:
```
go run cmd/client/main.go
```

## Proof of Work
I went with Hashcash algorithm for this project proof of work, among all other algorithm, Hashcash has simple implementation, and the documentation is easier to find only. Several articles with samples and descriptions. it is simple to validate server side.

it has disadvantages like slow compute time depending on the client machine and 

this could be solved with few implementations like caching and other mechanism.