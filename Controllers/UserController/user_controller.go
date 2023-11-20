package usercontroller

import (
	database "finance-tracker-api/Database"
	helpers "finance-tracker-api/Helpers"
	models "finance-tracker-api/Models"
	authrequest "finance-tracker-api/Requests/AuthRequest"
	userresponse "finance-tracker-api/Responses/UserResponse"
	userservice "finance-tracker-api/Services/UserService"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(context *gin.Context) {
	var body authrequest.BodySignUpRequest

	if context.Bind(&body) != nil {
		context.JSON(403, gin.H{
			"status_code": 403,
			"message":     "Failed read body",
		})
		return
	}

	if helpers.JsonIfErr(body.ValidateSignupRequest(), context, 400) {
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if helpers.JsonIfErr(err, context, 500) {
		return
	}

	// Create user
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	if helpers.JsonIfErr(userservice.Store(&user), context, 500) {
		return
	}

	// Response
	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success creating user",
		"data":        userresponse.UserResponseOne(user),
	})
}

func SignIn(context *gin.Context) {
	var body authrequest.BodySignInRequest

	if err := context.Bind(&body); err != nil {
		context.JSON(403, gin.H{
			"status_code": 403,
			"message":     "Failed read body",
		})
		return
	}

	if err := body.ValidateSignInRequest(); err != nil {
		context.JSON(400, gin.H{
			"status_code": 400,
			"message":     err.Error(),
		})
		return
	}

	var user models.User
	database.DB.Find(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(400, gin.H{
			"message": "Failed login, email or password is wrong : user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if helpers.JsonIfErr(err, context, 500) {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if helpers.JsonIfErr(err, context, 500) {
		return
	}

	accessToken := models.PersonalAccessToken{
		UserId: user.ID,
		Token:  tokenString,
	}

	result := database.DB.Create(&accessToken)

	if helpers.JsonIfErr(result.Error, context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"message": "Success login",
		"user":    userresponse.UserResponseOne(user),
		"token":   tokenString,
	})
}

func Profile(context *gin.Context) {
	user := context.MustGet("user").(models.User)

	userModel := models.User{}

	result := database.DB.Where("id = ?", user.ID).First(&userModel)
	if helpers.JsonIfErr(result.Error, context, 500) {
		return
	}

	if result.RowsAffected == 0 {
		// Jika tidak ada data ditemukan
		context.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Success get profile",
		"user":    userresponse.UserResponseOne(userModel),
	})
}
