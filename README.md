# golang-dojo

Golang dojo

## First challenge mock-server

### Primera definicion:

- Levanta configs de paths a mockear de archivo json (config.json)
- Cada path apunta a algun archivo a donde buscar el response mockeado (TODO)

### Features futuros
- Guardar la config en una DB:
  - Load config de DB
  - API CRUD completo
  - Refactorizar (modularizar, repository, model, etc)
- Guardar mocks en una DB:
  - CRUD
- Validacion del config( deuda tecnica error que se vea bonito en esta lib o otra)
- Eleccion de Timeouts
- Catch de multiples tipos de error
- Test
- Dockerizar
  
### Run

```go run main.go```

### Build

```go build```
