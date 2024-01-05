package api

import (
	"main/web/server/internal/service"
	"main/web/server/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetTree get tree.
func GetTree(c *gin.Context) {
	directories, err := service.NewHomeService().GetTree()
	response.Response(c, directories, err)
}
