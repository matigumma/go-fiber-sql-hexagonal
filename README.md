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

