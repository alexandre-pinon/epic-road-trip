# Epic road trip API

## Installation

- You first need [Go](https://golang.org/) installed (**version 1.18+ is required**), then you can use the below Go command to install the required packages.

```sh
$ go get
```

## Usage

You can run the server by using the below Go command at the root of the project

```sh
$ go run main.go
```

## Tests

You can run the tests by using the below Go command at the root of the project

```sh
$ go test ./...
```

You can see all available flags by using

```sh
$ go help test
```

## Documentation

1. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
go install github.com/swaggo/swag/cmd/swag
```

2. Run the [Swag](https://github.com/swaggo/swag) command at the root of the project

```sh
swag init
```

3. Access the documentation by visiting this [URL](http://localhost:8000/docs.html)