package main

import (
    "log"

    "github.com/narwar-veer/teen-wallet-api/internal/config"
    "github.com/narwar-veer/teen-wallet-api/internal/database"
    "github.com/narwar-veer/teen-wallet-api/internal/router"
)
func main() {
    cfg := config.MustLoad()

    db := database.MustConnect(cfg)
    database.AutoMigrate(db)

    r := router.New(cfg, db)

    log.Printf("🚀  Server starting on %s", cfg.HTTPServer.Addr)
    if err := r.Run(cfg.HTTPServer.Addr); err != nil {
        log.Fatalf("unable to start server: %v", err)
    }
}