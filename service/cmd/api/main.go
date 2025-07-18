package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/afrze/c9-private-server/service/internal/config"
	"github.com/afrze/c9-private-server/service/internal/ctx"
	"github.com/afrze/c9-private-server/service/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)

	ctx := ctx.Ctx{
		Mongo: db.InitMongo(cfg.MongoURI),
		MSSQL: db.InitMSSQL(cfg.MSSQLConn),
	}

	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	go func() {
		log.Printf("[api] App Listening on %s", cfg.ListenAddr)
		if err := r.Run(cfg.ListenAddr); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[api] App shutting down...")

	ctx.Mongo.Client().Disconnect(context.Background())
	ctx.MSSQL.Close()
}
