package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type (
	Form struct {
		ID      string `gorm:"column:ID"`
		Name    string `gorm:"column:Table_Name"`
		PKey    string `gorm:"column:TablePKField"`
		Base    string `gorm:"column:Target_DBName"`
		Taiwan  string `gorm:"column:Table_Taiwan_Name"`
		Chinese string `gorm:"column:Table_Chinese_Name"`
		English string `gorm:"column:Table_English_Name"`

		Fields  []Field  `gorm:"ForeignKey:Table_ID;AssociationForeignKey:ID"`
		Childs  []Child  `gorm:"ForeignKey:Relation_Target_Table;AssociationForeignKey:Table_Name"`
		Parents []Parent `gorm:"ForeignKey:Table_Name;AssociationForeignKey:Table_Name"`

		DataSource
		Controllers map[string]string
	}

	Field struct {
		TableID string `gorm:"column:Table_ID"`
		Name    string `gorm:"column:Field_Name"`
		Taiwan  string `gorm:"column:Field_Taiwan_Name"`
		Chinese string `gorm:"column:Field_Chinese_Name"`
		English string `gorm:"column:Field_English_Name"`
		Type    string `gorm:"column:Field_Type"`
		Length  string `gorm:"column:Field_Length"`
	}

	Child struct {
		Target string `gorm:"column:Relation_Target_Table"`
		Name   string `gorm:"column:Table_Name"`
		Self   string `gorm:"column:Relation_Self_Field"`
		Link   string `gorm:"column:Relation_Target_Field"`
	}

	Parent struct {
		Table string `gorm:"column:Table_Name"`
		Name  string `gorm:"column:Relation_Target_Table"`
		Link  string `gorm:"column:Relation_Target_Field"`
		Self  string `gorm:"column:Relation_Self_Field"`
	}
)

// TableName 将Form映射为tbsf_table_directory
func (table Form) TableName() string {
	return "tbsf_table_directory"
}

// TableName 将Field映射为tbsf_table_field_directory
func (table Field) TableName() string {
	return "tbsf_table_field_directory"
}

// TableName 将Child映射为tbsf_form_tableinfo
func (table Child) TableName() string {
	return "tbsf_form_tableinfo"
}

// TableName 将Parent映射为tbsf_form_tableinfo
func (table Parent) TableName() string {
	return "tbsf_form_tableinfo"
}

// getDataDictionaryFromSqlserver 从sqlserver获取数据字典
func getDataDictionaryFromSqlserver(db *gorm.DB, name string) (*Form, error) {
	var form Form

	query := db.
		Preload("Fields").
		Preload("Childs").
		Preload("Parents").
		Where("Table_Name=?", name).
		Find(&form)

	if query.Error != nil {
		return nil, fmt.Errorf("从mssql获取数据字典失败, %s", query.Error)
	}
	return &form, nil
}
