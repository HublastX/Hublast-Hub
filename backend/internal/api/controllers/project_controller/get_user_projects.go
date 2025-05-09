package projectcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserProjects godoc
// @Summary Get current user's projects
// @Description Retrieves all projects associated with the authenticated user
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns list of user's projects"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /api/user/projects [get]
func (c *ProjectController) GetUserProjects(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projects, err := c.projectService.GetUserProjects(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}
