// @APIVersion 1.0.0
// @Title 表单数据 API
// @Description 增删改查表单业务数据
// @Contact luckyfanyang@gmail.com
package routers

//数据库引擎需要放在首行导入
import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"


	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
	{{range $key,$value:=.Controllers}}
		beego.NSNamespace("/{{$key}}",
			beego.NSInclude(
				&controllers.{{$value}}Controller{},
			)),
	{{end}}	
	)
	beego.AddNamespace(ns.Filter("before", func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
		ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,PUT")

		//这里可以添加token认证
		//_, err := authMiddle.GetUser(ctx.Request)
		//if err != nil {
		//	ctx.Abort(401, err.Error())
		//	return
		//}
	}))
}
