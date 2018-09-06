package models

import (
	"fmt"
	"html/template"

	"github.com/astaxie/beego"
)

type DataSource struct {
	Name   string
	Host   string
	Port   int
	User   string
	Pwd    string
	Driver string
}

var ds DataSource

func init() {
	ds.Name = beego.AppConfig.String("dictionary::db")

	ds.Host = beego.AppConfig.String("dictionary::host")

	ds.Port, _ = beego.AppConfig.Int("dictionary::port")

	ds.User = beego.AppConfig.String("dictionary::user")

	ds.Pwd = beego.AppConfig.String("dictionary::pwd")

	ds.Driver = beego.AppConfig.String("dictionary::driver")
}

// newDataSource 数据源驱动和链接字符
func newDataSource() (string, string) {
	return ds.Driver, ds.connectionString()
}

func (d DataSource) getDir() string {
	return "models"
}

func (d DataSource) getName() string {
	return d.Name
}

func (d DataSource) getChilds() map[string]string {
	return make(map[string]string)
}

func (d DataSource) getTpl() string {
	return "./views/datasource.tpl"
}

func (d DataSource) getTplFunc() template.FuncMap {
	return template.FuncMap{
		"ToLower": ToLower,
	}
}

func (d DataSource) remove(dir string) error {
	return nil
}

func (d DataSource) connectionString() string {
	var connString string

	switch d.Driver {
	case "mssql":
		connString = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s", d.Host, d.Name, d.User, d.Pwd)
		if d.Port > 0 {
			connString = fmt.Sprintf("%s;port=%d", connString, d.Port)
		}
		connString = connString + ";encrypt=disable"

	case "mysql":
		connString = fmt.Sprintf("%s:%s@tcp(%s", d.User, d.Pwd, d.Host)
		if d.Port > 0 {
			connString = fmt.Sprintf("%s:%d)/%s", connString, d.Port, d.Name)
		} else {
			connString = fmt.Sprintf("%s)/%s", connString, d.Name)
		}
		connString = fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", connString)
	}
	return connString
}
