package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type handler struct {
}

func (h *handler) postHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello post")
}
func TestGin(t *testing.T) {
	engine := gin.Default()
	engine.GET("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello get")
	})
	h := &handler{}
	engine.POST("/user", h.postHandler)
	engine.Run(":8080")
}
