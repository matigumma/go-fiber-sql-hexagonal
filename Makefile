# Variables comunes
APP_NAME=go-fiber-sql-hexagonal
VERSION=0.0.1
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_CLEAN=$(GO_CMD) clean
CGO_ENABLED=0  # docs: https://pkg.go.dev/cmd/cgo

# Construir el proyecto
build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO_BUILD) -o ./build/$(APP_NAME)-dev

# Ejecutar las pruebas
test:
	CGO_ENABLED=$(CGO_ENABLED) $(GO_TEST) ./...

# Limpiar los archivos binarios
clean:
	$(GO_CLEAN)
	rm -rf ./build

# Construir el binario de producción
release:
	mkdir -p ./build/production
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) $(GO_BUILD) -o ./build/production/$(APP_NAME)-$(VERSION)
