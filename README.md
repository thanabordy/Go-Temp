# GO-TEMP
## Get started
```bash
# Init Project
$ go mod vendor
$ go run . keygen
$ cp .env.example .env

# HTTP/2 Server
$ go run . http

# GRPC Server
$ go run . grpc
```
## Structure
  - app
    - console
    - grpc
      - pb
    - http
      - controller
    - model
  - config
  - lib
  - router
  - storage
    - cert