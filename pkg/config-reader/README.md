# config-reader package

**Important**: currently tested only with strings and integer values.

## Define a properties struct

```go
package config

type Properties struct {
    Server   Server   `json:"server"`
    Database Database `json:"database"`
}

type Server struct {
    Port int `json:"port"`
}

type Database struct {
    Driver   string `json:"driver"`
    Username string `json:"username"`
    Password string `json:"password"`
    Name     string `json:"name"`
}
```

> _NOTE_: It's important to define the struct tag `json`

## Provide a configuration file in JSON format

```json
{
  "server": {
    "port": 8080
  },
  "database": {
    "driver": "postgres",
    "username": "postgres",
    "password": "password",
    "name": "postgres"
  }
}
```

## Read it

```go
var props config.Properties
reader := config_reader.DefaultConfigReader(&props)
err := reader.Read()
if err != nil {
    log.Printf("Could not read properties")
}
```

## More features

### Override/set anything with environment variables

> _NOTE_: This is currently a work in progress

Use the JSON tags to build environment variables, separating tags with `_`. For example:

- `SERVER_PORT=8081` would set the server port
- `DATABASE_PASSWORD` would set the database password

> _NOTE_: Environment variables override the values from the configuration file.

### Set your own properties instead of using the DefaultConfigReader() builder

```go
var props config.Properties{}
reader := config_reader.ConfigReader{
	File:        "/app/config.json",
    Environment: false,
	Props:       &props
}
err := reader.Read(&props)
```