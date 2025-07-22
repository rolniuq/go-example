package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	StoragePath string
	ReaderPath  string
	WriterPath  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{}
}

func (c *Config) Load() (*Config, error) {
	c.AppName = os.Getenv("APP_NAME")
	c.StoragePath = os.Getenv("STORAGE_PATH")
	c.ReaderPath = os.Getenv("READER_PATH")
	c.WriterPath = os.Getenv("WRITER_PATH")

	return c, nil
}
