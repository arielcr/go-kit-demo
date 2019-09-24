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
localhost:8000/status
```

**Say Hello**

```
localhost:8000/hello
```

With this request data:
```json
{
	"name": "You"
}
```
