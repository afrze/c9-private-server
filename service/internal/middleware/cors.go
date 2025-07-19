package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS return a gin.HandlerFunc with the project's CORS setting
func CORS() gin.HandlerFunc {
	cfg := cors.Config{
		AllowOrigins: []string{
			"https://c9Dark.netlify.app",
			"http://127.0.0.1:3003",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{}, // none needed
		AllowCredentials: true,       // keep cookies/JWT if needed
		MaxAge:           12 * time.Hour,
	}

	return cors.New(cfg)
}
