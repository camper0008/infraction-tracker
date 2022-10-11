package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (shared *SharedContext) ApiPostRemovePunishee(c *gin.Context) {
	name := c.Param("name")

	err := shared.Db.RemovePunishee(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "success"})
}
