package main

import (
	"cy/CloudERP/apiGen/models"

	"testing"
)

const (
	testDir  = "../formDataTest"
	testForm = "tblinformation_reworkapply"
)

func newDataSource() models.DataSource {
	return models.DataSource{
		Name:   "testdb",
		Host:   "127.0.0.1",
		Port:   0,
		User:   "testuser",
		Pwd:    "testpwd",
		Driver: "mssql",
	}
}

func TestGenApi(t *testing.T) {
	datasource := newDataSource()

	err := datasource.GenApiFile(testDir, testForm)
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
