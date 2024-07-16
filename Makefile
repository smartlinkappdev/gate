build:
	docker build -t my-go-app .

run:
	go run cmd/gate/main.go

recreate:
	#docker buildx build --platform linux/amd64 -t gate .
	#docker buildx build --platform linux/amd64 -t gate .
	docker build -t gate .
	docker tag gate akaletr/gate:latest
	docker push akaletr/gate:latest
	docker-compose up

migrate:
	go run cmd/gate/migrate.go


gen:
	go run cmd/gate/m.go