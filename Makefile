all: AzureFunctions/handler
	mkdir -p build
	cd AzureFunctions && zip -r ../build/azure_functions.zip *
	ls -lh build/azure_functions.zip

AzureFunctions/handler: cmd/api/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o $@ $<

clean:
	rm AzureFunctions/handler
	rm build/*

dev:
	# Install air if not already present 
	# https://github.com/cosmtrek/air
	command -v air > /dev/null || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH)/bin 
	air
