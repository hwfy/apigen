// @APIVersion 1.0.0
// @Title 表单数据 API
// @Description 增删改查表单业务数据
// @Contact yourEmail@gmail.com
package {{.Router}}

import (
	"{{.Name}}/{{.Ctrl}}"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
		{{.LayoutContent}}	
	)
	beego.AddNamespace(ns.Filter("before", func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
		ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,POST,DELETE,PUT")

		//这里可以添加对token的认证
	}))
}
