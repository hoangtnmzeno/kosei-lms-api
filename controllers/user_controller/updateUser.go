package usercontroller

import (
	"fmt"
	"kosei-lms-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) UpdateUserPhoto(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload models.SignPhoto

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := ctx.ShouldBindUri(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if err := ctx.SaveUploadedFile(payload.Photo, "uploads/"+payload.Photo.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	fmt.Println("currentUser := ", currentUser)
	var updatedUserPhoto models.User
	result := uc.DB.First(&updatedUserPhoto, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}
	now := time.Now()
	userResponse := &models.User{
		Photo:     payload.Photo.Filename,
		CreatedAt: now,
		UpdatedAt: now,
	}
	uc.DB.Model(&updatedUserPhoto).Updates(userResponse)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
