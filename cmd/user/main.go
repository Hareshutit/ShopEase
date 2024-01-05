package main

import (
	"os"

	"github.com/Hareshutit/ShopEase/internal/user"
	"github.com/rs/zerolog"

	config "github.com/Hareshutit/ShopEase/config/user"
)

func main() {
	log := zerolog.New(os.Stdout)
	cfg := config.CreateConfig()
	user.Run(log, cfg)
}
