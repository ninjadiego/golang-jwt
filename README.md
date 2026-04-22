# golang-jwt

API REST de ejemplo en Go que implementa autenticacion basada en JSON Web Tokens (JWT). Creado como primer paso para aprender desarrollo backend con Go.

## Caracteristicas

- Registro y login de usuarios
- Generacion de tokens JWT al autenticarse
- Middleware de validacion de JWT para proteger rutas
- Estructura modular y facil de entender

## Tecnologias

- Go (Golang)
- JWT para autenticacion
- HTTP estandar de Go

## Instalacion

Clona el repositorio y ejecuta:

git clone https://github.com/ninjadiego/golang-jwt.git
cd golang-jwt
go mod tidy
go run main.go

El servidor quedara escuchando en el puerto 8080.

## Endpoints

- POST /login -> Autentica al usuario y devuelve un JWT
- GET /protected -> Ruta protegida, requiere token JWT valido

## Objetivo

Este repo forma parte de mi camino aprendiendo backend con Go. Mi meta es dominar patrones como autenticacion, autorizacion y APIs REST limpias.

## Autor

ninjadiego - https://github.com/ninjadiego- - - - - - - - - 
