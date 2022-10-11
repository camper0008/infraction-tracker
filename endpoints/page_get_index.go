package endpoints

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (sharedContext *SharedContext) PageGetIndex(c *gin.Context) {
	list, err := sharedContext.Db.List()
	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": err.Error(),
		})
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list,
	})
}
