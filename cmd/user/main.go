package main

import (
	"github.com/Hareshutit/ShopEase/internal/user"

	config "github.com/Hareshutit/ShopEase/config/user"
)

func main() {
	cfg := config.CreateConfig()

	user.Run(cfg)
}
