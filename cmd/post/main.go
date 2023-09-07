package main

import (
	config "github.com/Hareshutit/ShopEase/config/post"
	"github.com/Hareshutit/ShopEase/internal/post"
)

func main() {
	cfg := config.CreateConfig()

	post.Run(cfg)
}
