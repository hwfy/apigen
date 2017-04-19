package models

import (
	"encoding/json"

	"github.com/hwfy/redis"
)

const (
	//路由文件模板： router.tpl
	ROUTER = "router.tpl"

	//路由文件目录： routers
	ROUTERDIR = "routers"

	//路由文件名称： router.go
	ROUTERFILE = "router.go"

	//命名空间文件模板： namespace.tpl
	NAMESPACE = "namespace.tpl"

	//模型文件目录: models
	MODELSDIR = "models"

	//控制器文件目录: controllers
	CTRLSDIR = "controllers"

	//文件扩展名:  .go
	EXT = ".go"
)

type (
	Form struct {
		Name   string
		Flag   string
		Base   string
		Lang   Version
		Fields []Field
		Childs []Child
		DBInfo
		Package
	}
	Version struct {
		Cn string
		En string
		Tw string
	}
	Field struct {
		Name string
		Flag string
		Type string
		Lang Version
	}
	Child struct {
		Name string
		Self string
		Link string
	}
	DBInfo struct {
		Name   string
		Host   string
		Port   int
		User   string
		Pwd    string
		Driver string
	}
	Package struct {
		Ctrl  string
		Model string
	}
)

// NewForm 根据名称获取数据字典
func NewForm(name string) (form *Form, err error) {
	client, err := redis.NewClient("form")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bytes, err := client.Get(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &form)
	if err != nil {
		return nil, err
	}
	return form, nil
}

// GetAllChilds 获取所有关联的子表，tables必须用make()初始化
func (form *Form) GetAllChilds(tables map[string]string) {
	for _, child := range form.Childs {
		if form.Name == child.Name {
			continue
		}
		tables[child.Name] = child.Self

		table, err := NewForm(child.Name)
		if err != nil {
			return
		}
		if len(table.Childs) != 0 {
			table.GetAllChilds(tables)
		}
	}
}
