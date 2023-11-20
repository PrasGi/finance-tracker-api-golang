package helpers

import "github.com/gin-gonic/gin"

func JsonIfErr(err error, context *gin.Context, status int) bool {
	if err != nil {
		context.JSON(500, gin.H{
			"status_code": status,
			"message":     err.Error(),
		})
		return true
	}

	return false
}
