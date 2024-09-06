# EGTS Demo Application

This is a demo application written in Go (Golang) for sending and receiving EGTS (Era Glonass Transport System) data packets. The project demonstrates how to implement the EGTS protocol, commonly used for transmitting data between tracking devices and monitoring servers.

# Installation

Clone the repository:

```sh
git clone git@github.com:QED-tech/egts-example.git
cd egts-demo
```

Install dependencies:
```sh
go mod tidy
```

# Usage

To run the demo:

```sh
go run ./cmd/main.go --address=127.0.0.1:8080
```