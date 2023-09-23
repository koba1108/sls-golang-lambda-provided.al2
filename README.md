# sls-golang-lambda-provided.al2
sls-golang-lambda-provided.al2

## build
```bash
GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o .bin/hello-world/bootstrap src/hello-world/main.go
GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o .bin/hello-world2/bootstrap src/hello-world2/main.go
```
