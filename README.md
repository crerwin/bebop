# Bebop

Bebop is a basic Hello World-stype web service used to demonstrate various CI/CD concepts.

# Building Bebop

To build the Go code locally, from the root of the repo run `go install ./...`.  You can then execute Bebop using the `bebop` command:

```
$ bebop      
Serve or interact with the Bebop API

Usage:
  bebop [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  serve       Serve the Bebop API
  version     Displays the Bebop version

Flags:
      --debug   Enable debug logging
  -h, --help    help for bebop

Use "bebop [command] --help" for more information about a command.
```

# Serving the Bebop API

Serve the Bebop API with the `bebop serve` command:

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/crerwin/bebop/internal/server.setRouter.func1 (3 handlers)
[GIN-debug] GET    /api/hello                --> github.com/crerwin/bebop/internal/server.setRouter.func2 (3 handlers)
[GIN-debug] GET    /api/version              --> github.com/crerwin/bebop/internal/server.setRouter.func3 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

# Building and running in Docker

You can build a Bebop Docker image using the included `Dockerfile` with `docker build`:
```
docker build -t bebop .
```

You can then run Bebop in a container:

```
$ docker run -p 8080:8080 bebop serve
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/crerwin/bebop/internal/server.setRouter.func1 (3 handlers)
[GIN-debug] GET    /api/hello                --> github.com/crerwin/bebop/internal/server.setRouter.func2 (3 handlers)
[GIN-debug] GET    /api/version              --> github.com/crerwin/bebop/internal/server.setRouter.func3 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```