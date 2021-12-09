# Needs to be deployed to a linux environment
GOOS=linux
GOARCH=amd64

# Compile a static binary for better compatibility
CGO_ENABLED=0
LD_FLAGS='-extldflags "-static"'

all: build/azure_functions.zip

build/azure_functions.zip: AzureFunctions/handler
	mkdir -p build
	cd AzureFunctions && zip -r ../build/azure_functions.zip *
	# Display size of output zip
	@ls -lh build/azure_functions.zip | awk '{print $$5 " " $$9}'

AzureFunctions/handler: cmd/api/main.go
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags $(LD_FLAGS) -o $@ $<

clean:
	rm AzureFunctions/handler
	rm build/*

dev:
	# Install air if not already present 
	# https://github.com/cosmtrek/air
	command -v air > /dev/null || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH)/bin 
	air
