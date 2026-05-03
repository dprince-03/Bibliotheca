package main

import (
	"bibliotheca/internal/config"
	"bibliotheca/internal/database"
	"bibliotheca/pkg/mysqlclient"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	// validate := validator.New()
	
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// jwtSecret := cfg.JWTSecret

	db, err := mysqlclient.ConnectMySqlClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to MySql client: %v", err)
	}

	defer db.Close()

	if err := database.RunMigration(db, "./migrations"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "root url is healthy",
		})
	})

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "url is working",
		})
	})

	log.Printf("Bibliotheca starting on port %s [%s mode]\n", cfg.ServerPort, cfg.AppEnv)
	server := fmt.Sprintf("0.0.0.0:%s", cfg.ServerPort)
	r.Run(server)
}