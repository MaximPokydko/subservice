package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"subservice/internal/config"
	"subservice/internal/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("database connected")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":" + cfg.Port)

	_ = db
}
