# golang-dojo

Golang dojo

## First challenge mock-server

### Primera definicion:

- Levanta configs de paths a mockear de archivo json (config.json)
- Cada path apunta a algun archivo a donde buscar el response mockeado (TODO)

### Features futuros

- Panic con error con mensaje firendly (cuando ruta o file no esta definido )
- Eleccion de Response codes
- Eleccion de Headers
- Eleccion de Timeouts
- Guardar la config en una DB
- CRUD para los mocks
- Proxy para ir al mock o servicio real segun algun flag o header

### Run

```go run main.go```

### Build

```go build```
