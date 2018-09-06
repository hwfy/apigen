package models

import (
	"html/template"
	"log"
	"os"
)

type controller struct {
	name string
}

func (c controller) getDir() string {
	return "controllers"
}

func (c controller) getName() string {
	return c.name
}

func (c controller) getChilds() map[string]string {
	return make(map[string]string)
}

func (c controller) getTpl() string {
	return "./views/ctrl.tpl"
}

func (c controller) getTplFunc() template.FuncMap {
	return template.FuncMap{
		"ExportStructTag": ExportStructTag,
	}
}

func (c controller) remove(dir string) error {
	name := dir + "/" + c.getDir() + "/" + c.name + ".go"

	if err := os.Remove(name); err == nil {
		log.Println("删除控制文件:", name)
	}
	return nil
}
