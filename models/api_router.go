package models

import (
	"fmt"
	"html/template"
	"log"
)

type router struct {
	name string
}

func (r router) getDir() string {
	return "routers"
}

func (r router) getName() string {
	return "router"
}

func (r router) getChilds() map[string]string {
	return make(map[string]string)
}

func (r router) getTpl() string {
	return "./views/router.tpl"
}

func (r router) getTplFunc() template.FuncMap {
	return nil
}

func (r router) remove(dir string) error {
	path := fmt.Sprintf("%s/%s/%s.go", dir, r.getDir(), r.getName())

	routerTpl, err := parseTpl(r.getName(), r.getTpl(), r.getTplFunc())
	if err != nil {
		return fmt.Errorf("解析模板失败,%s", err)
	}
	var form Form

	if err = form.setControllers(dir); err != nil {
		return err
	}
	if err = form.genFile(routerTpl, path); err != nil {
		return fmt.Errorf("修改路由引用失败:%s", err)
	}
	log.Println("修改路由引用:", path+"->"+r.name, "\n")

	return nil
}
