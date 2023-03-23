package i18n

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const localizerKey = "Localizer"

// Localizer options
type Options struct {
	DefaultLang  string // default language
	SupportLangs string // list of supported languages ​​(must include default language)
	FilePath     string // multilingual file directory
}

// Localizer middleware
func Localizer(opt *Options) gin.HandlerFunc {
	defaultStr := strings.TrimSpace(opt.DefaultLang)
	supportStr := strings.TrimSpace(opt.SupportLangs)
	filePathStr := strings.TrimSpace(opt.FilePath)
	if len(defaultStr) == 0 || len(supportStr) == 0 {
		panic("bad defaultLang or supportLang")
	}

	l := newLocalize(defaultStr, supportStr, filePathStr)

	return func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		localizer := l.userLocalize(acceptLang)
		c.Set(localizerKey, localizer)
		c.Next()
	}
}

// Get localizer message
func Msg(c *gin.Context, tag string, args ...interface{}) string {
	localizer := c.MustGet(localizerKey).(*userLocalize)
	return localizer.getMsg(tag, args...)
}
