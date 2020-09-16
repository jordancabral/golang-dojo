# golang-dojo

Golang dojo

## First challenge mock-server

### Primera definicion:

- Levanta configs de paths a mockear de archivo json (config.json)
- Cada path apunta a algun archivo a donde buscar el response mockeado (TODO)

### Features futuros
- Validacion del config( deuda tecnica error que se vea bonito en esta lib o otra)
- Catch de multiples tipos de error
- Test
  
### Run

```go run main.go```

### Run with docker-compose

```docker-compose up -d```

### Run with docker-compose (with rebuild)

```docker-compose up -d --build```
