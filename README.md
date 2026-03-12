# EnvLoader

EnvLoader is a lightweight Go library to load environment variables into structs using tags.

It supports default values, required validation, nested structs, slices, duration parsing, and environment prefixes.

## Features

- Struct tag based configuration
- Default values
- Required validation
- Nested struct support
- Prefix for environment variables
- Slice parsing (`[]string`)
- `time.Duration` parsing
- Simple and lightweight

---

## Installation

```bash
go get github.com/syaeful16/envloader
```

---

## Example

### Config Struct

```go
type Config struct {
	AppName string        `env:"APP_NAME" default:"MyApp"`
	AppPort int           `env:"APP_PORT" default:"8080"`
	Debug   bool          `env:"APP_DEBUG" default:"false"`
	Expiry  time.Duration `env:"TOKEN_EXPIRY" default:"10m"`
	Origins []string      `env:"ORIGINS"`
}
```

### Example .env File

```
APP_NAME=DemoApp
APP_PORT=9000
APP_DEBUG=true
TOKEN_EXPIRY=10m
ORIGINS=a.com,b.com
```

### Load configuration

```go
package main

import (
    "log"

    "github.com/joho/godotenv"
    "github.com/syaeful16/envloader"
)

func main() {

    godotenv.Load()

    cfg := &Config{}

    err := envloader.Load(cfg)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(cfg.App.Name)
}
```

### Supported Tags

| Tag        | Description                      |
| ---------- | -------------------------------- |
| `env`      | Environment variable name        |
| `default`  | Default value (optional)         |
| `required` | Field is required (optional)     |
| `prefix`   | Prefix for environment variables |

Example:

```go
Port int `env:"PORT" prefix:"DB" default:"5432"`
```

Env variable

```
DB_PORT=5432
```
