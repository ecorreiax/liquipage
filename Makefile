build:
	go build -o liquipage main.go

docker-build:
	docker build . -t liquipage -f ./Dockerfile

docker-run:
	docker run --rm -it liquipage --help

clean:
	rm ./liquipage