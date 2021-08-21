hello:
	echo "Hello"

build:
	go build -a

run:
	go run ./main.go

docker-run:
	docker-compose up --build