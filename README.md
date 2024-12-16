# debug-server
A minimalistic http(s) server to aid remote access http debugging.

## Getting the binary for your os

### Option 1 (easier)
Download the latest [release](https://github.com/Gr3at/debug-server/releases) for your os-arch and use directly on your machine.

### Option 2 (more steps)
Build the executable for the target os (e.g. linux):
```bash
CGO_ENABLED=0 GOOS=linux go build -o server main.go
```

## Usage

Run the server with optional flags:
```bash
# start the http server on port 8080
./server --port 8080

# start the http(s) server on port 443
./server --tls --port 443 --cert <path_to_tls_cert.pem> --key <path_to_tls_key.pem>
```
