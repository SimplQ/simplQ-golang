all: AzureFunctions/handler

AzureFunctions/handler: cmd/handler/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o $@ $<

clean:
	rm AzureFunctions/handler