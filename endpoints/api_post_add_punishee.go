package endpoints

import (
	"net/http"

	"git.skillissue.dk/models"
	"github.com/gin-gonic/gin"
)

func (shared *SharedContext) ApiPostAddPunishee(c *gin.Context) {
	name := c.Param("name")

	err := shared.Db.AddPunishee(models.Punishee{Name: name, Count: 0})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "success"})
}
