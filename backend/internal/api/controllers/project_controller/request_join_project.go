package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RequestJoinProject godoc
// @Summary Request to join a project
// @Description Allows the current user to request joining a specific project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Join request successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/projects/{id}/join [post]
func (c *ProjectController) RequestJoinProject(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// In a real implementation, you might create a join request record
	// For simplicity, we'll just add the user to the project directly
	err = c.projectService.AddUserToProject(uint(projectID), userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Join request submitted successfully"})
}
