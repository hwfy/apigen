package models

import (
	"fmt"
	"html/template"
	"strings"
)

// HavePk 判断是否包含主键
func (m *Form) HavePk() bool {
	return len(m.PkColumnsSchema()) > 0
}

// PkColumns 获取主键名称
func (m *Form) PkColumns() []string {
	pkColumnsSchema := m.PkColumnsSchema()
	result := make([]string, 0, len(pkColumnsSchema))
	for _, t := range pkColumnsSchema {
		result = append(result, t.Name)
	}
	return result
}

// PkColumnsSchema 获取主键字段
func (m *Form) PkColumnsSchema() []Field {
	result := make([]Field, 0, len(m.Fields))
	for _, t := range m.Fields {
		if strings.ToUpper(t.Name) == "ID" {
			result = append(result, t)
		}
	}
	return result
}

// NoPkColumns 获取所有非主键名
func (m *Form) NoPkColumns() []string {
	noPkColumnsSchema := m.NoPkColumnsSchema()
	result := make([]string, 0, len(noPkColumnsSchema))
	for _, t := range noPkColumnsSchema {
		result = append(result, t.Name)
	}
	return result
}

// NoPkColumnsSchema 获取所有非主键字段
func (m *Form) NoPkColumnsSchema() []Field {
	result := make([]Field, 0, len(m.Fields))
	for _, t := range m.Fields {
		if strings.ToUpper(t.Name) != "ID" {
			result = append(result, t)
		}
	}
	return result
}

// Conn 根据驱动名返回不同的数据库链接字符串,目前仅支持mssql、 mysql、 postgres
func (m *Form) Conn() (connString string) {
	switch m.Driver {
	case "mssql":
		connString = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s", m.Host, m.DBInfo.Name, m.User, m.Pwd)
		if m.Port > 0 {
			connString = fmt.Sprintf("%s;port=%d", connString, m.Port)
		}
		connString = connString + ";encrypt=disable"
	case "mysql":
		connString = fmt.Sprintf("%s:%s@tcp(%s", m.User, m.Pwd, m.Host)
		if m.Port > 0 {
			connString = fmt.Sprintf("%s:%d)/%s", connString, m.Port, m.DBInfo.Name)
		} else {
			connString = fmt.Sprintf("%s)/%s", connString, m.DBInfo.Name)
		}
		connString = fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", connString)
	case "postgres":
		connString = fmt.Sprintf("host=%s ", m.Host)
		if m.Port > 0 {
			connString = fmt.Sprintf("%s:%d", connString, m.Port)
		}
		connString = fmt.Sprintf("%s user=%s dbname=%s sslmode=disable password=%s", connString, m.User, m.DBInfo.Name, m.Pwd)
	}
	return
}

// GetChilds 获取所有子表名和外键并去重
func (m *Form) GetChilds() map[string]string {
	childs := make(map[string]string)
	for _, child := range m.Childs {
		childs[child.Name] = child.Self
	}
	return childs
}

// Tags 根据字段名返回gorm标记,例如: Tags(name)->gorm:"column:name" json:"name"
func Tags(columnName string) template.HTML {
	if strings.ToUpper(columnName) == "ID" {
		return template.HTML(" `gorm:\"column:" + columnName + ";primary_key\" json:\"" + columnName + "\"` ")
	}
	return template.HTML(" `gorm:\"column:" + columnName + "\" json:\"" + columnName + "\"` ")
}

// AssoTags 根据子表信息返回gorm关联标记,例如
// {Name: "score",  Self: "stuid",  Link:"id"}, 返回
// gorm:"ForeignKey:stuid; AssociationForeignKey:id"  json:"score_id"
func AssoTags(c Child) template.HTML {
	return template.HTML(" `gorm:\"ForeignKey:" + c.Self + ";AssociationForeignKey:" + c.Link + "\" json:\"" + c.Name + "_" + c.Link + "\"` ")
}

// ColumnAndType 根据字段信息返回参数定义，例如
// [
//   {Name: "id",   Type: "int"},
//   {Name: "sex", Type: "text"}
// ]
// 把类型text转换成string最后返回: id int,sex string
func ColumnAndType(table_schema []Field) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.Name+" "+TypeConvert(t.Type))
	}
	return strings.Join(result, ",")
}

// ColumnWithPostfix  根据字段名返回sql条件，例如参数
// columns= ["name","sex"]  postfix= "=？"  sep= "and"
// 返回 name=？ and sex=？
func ColumnWithPostfix(columns []string, postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+postfix)
	}
	return strings.Join(result, sep)
}

// TypeConvert 将数据库数据类型转换go数据类型
// text->string、 datetime->*time.Time、 decimal->float64
func TypeConvert(str string) string {
	switch str {
	case "text":
		return "string"

	case "datetime":
		return "*time.Time"

	case "decimal":
		return "float64"
	default:
		return str //int or bool
	}
}

// ExportColumn 将列名转换成结构体字段名,例如：user_name->UserName
func ExportColumn(columnName string) string {
	columnItems := strings.Split(columnName, "_")
	for i := 0; i < len(columnItems); i++ {
		columnItems[i] = strings.Title(columnItems[i])
	}
	return strings.Join(columnItems, "")
}

/*
//所有字段名
func (m *Form) ColumnNames() []string {
	result := make([]string, 0, len(m.Fields))
	for _, t := range m.Fields {
		result = append(result, t.Name)
	}
	return result
}
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

//字段总数
func (m *Form) ColumnCount() int {
	return len(m.Fields)
}
func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}
*/
