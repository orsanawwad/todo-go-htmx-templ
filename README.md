# The templ of HTMX and Go

A TODO project

### How?
Install air
```
go install github.com/cosmtrek/air@latest
```
Install templ
```
go install github.com/a-h/templ/cmd/templ@latest
```

Tidy up
```
go mod tidy
```

Run
```
air -c .air.toml
```

### What?

- templates are defined in handlers/templates as templ files
- templ converts templ files to go files with

    `templ generate --path ./internal/templates/`

- assets under ./internal/handlers/assets/* are bundled using go:embed in ./internal/handlers/assets.go, then exposed as its own route under /assets
