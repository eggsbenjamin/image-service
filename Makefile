install:
	glide install

unit_tests:
	echo TODO

integration_tests:
	echo TODO

system_tests:
	go test -v ./systemtest -tags=system

docker_tests:
	docker-compose -f docker/docker-compose-tests.yml -p image-service rm -v -f
	docker-compose -f docker/docker-compose-tests.yml -p image-service up --abort-on-container-exit

prod_build:
	GOOS=linux GOARCH=386 go build -o bin/image-service cmd/main.go

prod_image:
	make docker_tests && make prod_build && docker build -t image-service -f docker/Dockerfile .
