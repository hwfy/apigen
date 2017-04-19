package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/hwfy/file"
)

// genRouterFile 添加路由引用
func genRouterFile(form *Form, path, name string) error {
	err := file.AppendContent("./views/"+NAMESPACE, getText(form))
	if err != nil {
		return err
	}
	return genRouter(path, name)
}

// delRouterLine 删除路由引用
func delRouterLine(form *Form, path, name string) error {
	err := file.RemoveLine("./views/"+NAMESPACE, getText(form))
	if err != nil {
		return err
	}
	return genRouter(path, name)
}

func getText(form *Form) string {
	return fmt.Sprintf("beego.NSNamespace(%q ,\n	beego.NSInclude(\n		&controllers.%sController{},\n)),\n", "/"+form.Name, form.Flag)
}

func genRouter(path, name string) error {
	var tpl beego.Controller
	tpl.Data = make(map[interface{}]interface{})

	if strings.HasPrefix(path, "../") {
		path = strings.TrimPrefix(path, "../")
	}
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	tpl.Layout = ROUTER
	tpl.TplName = NAMESPACE
	tpl.Data["Name"] = path
	tpl.Data["Ctrl"] = CTRLSDIR
	tpl.Data["Router"] = ROUTERDIR

	fileBuf, err := tpl.RenderBytes()
	if err != nil {
		return err
	}
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(fileBuf)

	return err
}
