package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (shared *SharedContext) ApiGetList(c *gin.Context) {
	list, err := shared.Db.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "success", "list": list})
}
