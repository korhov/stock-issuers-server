CGO_ENABLED=0 go build -v -ldflags="-s -w" -o stock-issuers-server ./cmd/main.go && ./stock-issuers-server -MODE=development
