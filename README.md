# simplQ-golang

## Local Setup Notes

1. Golang setup: https://go.dev/doc/install
2. Install docker and docker-compose: https://docs.docker.com/compose/install/
2. Start development server:
```
git clone git@github.com:SimplQ/simplQ-golang.git
cd simplQ-golang/
make dev
```

Mongo Express UI will be available at http://localhost:8081/

We use Visual Studio Code as the IDE but this is not a strict requirement. [Azure Functions](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurefunctions) plugin can be used to deploy code to Azure, but this is optional.
