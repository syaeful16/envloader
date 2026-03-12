package envloader

import (
	"os"
	"testing"
	"time"
)

type TestConfig struct {
	AppName string        `env:"APP_NAME" default:"TestApp"`
	Port    int           `env:"APP_PORT" default:"8080"`
	Debug   bool          `env:"APP_DEBUG" default:"false"`
	Expiry  time.Duration `env:"TOKEN_EXPIRY" default:"5m"`
	Origins []string      `env:"ORIGINS"`
}

func TestLoad(t *testing.T) {

	os.Setenv("APP_NAME", "DemoApp")
	os.Setenv("APP_PORT", "9000")
	os.Setenv("APP_DEBUG", "true")
	os.Setenv("TOKEN_EXPIRY", "10m")
	os.Setenv("ORIGINS", "a.com,b.com")

	cfg := &TestConfig{}

	err := Load(cfg)

	if err != nil {
		t.Fatal(err)
	}

	if cfg.AppName != "DemoApp" {
		t.Fatal("wrong app name")
	}

	if cfg.Port != 9000 {
		t.Fatal("wrong port")
	}

	if cfg.Debug != true {
		t.Fatal("wrong debug")
	}

	if cfg.Expiry != 10*time.Minute {
		t.Fatal("wrong expiry")
	}

	if len(cfg.Origins) != 2 {
		t.Fatal("slice parsing failed")
	}
}
