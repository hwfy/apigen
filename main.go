package main

import (
	"cy/CloudERP/apiGen/models"

	"flag"
	"log"

	"github.com/astaxie/beego"
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

func newDatasource() models.DataSource {
	if *db == "" {
		log.Fatalf("数据库名称为空!")
	}
	return models.DataSource{
		Name:   *db,
		Host:   *host,
		Port:   *port,
		User:   *user,
		Pwd:    *pwd,
		Driver: *driver,
	}
}

func main() {
	flag.Parse()

	datasource := newDatasource()

	if *del == "" {
		if *form != "" {
			if err := datasource.GenApiFile(*path, *form); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := datasource.GenApiFiles(*path, *forms); err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		if err := models.DelApiFile(*path, *del); err != nil {
			log.Fatalln(err)
		}
	}
}
