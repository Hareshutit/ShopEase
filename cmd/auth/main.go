package main

import (
	config "github.com/Hareshutit/ShopEase/config/auth"
	"github.com/Hareshutit/ShopEase/internal/auth"
)

func main() {
	cfg := config.CreateConfig()

	auth.Run(cfg)
}
