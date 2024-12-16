# debug-server
A minimalistic http(s) server to aid remote access http debugging.

## Usage

Build a linux executable:
```bash
CGO_ENABLED=0 GOOS=linux go build -o server main.go
```
Then run the server with:
```bash
# start the http server on port 8080
./server --port 8080

# start the http(s) server on port 443
./server --tls --port 443 --cert <path_to_tls_cert.pem> --key <path_to_tls_key.pem>
```

Run locally (assuming you already have golang installed in your system)
```bash
go run main.go
```

