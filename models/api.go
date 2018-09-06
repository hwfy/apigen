package models

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type component interface {
	//获取组件目录
	getDir() string
	//获取组件名称
	getName() string
	//获取子组件名称
	getChilds() map[string]string
	//获取组件模板
	getTpl() string
	//获取组件模板函数
	getTplFunc() template.FuncMap
	//根据项目名dir移除api组件
	remove(dir string) error
	//根据项目名dir生成api组件
	//generator(dir string, data *Form) error
}

func newComponents(name, base string, childs map[string]string) []component {
	return []component{
		DataSource{Name: base},

		model{name: name, childs: childs},

		controller{name: name},
		//生成控制器后添加路由引用
		router{name: name},
	}
}

// parseTpl 解析模板文件
func parseTpl(name, path string, fun template.FuncMap) (*template.Template, error) {
	tpl := template.New(name).Funcs(fun)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return tpl.Parse(string(data))
}

// genFile 创建文件并添加渲染内容
func (form *Form) genFile(render *template.Template, fileName string) error {
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

// GenApiFiles 从表单配置文件生成API
func (ds DataSource) GenApiFiles(path, forms string) error {
	file, err := os.Open(forms)
	if err != nil {
		return err
	}
	defer file.Close()

	buff := bufio.NewReader(file)
	for {
		line, readErr := buff.ReadString('\n')
		form := strings.TrimSuffix(line, "\n")

		err = ds.GenApiFile(path, form)
		if err != nil {
			return err
		}
		if readErr != nil {
			break
		}
	}
	return nil
}

// genApiFile 从命令行生成API
func (ds DataSource) GenApiFile(path, name string) error {
	childs := make(map[string]string)

	form, err := getDatadictionaryByName(name)
	if err != nil {
		return err
	}
	form.DataSource = ds   //模型文件中需要数据源名称
	form.getChilds(childs) //子模型文件

	path = strings.TrimSuffix(path, "/")
	//创建主目录
	if _, err := os.Stat(path); err != nil {
		os.Mkdir(path, os.ModePerm)
	}
	for _, comp := range newComponents(name, form.Base, childs) {
		//生成控制器以后设置
		form.setControllers(path)
		//定义名称
		dir := path + "/" + comp.getDir() + "/"
		fileName := dir + comp.getName() + ".go"
		//创建目录
		if _, err := os.Stat(dir); err != nil {
			os.Mkdir(dir, os.ModePerm)
		}
		//解析模板
		tpl, err := parseTpl(comp.getName(), comp.getTpl(), comp.getTplFunc())
		if err != nil {
			return fmt.Errorf("解析模板失败,%s", err)
		}
		//生成文件
		if err = form.genFile(tpl, fileName); err != nil {
			return fmt.Errorf("生成文件失败:%s", err)
		}
		log.Println("生成文件:", fileName)
		//生成子文件
		for name, _ := range comp.getChilds() {
			err = ds.genChildFiles(dir, name, tpl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ds DataSource) genChildFiles(dir, name string, tpl *template.Template) error {
	path := dir + name + ".go"
	//如果文件存在则继续其他子表
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	form, err := getDatadictionaryByName(name)
	if err != nil {
		return err
	}
	//子模型文件和子模型数据源文件都需要
	form.DataSource = ds
	//生成子模型数据源
	if err = form.genDataSource(dir); err != nil {
		return err
	}
	if err = form.genFile(tpl, path); err != nil {
		return fmt.Errorf("生成子文件失败:%s", err)
	}
	log.Println("生成子文件:", path)

	return nil
}

func (form Form) genDataSource(dir string) error {
	name := dir + form.Base + ".go"
	//如果数据源文件存在则不生成
	_, err := os.Stat(name)
	if err == nil {
		return nil
	}
	ds := form.DataSource
	//解析数据源模板
	tpl, err := parseTpl(ds.getName(), ds.getTpl(), ds.getTplFunc())
	if err != nil {
		return err
	}
	if err = form.genFile(tpl, name); err != nil {
		return fmt.Errorf("生成数据文件失败:%s", err)
	}
	log.Println("生成数据文件:", name)

	return nil
}

// DelApiFile 从命令行删除API
func DelApiFile(path, name string) error {
	for _, comp := range newComponents(name, "", make(map[string]string)) {
		if err := comp.remove(path); err != nil {
			return err
		}
	}
	return nil
}
