# Datetime Server using Go

# Project Description

This project is an implementation of datetime server that returns the current date and time using Golang with two different approaches. The first one is using http package and the second one is using Gin framwork. The project extends beyond mere implementation; it delves into docker files, docker images, and docker compose. Each implementation has its own docker file and there is a docker compose file that runs both implementations at the same time. Running and configuring each process is made easy by a Makefile that hosts targets of all the commands that a person who wants to run the project will need.

# How to Install and Run

There are different approaches for building and running the project, and it depends on what the user wants. But before into the approaches, git clone the repository using:

```bash
git clone git@github.com:codescalersinternships/datetime-server-eyadhussein.git
```

## Approach 1 - Building and Executing Binaries

### Building and Running Each Implmentation Alone

- HTTP Package

1- To build binary, run:

```bash
make build-http-bin
```

2- To run binary, run:

```bash
./datetimehttp
```

- Gin Package

1- To build binary, run:

```bash
make build-gin-bin
```

2- To run gin server binary, run:

```bash
./datetimegin
```

3- To run http server binary, run:

```bash
./datetimehttp
```

### Building All Binaries at the Same Time

1- To build binaries, run:

```bash
make build-all-bin
```

2- To run gin server binary, run:

```bash
./datetimegin
```

3- To run http server binary, run:

```bash
./datetimehttp
```

## Approach 2 - Using Docker Images

## Building and Running Each Image Separately

- HTTP Package

1- To build image, run:

```bash
make build-http-img
```

2- To lauch container, run:

```bash
make run-http
```

- Gin Package

1- To build image, run:

```bash
make build-gin-img
```

2- To lauch container, run:

```bash
make run-gin
```

## Building and Running Both Services at the Same Time

To build and run both services at the same time

```bash
make run-all
```

# How to Use

- There are utilties for formatting and linting the project

To format the project, run:

```bash
make format
```

To run lintter on the project, run:

```bash
make lint
```

After building and running the binaries using one of the previously stated approaches, it is possible to make a GET request to /datetime endpoint and recieve a response.

This will be illustrated using curl:
If the Approach 1 is chosen, then use port 8080 when running either binaries.

```bash
curl http://localhost:8080/datetime
```

Response:

```bash
2024-07-04 15:11:44
```

If the Approach 2 is chosen, then:

If running the two services is not the desirable behavior, then use:

```bash
curl http://localhost:8080/datetime
```

Response:

```bash
2024-07-04 15:11:44
```

If running two services at the same time is the desirable behavior then to run:

- HTTP Package will run on port 5000

```bash
curl http://localhost:5000/datetime
```

Response:

```bash
2024-07-04 15:11:44
```

- Gin Package wil run on port 3000

```bash
curl http://localhost:3000/datetime
```

Response:

```bash
2024-07-04 15:11:44
```

# How to Test

Run

```bash
make test
```
