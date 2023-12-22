package config

import "os"

type Config struct {
	PublicKey  string
	PrivateKey string
}

func NewConfig() Config {
	return Config{
		PublicKey:  os.Getenv("IMAGEKIT_PUBLICKEY"),
		PrivateKey: os.Getenv("IMAGEKIT_PRIVATEKEY"),
	}
}
