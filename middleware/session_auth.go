package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/public"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo, ok := session.Get(public.AdminSessionInfoKey).(string); !ok || adminInfo == "" {
			fmt.Println(ok)
			fmt.Println("$$$$$$$$")
			fmt.Println(adminInfo)
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
