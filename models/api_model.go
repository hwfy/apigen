package models

import (
	"html/template"
	"log"
	"os"
)

type model struct {
	name   string
	childs map[string]string //子表名称
}

func (m model) getDir() string {
	return "models"
}

func (m model) getName() string {
	return m.name
}

func (m model) getChilds() map[string]string {
	return m.childs
}

func (m model) getTpl() string {
	return "./views/orm.tpl"
}

func (m model) getTplFunc() template.FuncMap {
	return template.FuncMap{
		"Tags":              Tags,
		"ToLower":           ToLower,
		"AssoTags":          AssoTags,
		"TypeConvert":       TypeConvert,
		"ExportStructTag":   ExportStructTag,
		"ColumnAndType":     ColumnAndType,
		"ColumnWithPostfix": ColumnWithPostfix,
	}
}

func (m model) remove(dir string) error {
	name := dir + "/" + m.getDir() + "/" + m.name + ".go"

	if err := os.Remove(name); err == nil {
		log.Println("删除模型文件:", name)
	}
	return nil
}
