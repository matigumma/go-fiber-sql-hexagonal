# GO - Fiber V3 - SQL - Hexagonal Architecture boilerplate

Esta es una aplicación de ejemplo que utiliza el framework [Fiber](https://gofiber.io/) para crear una API que se conecta a una base de datos SQL Server.

## Requisitos

- Go 1.16 o superior
- SQL Server
- Archivo `.env` con las siguientes variables:
  - `DB_USER`: Usuario de la base de datos
  - `DB_PASSWORD`: Contraseña de la base de datos
  - `DB_SERVER`: URL del servidor de la base de datos
  - `DB_PORT`: Puerto del servidor de la base de datos
  - `DB_DATABASE`: Nombre de la base de datos

## Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/matigumma/go-fiber-sql-hexagonal.git
    cd go-fiber-sql-hexagonal
    ```

2. Instala las dependencias:
    ```sh
    go mod tidy
    ```

3. Crea un archivo `.env` en la raíz del proyecto con las variables de entorno necesarias.

## Uso

1. Inicia la aplicación:
    ```sh
    go run main.go
    ```

2. La aplicación estará disponible en `http://localhost:3000`.

## Endpoints 
- TODO

## Estructura del Proyecto

- `main.go`: Archivo principal de la aplicación.
- `public/`: Directorio para archivos estáticos.
- `internal/`: Código interno de la aplicación.
  - `app/`: Lógica de la aplicación.
    - `api/`: Lógica del caso de uso.
  - `ports/`: Interfaces de los puertos.
  - `adapters/`: Implementaciones de infraestructura.
    - `left/`: Driving side.
      - `rest/`: Servicio Rest.
    - `right/`: Driven side.
      - `db/`: DB adapter.

## Dependencias

- [Fiber](https://gofiber.io/)
- [Go MSSQL Driver](https://github.com/denisenkom/go-mssqldb)
- [godotenv](https://github.com/joho/godotenv)

## Makefile Commands

This project uses a Makefile to manage common tasks. Below are the available commands:

### Variables
- `APP_NAME`: The name of the application (default: `go-fiber-sql-hexagonal`)
- `VERSION`: The version of the application (default: `0.0.1`)
- `GO_CMD`: The Go command (default: `go`)
- `CGO_ENABLED`: CGO enable flag (default: `0`)

### Commands

#### `make build`
Builds the project with CGO disabled and outputs the binary to the `./build` directory with the name `$(APP_NAME)-dev`.

#### `make test`
Runs all tests in the project with CGO disabled.

#### `make clean`
Cleans up the binary files and removes the `./build` directory.

#### `make release`
Builds the production binary with CGO disabled, using the current OS and architecture. The binary is output to the `./build/production` directory with the name `$(APP_NAME)-$(VERSION)`.

### Usage Examples

- To build the project:
  ```sh
  make build
  ```

- To run tests:
  ```sh
  make test
  ```

- To clean up binaries:
  ```sh
  make clean
  ```

- To create a production release:
  ```sh
  make release
  ```

### Notes
- Ensure you have Go installed and properly set up in your environment.
- The `CGO_ENABLED` variable is set to `0` by default to disable CGO. You can modify this if your project requires CGO.
