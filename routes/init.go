package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(c *gin.Context) {
	c.HTML(http.StatusOK, "init", nil)
}
