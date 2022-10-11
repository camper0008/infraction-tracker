package endpoints

import (
	"net/http"

	"git.skillissue.dk/env"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	ApiKey string `header:"X-API-KEY"`
}

func (shared *SharedContext) Auth(c *gin.Context) {
	h := authHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "msg": "invalid api key"})
	}
	requiredKey := env.ApiKey()
	if h.ApiKey != requiredKey {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "msg": "invalid api key"})
		return
	}

	c.Next()
}
