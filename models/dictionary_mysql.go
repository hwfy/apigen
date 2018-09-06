package models

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type (
	Table struct {
		Schema  string `gorm:"column:TABLE_SCHEMA"`  //数据库
		Name    string `gorm:"column:TABLE_NAME"`    //表名
		Comment string `gorm:"column:TABLE_COMMENT"` //备注

		Usage   Usage    `gorm:"ForeignKey:TABLE_NAME;AssociationForeignKey:TABLE_NAME"`
		Columns []Column `gorm:"ForeignKey:TABLE_NAME;AssociationForeignKey:TABLE_NAME"`
		Childs  []child  `gorm:"ForeignKey:REFERENCED_TABLE_NAME;AssociationForeignKey:TABLE_NAME"`
		Parents []parent `gorm:"ForeignKey:TABLE_NAME;AssociationForeignKey:TABLE_NAME"`
	}

	Column struct {
		Table                  string `gorm:"column:TABLE_NAME" `               //表名
		Name                   string `gorm:"column:COLUMN_NAME" `              //列名
		Type                   string `gorm:"column:DATA_TYPE" `                //数据类型
		CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH" ` //字符长度
		NumericPrecision       int    `gorm:"column:NUMERIC_PRECISION" `        //数字长度
		NumericScale           int    `gorm:"column:NUMERIC_SCALE" `            //小数位数
		Comment                string `gorm:"column:COLUMN_COMMENT" `           //备注
	}

	child struct {
		Parent string `gorm:"column:REFERENCED_TABLE_NAME" `  //表名
		Name   string `gorm:"column:TABLE_NAME" `             //子表
		Self   string `gorm:"column:COLUMN_NAME" `            //外键
		Link   string `gorm:"column:REFERENCED_COLUMN_NAME" ` //主键
	}

	parent struct {
		Child string `gorm:"column:TABLE_NAME" `             //子表
		Name  string `gorm:"column:REFERENCED_TABLE_NAME" `  //表名
		Link  string `gorm:"column:REFERENCED_COLUMN_NAME" ` //主键
		Self  string `gorm:"column:COLUMN_NAME" `            //外键
	}

	Usage struct {
		Table  string `gorm:"column:TABLE_NAME" `      //表名
		Column string `gorm:"column:COLUMN_NAME" `     //列名
		Name   string `gorm:"column:CONSTRAINT_NAME" ` //是否主键

	}
)

// TableName 将Table映射为information_schema.tables
func (table Table) TableName() string {
	return "information_schema.tables"
}

// TableName 将Column映射为information_schema.columns
func (table Column) TableName() string {
	return "information_schema.columns"
}

// TableName 将child映射为information_schema.KEY_COLUMN_USAGE
func (table child) TableName() string {
	return "information_schema.KEY_COLUMN_USAGE"
}

// TableName 将parent映射为information_schema.KEY_COLUMN_USAGE
func (table parent) TableName() string {
	return "information_schema.KEY_COLUMN_USAGE"
}

// TableName 将Usage映射为information_schema.KEY_COLUMN_USAGE
func (table Usage) TableName() string {
	return "information_schema.KEY_COLUMN_USAGE"
}

// getDataDictionaryFromMysql 从mysql获取数据字典
func getDataDictionaryFromMysql(db *gorm.DB, schema, name string) (*Form, error) {
	var table Table

	query := db.
		Preload("Columns").
		Preload("Childs").
		Preload("Parents").
		Where("table_schema=? and table_name=?", schema, name).
		Find(&table)

	if query.Error != nil {
		return nil, fmt.Errorf("从mysql获取数据字典失败,%s", query.Error)
	}
	return table.update(), nil
}

func (table Table) update() *Form {
	return &Form{
		Name:    table.Name,
		Base:    table.Schema,
		Chinese: table.Comment,
		Fields:  table.getFields(),
		Parents: table.getParents(),
		Childs:  table.getChilds(),
	}
}

func (table Table) getFields() []Field {
	var fields []Field

	for _, column := range table.Columns {
		var field Field

		field.Name = column.Name
		field.Type = column.Type
		field.Chinese = column.Comment

		if column.CharacterMaximumLength != 0 {
			field.Length = strconv.Itoa(column.CharacterMaximumLength)
		} else {
			field.Length = strconv.Itoa(column.NumericPrecision)
		}
		fields = append(fields, field)
	}
	return fields
}

func (table Table) getParents() []Parent {
	var parents []Parent

	for _, p := range table.Parents {
		parent := Parent{
			Name: p.Name,
			Link: p.Link,
			Self: p.Self,
		}
		parents = append(parents, parent)
	}
	return parents
}

func (table Table) getChilds() []Child {
	var childs []Child

	for _, c := range table.Childs {
		child := Child{
			Name: c.Name,
			Self: c.Self,
			Link: c.Link,
		}
		childs = append(childs, child)
	}
	return childs
}
