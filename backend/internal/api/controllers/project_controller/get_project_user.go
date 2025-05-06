package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProjectUsers godoc
// @Summary Get users in a project
// @Description Retrieves all users associated with a specific project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{} "Returns project details with users"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Project not found"
// @Security ApiKeyAuth
// @Router /api/projects/{id}/users [get]
func (c *ProjectController) GetProjectUsers(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := c.projectService.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var users []gin.H
	for _, user := range project.Users {
		users = append(users, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"project_id":    project.ID,
		"project_title": project.Title,
		"responsible_user": gin.H{
			"id":       project.ResponsibleUser.ID,
			"username": project.ResponsibleUser.Username,
		},
		"users": users,
	})
}
