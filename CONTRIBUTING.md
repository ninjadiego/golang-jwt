# Guía de contribución

¡Gracias por tu interés en contribuir a `golang-jwt`! Este proyecto es una implementación de práctica de autenticación con JSON Web Tokens en Go, y las contribuciones son bienvenidas.

## Cómo contribuir

1. Haz un fork del repositorio.
2. Crea una rama descriptiva para tu cambio:

   ```bash
   git checkout -b feat/nombre-de-tu-feature
   ```

3. Realiza los cambios y asegúrate de que el código compila:

   ```bash
   go build ./...
   ```

4. Si añades funcionalidad nueva, agrega pruebas en archivos `_test.go`.
5. Haz commit con un mensaje claro siguiendo el estilo Conventional Commits:

   - `feat:` para nuevas funcionalidades
   - `fix:` para correcciones de bugs
   - `docs:` para cambios en la documentación
   - `refactor:` para refactorizaciones
   - `test:` para añadir o modificar pruebas

6. Haz push a tu fork y abre un Pull Request.

## Estilo de código

- Sigue las convenciones idiomáticas de Go.
- Ejecuta `gofmt` antes de hacer commit.
- Ejecuta `go vet ./...` para detectar problemas comunes.
- Mantén las funciones pequeñas y con responsabilidad única.

## Reportar bugs

Si encuentras un bug, abre un issue incluyendo:

- Versión de Go que usas (`go version`).
- Sistema operativo.
- Pasos para reproducir el problema.
- Comportamiento esperado vs comportamiento real.
- Logs o mensajes de error relevantes.

## Sugerencias de mejora

Las ideas y propuestas son bienvenidas. Abre un issue con la etiqueta `enhancement` describiendo:

- El problema que resuelve.
- La solución propuesta.
- Alternativas consideradas.

## Seguridad

Si descubres una vulnerabilidad de seguridad, **no la reportes en un issue público**. Contacta directamente al mantenedor por email para coordinar la divulgación responsable.

## Código de conducta

Sé respetuoso con los demás colaboradores. Las críticas constructivas son bienvenidas; los ataques personales no.

¡Gracias por contribuir!
