build:
	@go build -o bin/onionhead -ldflags "-s -w" cmd/onionhead/*