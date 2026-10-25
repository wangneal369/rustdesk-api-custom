package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
	"github.com/lejianwen/rustdesk-api/v2/http/response"
	"net/http"
	"strings"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginLimiter := global.LoginLimiter
		clientIp := c.ClientIP()
		banned, _ := loginLimiter.CheckSecurityStatus(clientIp)
		if banned {
			// custom: 管理端沿用管理端信封；客户端路由使用客户端原生错误结构 {"error": ...}，协议不变
			if strings.HasPrefix(c.Request.URL.Path, "/api/console") {
				response.Fail(c, http.StatusLocked, response.TranslateMsg(c, "Banned"))
			} else {
				c.JSON(http.StatusLocked, gin.H{"error": response.TranslateMsg(c, "Banned")})
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
