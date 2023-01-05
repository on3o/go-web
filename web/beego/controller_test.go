package beego

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"testing"
)

func TestUserController(t *testing.T) {
	go func() {
		suc := &UserController{}
		//web.Router("/user", suc, fmt.Sprintf("%s:%s", http.MethodGet, "Shadow"))
		s := web.NewHttpSever()
		s.Router("/user", suc, fmt.Sprintf("%s:%s", http.MethodGet, "Shadow"))
		s.Run(":9999")
	}()
	web.BConfig.CopyRequestBody = true
	uc := &UserController{}
	web.Router("/user", uc, fmt.Sprintf("%s:%s", http.MethodGet, "GetUser"))
	web.Run(":8080")
}
