package handlers

import (
	"github.com/airdb/passport/model/bo"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	c.Redirect(307, bo.GetRewriteURI())
}
