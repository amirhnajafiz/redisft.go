# event-man

Event man is an event publisher implemented with Golang.

With event publisher, you can send events to your clients, and
also you can check the events' history.

Event man uses SSE protocol to create streams between clients
and the main server.

## Routes

## How to run?
To run the project, clone the repository:
```shell
git clone https://github.com/amirhnajafiz/event-man.git
```

Then start the project with docker compose:
```shell
docker-compose up -d
```

## Clients

## Deploy
If you want to deploy the project on a kubernetes cluster, use the following command:
```shell
helm install event-man-v1 ./charts
```
