package i18n

import (
	"github.com/gin-gonic/gin"
)

// GinLocalizer middleware
func GinLocalizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		localizer := NewUserLocalize(acceptLang)
		c.Set("Localizer", localizer)
		c.Next()
	}
}
