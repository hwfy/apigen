package main

import (
	"apigen/models"
	"flag"
	"html/template"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/hwfy/file"
)

var (
	db   = flag.String("db", "", "the name of db")
	host = flag.String("host", "", "the host of db")
	port = flag.Int("port", 0, "the port of db host")

	user   = flag.String("user", "", "the user name of db")
	pwd    = flag.String("pwd", "", "the password of db")
	driver = flag.String("driver", "", "the driver of db")

	form  = flag.String("form", "", "the name of form")
	forms = flag.String("forms", "", "the file of forms")

	path = flag.String("path", "", "the path of folder for model and controller files")

	del = flag.String("del", "", "delete the form name")
)

func init() {
	logDir := "../logs/apiGen/"

	file.MkDir(logDir)

	beego.SetLogger("file", `{
		"filename":"`+logDir+`app.log",
		"level":7,
		"maxlines":0,
		"maxsize":0,
		"daily":true,
		"maxdays":10 
	}`)

	switch "" {
	case *db:
		*db = beego.AppConfig.String("db")
		fallthrough
	case *host:
		*host = beego.AppConfig.String("host")
		fallthrough
	case *user:
		*user = beego.AppConfig.String("user")
		fallthrough
	case *pwd:
		*pwd = beego.AppConfig.String("pwd")
		fallthrough
	case *driver:
		*driver = beego.AppConfig.String("driver")
		fallthrough
	case *path:
		*path = beego.AppConfig.String("path")
		fallthrough
	case *forms:
		*forms = beego.AppConfig.String("forms")
	}
	if *port == 0 {
		*port, _ = beego.AppConfig.Int("port")
	}
}

func main() {
	flag.Parse()

	dbInfo := models.DBInfo{
		Name:   *db,
		Host:   *host,
		Port:   *port,
		User:   *user,
		Pwd:    *pwd,
		Driver: *driver,
	}

	mdata, err := ioutil.ReadFile("./views/orm.tpl")
	if err != nil {
		beego.Error(err)
		return
	}
	cdata, err := ioutil.ReadFile("./views/ctrl.tpl")
	if err != nil {
		beego.Error(err)
		return
	}

	mTpl := template.Must(
		template.New("models").
			Funcs(template.FuncMap{
				"Tags":              models.Tags,
				"AssoTags":          models.AssoTags,
				"TypeConvert":       models.TypeConvert,
				"ExportColumn":      models.ExportColumn,
				"ColumnAndType":     models.ColumnAndType,
				"ColumnWithPostfix": models.ColumnWithPostfix,
			}).Parse(string(mdata)))

	cTpl := template.Must(
		template.New("ctrls").
			Parse(string(cdata)))

	if *del != "" {
		err = models.DelApiFile(*path, *del)
		if err != nil {
			beego.Error(err)
		}
	} else {
		if *form != "" {
			err = models.GenApiFile(dbInfo, mTpl, cTpl, *path, *form)
			if err != nil {
				beego.Error(err)
			}
		} else {
			err = models.GenApiFiles(dbInfo, mTpl, cTpl, *path, *forms)
			if err != nil {
				beego.Error(err)
			}
		}
	}
}
