# executing the application
exec:
	go run cmd/main.go

# building the application
build:
	go build -o /execute cmd/main.go

# create docker image
docker-build:
	docker build . -t amirhossein21/event-man

# docker push to docker hub
docker-push:
	docker push amirhossein21/event-man
