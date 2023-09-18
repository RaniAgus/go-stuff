# hello

### Instalar una aplicación

```bash
go install
go install .
go install github.com/RaniAgus/hello
```

### Ejecutar la app instalada
```bash
hello
```

### Importar paquetes desde tu módulo
- Creamos una subcarpeta `morestrings`
- Usamos la ruta completa del package: `github.com/RaniAgus/hello/morestrings`

Si ejecutamos `go build` desde el módulo eso solo actualiza una caché.

### Importar paquetes remotos
- Ejecutamos `go mod tidy` para descargarlos:

```bash
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.9
```

- Para borrar todos los módulos descargados, ejecutamos `go clean -modcache`
