package beego

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (uc *UserController) GetUser() {
	uc.Ctx.WriteString("你好！我是one-neo")
}
func (uc *UserController) Shadow() {
	uc.Ctx.WriteString("这是一个影子！")
}
