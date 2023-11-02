# A Go, A Templ, and HTMX, walked into a bar

And said lets build a TODO app

### How
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

### But how

- templates are in handlers/templates
- templ converts them to go files with

    `templ generate --path ./internal/templates/`

- assets under ./internal/handlers/assets/* are bunbled in ./internal/handlers/assets.go, then exposed as its own route /assets
