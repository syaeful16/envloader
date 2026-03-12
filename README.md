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

## Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/syaeful16/envloader"
)

type Config struct {
	AppName string        `env:"APP_NAME" default:"MyApp"`
	AppPort int           `env:"APP_PORT" default:"8080"`
	Debug   bool          `env:"APP_DEBUG" default:"false"`
	Expiry  time.Duration `env:"TOKEN_EXPIRY" default:"10m"`
	Origins []string      `env:"ORIGINS"`
}

func main() {
	os.Setenv("APP_NAME", "DemoApp")
	os.Setenv("APP_PORT", "9000")
	os.Setenv("APP_DEBUG", "true")
	os.Setenv("TOKEN_EXPIRY", "10m")
	os.Setenv("ORIGINS", "a.com,b.com")

	cfg := &Config{}

	err := envloader.Load(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("AppName: %s, AppPort: %d, Debug: %v, Expiry: %v, Origins: %v\n",
		cfg.AppName, cfg.AppPort, cfg.Debug, cfg.Expiry, cfg.Origins)
}
```
