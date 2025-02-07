package main

import (
	"backend/cmd"
	"backend/config"
	"backend/db"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig()
	port := cfg.AppPort

	db := db.DBconn(ctx, cfg)

	server := cmd.SetupServer(db, *cfg)

	if err := server.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
