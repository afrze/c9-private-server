package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/afrze/c9-private-server/service/internal/ctx"
	"github.com/afrze/c9-private-server/service/internal/db"
	"github.com/afrze/c9-private-server/service/internal/models"
)

func RegisterRoutes(r *gin.RouterGroup, app *ctx.Ctx, jwtSecret string, jwtExpiry int) {
	h := &authHandler{app: app, jwtSecret: []byte(jwtSecret), jwtExpiry: time.Duration(jwtExpiry) * time.Minute}

	r.POST("/register", h.register)
	r.POST("/login", h.login)
}

type authHandler struct {
	app       *ctx.Ctx
	jwtSecret []byte
	jwtExpiry time.Duration
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

// POST /api/login
// @Summary      Login
// @Description  Verify credentials from Mongo; returns JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  models.LoginDTO  true  "user + password"
// @Success      200   {object}  map[string]string { token=string }
// @Failure      401   {string}  string  "invalid credentials"
func (h *authHandler) login(c *gin.Context) {
	var dto models.LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := db.FindWebUser(ctx, h.app.Mongo, dto.PAccId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Password hash check
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.PPassword)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, err := h.signJWT(user.AccId, user.AccNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// JWT Helper
func (h *authHandler) signJWT(accId string, accNo int64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   accId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.jwtExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        fmt.Sprint(accNo),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(h.jwtSecret)
}
