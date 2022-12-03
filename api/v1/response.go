package v1

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func SuccessResponse() gin.H {
	return gin.H{"status": "success"}
}
