# Go kit Demo

[![Build Status](https://travis-ci.org/arielcr/go-kit-demo.svg?branch=master)](https://travis-ci.org/arielcr/go-kit-demo)

This is a simple implementation of a micro service using Go Kit.

### Usage
Run this command to start the server:
```sh
$ go run cmd/main.go
```

### Routes
**Check status**

```
localhost:8080/status
```

**Say Hello**

```
localhost:8080/hello
```

With this request data:
```json
{
	"name": "You"
}
```

**Say Hello... with more stuff**

```
localhost:8080/complex
```

With this request data:
```json
{
	"name": "You",
	"phone": "1234",
	"address": "Here",
	"zipcode": "12222"
}
```