default: build

build:
	cd docker/mysql && docker-compose build
	cd docker/rethinkdb && docker-compose build
	cd docker/mongodb && docker-compose build
	cd docker/redis && docker-compose build

doc:
	go run generate.go
