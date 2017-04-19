package models

import (
	"os"

	"github.com/astaxie/beego"
)

func DelApiFile(path, name string) error {
	routerName := path + "/" + ROUTERDIR + "/" + ROUTERFILE
	modelName := path + "/" + MODELSDIR + "/" + name + EXT
	ctrlName := path + "/" + CTRLSDIR + "/" + name + EXT

	form, err := NewForm(name)
	if err != nil {
		return err
	}
	err = os.Remove(modelName)
	if err == nil {
		beego.Info("删除模型文件:", modelName)
	}
	err = os.Remove(ctrlName)
	if err == nil {
		beego.Info("删除控制文件:", ctrlName)
	}
	err = delRouterLine(form, path, routerName)
	if err == nil {
		beego.Info("删除路由引用:", routerName+"->"+name)
	}
	return err
}
