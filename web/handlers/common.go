package handlers

import "github.com/gin-gonic/gin"

func renderError(c *gin.Context, template string, message string, statusCode int, extraData ...gin.H) {
	data := gin.H{
		"Title": "Error",
		"Error": message,
	}

	// Merge any extra data provided
	if len(extraData) > 0 {
		for key, value := range extraData[0] {
			data[key] = value
		}
	}

	c.HTML(statusCode, template, data)
}
