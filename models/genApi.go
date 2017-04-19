package models

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
	"github.com/hwfy/file"
)

// GenApiFiles 从文件列表批量生成API
func GenApiFiles(db DBInfo, modelTpl, ctrlTpl *template.Template, path, forms string) error {
	file, err := os.Open(forms)
	if err != nil {
		return err
	}
	defer file.Close()

	buff := bufio.NewReader(file)
	for {
		line, readErr := buff.ReadString('\n')
		form := strings.TrimSuffix(line, "\n")

		err = GenApiFile(db, modelTpl, ctrlTpl, path, form)
		if err != nil {
			return err
		}
		if readErr != nil {
			break
		}
	}
	return nil
}

// GenApiFile 生成模型、控制、路由文件
func GenApiFile(db DBInfo, modelTpl, ctrlTpl *template.Template, path, name string) error {
	form, err := NewForm(name)
	if err != nil {
		return err
	}
	form.DBInfo = db
	form.Package.Ctrl = CTRLSDIR
	form.Package.Model = MODELSDIR

	childs := make(map[string]string)
	form.GetAllChilds(childs)

	modelDir := path + "/" + MODELSDIR + "/"
	modelName := modelDir + name + EXT

	ctrlDir := path + "/" + CTRLSDIR + "/"
	ctrlName := ctrlDir + name + EXT

	routerDir := path + "/" + ROUTERDIR + "/"
	routerName := routerDir + ROUTERFILE

	file.MkDir(path, modelDir, ctrlDir, routerDir)

	err = genFile(form, modelTpl, modelName)
	if err != nil {
		return fmt.Errorf("生成模型文件失败:%s", err)
	}
	beego.Info("生成模型文件:", modelName)

	for child, _ := range childs {
		table, err := NewForm(child)
		if err != nil {
			return err
		}
		table.DBInfo = db
		table.Package.Ctrl = CTRLSDIR
		table.Package.Model = MODELSDIR

		modelName := modelDir + child + EXT

		err = genFile(table, modelTpl, modelName)
		if err != nil {
			return fmt.Errorf("生成子模型文件失败:%s", err)
		}
		beego.Info("生成子模型文件:", modelName)
	}

	err = genFile(form, ctrlTpl, ctrlName)
	if err != nil {
		return fmt.Errorf("生成控制文件失败:%s", err)
	}
	beego.Info("生成控制文件:", ctrlName)

	err = genRouterFile(form, path, routerName)
	if err != nil {
		return fmt.Errorf("添加路由引用失败:%s", err)
	}
	beego.Info("添加路由引用:", routerName+"->"+name, "\n")

	return nil
}

func genFile(form *Form, render *template.Template, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = render.Execute(file, form)
	if err != nil {
		return err
	}
	return exec.Command("goimports", "-w", fileName).Run()
}
