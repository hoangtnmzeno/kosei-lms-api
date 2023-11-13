package auth_controller

import (
	"kosei-lms-api/initializers"
	"kosei-lms-api/models"
	"kosei-lms-api/otp"
	"kosei-lms-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ac.Auth.GetUserByName(payload.Name)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "name already exists"})
		return
	}
	// emails := make([]string, 0)
	// emails = append(emails, user.Email)
	otpGen, err := otp.GenerateOTP(6)
	if err != nil {
		return
	}
	// go func() {
	// 	mail.SendMail(emails, otpGen)
	// }()
	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token, "otp": otpGen})
}
