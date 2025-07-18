package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/afrze/c9-private-server/service/internal/ctx"
	"github.com/afrze/c9-private-server/service/internal/db"
	"github.com/afrze/c9-private-server/service/internal/models"
)

func RegisterRoutes(r *gin.RouterGroup, app *ctx.Ctx) {
	h := &authHandler{app: app}

	r.POST("/register", h.register)
	// r.POST("/login", h.login)
}

type authHandler struct {
	app *ctx.Ctx
}

// POST 		/api/register
// @Summary 	Register
// @Description Creates an account in MSSQL and Mongo
// @Tags 		auth
// @Accept 		json
// @Produce 	json
// @param 		body body model.RegisterDTO true "user+password"
// @Success 	200 {string} string "Account Created..."
// @Failure 	409 {string} string "User Already exists..."
// @Router 		/api/register [post]
func (h *authHandler) register(c *gin.Context) {
	var dto models.RegisterDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Call Store Procedure
	dup, accNo, err := db.CreateAccount(ctx, h.app.MSSQL, dto.PAccId, dto.PPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if dup {
		c.JSON(http.StatusConflict, gin.H{"error": "User exists"})
		return
	}

	// Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(dto.PPassword), bcrypt.DefaultCost)

	// Insert into Mongo
	err = db.InsertWebUser(ctx, h.app.Mongo, db.WebUser{
		AccNo:        accNo,
		AccId:        dto.PAccId,
		PasswordHash: string(hash),
		Created:      time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Account created...")
}
