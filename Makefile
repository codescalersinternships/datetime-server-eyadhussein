BINARY_NAME_HTTP=datetimehttp
BINARY_NAME_GIN=datetimegin

build-http-bin:
	go build -o ${BINARY_NAME_HTTP} ./cmd/${BINARY_NAME_HTTP}

build-gin-bin:
	go build -o ${BINARY_NAME_GIN} ./cmd/${BINARY_NAME_GIN}

build-all-bin:
	go build -o ${BINARY_NAME_HTTP} ./cmd/${BINARY_NAME_HTTP}
	go build -o ${BINARY_NAME_GIN} ./cmd/${BINARY_NAME_GIN}


format:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...


build-http-img:
	sudo docker build -f Dockerfile.http -t datetimehttpimg .

build-gin-img:
	sudo docker build -f Dockerfile.gin -t datetimeginimg .

build-all-img:
	sudo docker build -f Dockerfile.gin -t datetimeginimg .
	sudo docker build -f Dockerfile.http -t datetimehttpimg .


run-http:
	sudo docker run --name datetimehttpcontainer datetimehttpimg
run-gin:
	sudo docker run --name datetimegincontainer datetimeginimg
run-all:
	sudo docker-compose up
