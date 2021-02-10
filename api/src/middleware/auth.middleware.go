package middleware

import (
	"net/http"

	"github.com/cheesecat47/webpractice/api/helper"
	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware func is middleware for verifying authorization
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqHeader := c.Request.Header.Get("Authorization")
		if reqHeader == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}
		clientToken := helper.ExtractToken(reqHeader)
		if clientToken == "" {
			c.JSON(400, "Invalid format of authorization token")
			c.Abort()
			return
		}
		//JwtInfo:ACCESS_TOKEN is temporal
		jwtInfo := helper.JwtInfo{AccessKey: "ACCESS_TOKEN", RefreshKey: "REFRESH_TOKEN", Issuer: "atg0831"}
		claims, err := jwtInfo.ValidateToken(clientToken)

		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		//c.Get("userInfo")으로 받아야함...
		c.JSON(http.StatusOK, "Authorization success\n")
		c.Set("userInfo", claims)
		c.Next()
	}
}
