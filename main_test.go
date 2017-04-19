package main

import (
	"apigen/models"
	"html/template"
	"io/ioutil"
	"testing"
)

const (
	testDir  = "../formDataTest"
	testForm = "tblinformation_reworkapply"

	testDB     = "testdb"
	testHost   = "127.0.0.1"
	testPort   = 0
	testUser   = "testuser"
	testPwd    = "testpwd"
	testDriver = "mssql"
)

func TestGenApi(t *testing.T) {
	dbInfo := models.DBInfo{
		Name:   testDB,
		Host:   testHost,
		Port:   testPort,
		User:   testUser,
		Pwd:    testPwd,
		Driver: testDriver,
	}

	mdata, err := ioutil.ReadFile("./views/orm.tpl")
	if err != nil {
		t.Fatal(err)
	}
	cdata, err := ioutil.ReadFile("./views/ctrl.tpl")
	if err != nil {
		t.Fatal(err)
	}
	modelTpl := template.Must(
		template.New("models").
			Funcs(template.FuncMap{
				"Tags":              models.Tags,
				"AssoTags":          models.AssoTags,
				"TypeConvert":       models.TypeConvert,
				"ExportColumn":      models.ExportColumn,
				"ColumnAndType":     models.ColumnAndType,
				"ColumnWithPostfix": models.ColumnWithPostfix,
			}).Parse(string(mdata)))

	ctrlTpl := template.Must(
		template.New("ctrls").
			Parse(string(cdata)))

	err = models.GenApiFile(dbInfo, modelTpl, ctrlTpl, testDir, testForm)
	if err != nil {
		t.Error(err)
	}
}

func TestDelApi(t *testing.T) {
	err := models.DelApiFile(testDir, testForm)
	if err != nil {
		t.Error(err)
	}
}
