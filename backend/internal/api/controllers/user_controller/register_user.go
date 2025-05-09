package usercontrollers

import (
	"net/http"

	schemas "github.com/HublastX/HubLast-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body schemas.RegisterRequest true "User registration data"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Router /auth/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req schemas.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Register(
		req.Username,
		req.Email,
		req.Password,
		req.Level,
		req.Experience,
		req.Employment,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"level":      user.Level,
			"experience": user.Experience,
			"employment": user.Employment,
		},
	})
}
