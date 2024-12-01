# Store API

This is a simple store API written in Go using the Gin framework. It allows you to create, read, update, and delete products and categories.

## Installation

1. Clone the repository.
2. Run `make install` to download the dependencies.
3. Set the environment variable `DATABASE_URL` with your database connection string.
4. Run `make run` to start the server.

> **Note:** The server will start on port `8080` by default. You can change this by setting the `PORT` environment variable.

## Endpoints

- `GET /` - Returns a welcome message.
- `GET /ping` - Returns a pong message.
- `GET /products` - Returns a list of all products.

## Using the Makefile

- `make install` - Downloads the dependencies.
- `make build` - Builds the project binary.
- `make run` - Builds and runs the project binary.
- `make start` - Runs the project without building.
- `make test` - Runs the tests.
- `make clean` - Removes the generated binary files.
