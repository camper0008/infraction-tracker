package endpoints

import (
	"net/http"
	"strconv"

	"infraction-tracker/models"

	"github.com/gin-gonic/gin"
)

func (shared *SharedContext) ApiPostUpdatePunishee(c *gin.Context) {
	name := c.Param("name")
	count := c.Param("count")

	intCount, err := strconv.Atoi(count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "msg": "count is not a valid integer"})
	}

	err = shared.Db.UpdatePunishee(models.Punishee{Name: name, Count: intCount})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "success"})
}
