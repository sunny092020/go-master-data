package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunny092020/go-master-data/datatransfers"
)

func GETPing(c *gin.Context) {
	c.JSON(http.StatusOK, datatransfers.Response{Data: "pong"})
}
