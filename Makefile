all: AzureFunctions/handler
	cd AzureFunctions && zip -r ../build/azure_functions.zip *
	ls -lh build/azure_functions.zip

AzureFunctions/handler: cmd/handler/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o $@ $<

clean:
	rm AzureFunctions/handler
	rm build/*