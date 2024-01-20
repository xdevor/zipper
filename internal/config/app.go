package config

type AppConfig struct {
	Version string
}

var App AppConfig = AppConfig{
	Version: "v0.9.0",
}
