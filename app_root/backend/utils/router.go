package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sunny092020/go-master-data/constants"
	"github.com/sunny092020/go-master-data/datatransfers"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: "user not authenticated"})
	}
}
