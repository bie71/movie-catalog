package app // berisi file untuk menjalankan aplikasi dan merupakan file menginitialisasi aplikasi dan logic dan menghubungkan ke komponen utama seperti service dan endpoint

import (
	"context"
	"log"
	"movie-catalog/config"
	"movie-catalog/lib/auth"
	"movie-catalog/lib/midleware"
	"movie-catalog/lib/pagination"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() {
	cfg := config.Newconfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
		return
	}
	cdfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cdfR2)

	_ = auth.NewJwt(cfg)
	_ = midleware.NewMidleware(cfg)
	_ = pagination.NewPagination()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${latency} ${method} ${path}\n",
	}))

	_ = app.Group("/api")

	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}
		err := app.Listen(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatalf("error starting server : %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
	log.Println("server stopped")
}
