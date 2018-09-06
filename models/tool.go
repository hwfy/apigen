package models

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"sync"
)

var mu sync.Mutex

// HavePk 判断是否包含主键
func (m *Form) HavePk() bool {
	return len(m.PkColumnsSchema()) > 0
}

// IsNotPK 如果字段不是主键返回true否则false
func (m *Form) IsNotPK(field string) bool {
	for _, pk := range m.PkColumns() {
		if strings.ToUpper(field) == pk {
			return false
		}
	}
	return true
}

// HaveChild 判断是否包含子表
func (m *Form) HaveChild() bool {
	return len(m.Childs) > 0
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

// FormName 获取表单完整的名称
func (m *Form) FormName() string {
	if m.Driver == "mssql" {
		return ".dbo." + m.Name
	} else {
		return "." + m.Name
	}
}

// GetChilds 获取所有子表名和外键并去重
func (m *Form) GetChilds() map[string]string {
	childs := make(map[string]string)

	for _, child := range m.Childs {
		childs[child.Name] = child.Self
	}
	return childs
}

// IsNotPrevChild 判断当前位置的子表名是否和前一个子表名相同
func (m *Form) IsNotPrevChild(index int) bool {
	if index < 1 {
		return true
	}
	return m.Childs[index].Name != m.Childs[index-1].Name
}

// Tags 根据字段名返回gorm标记,例如: Tags(name)->gorm:"column:name" json:"name"
func Tags(field Field) template.HTML {
	mu.Lock()
	defer mu.Unlock()

	if field.Name == "" {
		return ""
	}
	flag := ExportStructTag(field.Name)
	type_ := TypeConvert(field.Type)
	length := strings.SplitN(field.Length, ",", 2)
	column := fmt.Sprintf("column:%s", field.Name)

	if length[0] != "" {
		column += fmt.Sprintf(";size:%s", length[0])
	}
	if strings.ToUpper(field.Name) == "ID" {
		column += fmt.Sprintf(";primary_key;AUTO_INCREMENT")
	}
	return template.HTML(fmt.Sprintf("%s %s `gorm:%q json:%q`//%s",
		flag, type_, column, field.Name, field.Chinese))
}

// AssoTags 根据子表信息返回gorm关联标记
// 例如: {Name: "score",  Self: "stuid",  Link:"id"}
// 返回: ScoreId []Score gorm:"ForeignKey:stuid; AssociationForeignKey:id"  json:"score_id"
func AssoTags(c Child) template.HTML {
	if c.Name == "" || c.Self == "" || c.Link == "" {
		return ""
	}
	name := ExportStructTag(c.Name)
	link := ExportStructTag(c.Link)
	asso := fmt.Sprintf("ForeignKey:%s;AssociationForeignKey:%s", c.Self, c.Link)

	return template.HTML(fmt.Sprintf("%s%s []%s `gorm:%q json:%q`",
		name, link, name, asso, c.Name+"_"+c.Link))
}

// ColumnAndType 根据字段信息返回参数定义，例如
// [
//   {Name: "id",   Type: "int"},
//   {Name: "sex", Type: "text"}
// ]
// 把类型text转换成string,返回字符: id int,sex string
func ColumnAndType(fields []Field) string {
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, field.Name+" "+TypeConvert(field.Type))
	}
	return strings.Join(result, ",")
}

// ColumnWithPostfix 根据字段名返回sql条件
// 例如:  columns= ["name","sex"]  postfix= "=？"  sep= "and"
// 返回:  name=？ and sex=？
func ColumnWithPostfix(columns []string, postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, column := range columns {
		result = append(result, column+postfix)
	}
	return strings.Join(result, sep)
}

// TypeConvert 将数据库字段类型标识转换go数据类型
// 1->string、 2->float64、 3->*time.Time、 4,5->int 6->bool
func TypeConvert(str string) string {
	switch str {
	case "4", "5", "int":
		return "int"

	case "6", "tinyint":
		return "bool"

	case "2", "decimal", "double":
		return "float64"

	case "3", "timestamp", "datetime":
		return "*time.Time"

	default:
		return "string" // "", "1", "varchar","text"
	}
}

// ExportStructTag 将标识名转换成结构体标记名,例如：user_name->UserName
func ExportStructTag(name string) string {
	columnItems := strings.Split(name, "_")

	for i := 0; i < len(columnItems); i++ {
		columnItems[i] = strings.Title(columnItems[i])
	}
	return strings.Join(columnItems, "")
}

// ToLower 转换字符为小写
func ToLower(str string) string {
	return strings.ToLower(str)
}

// readnames if the current directory exists to return a map key is the file name,
// value is to remove the underscore first character in the file name
// -library
//|____book_a.go
//|____book_b.go
//|____food_c.go
//
// readnames(library): map[string]string{"book_a":"BookA","book_b":"BookB","food_c":"FoodC"}
func readnames(dirName string) (map[string]string, error) {
	list := make(map[string]string)

	file, err := os.OpenFile(dirName, os.O_RDONLY, 0644)
	if err != nil {
		return list, err
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		return list, err
	}
	for _, name := range names {
		//Remove the extension
		prefix := strings.SplitN(name, ".", 2)
		//Remove the underline
		items := strings.Split(prefix[0], "_")
		//Combination of strings
		for i := 0; i < len(items); i++ {
			items[i] = strings.Title(items[i])
		}
		list[prefix[0]] = strings.Join(items, "")
	}
	return list, nil
}

//// ColumnNames 所有字段名
//func (m *Form) ColumnNames() []string {
//	result := make([]string, 0, len(m.Fields))
//	for _, t := range m.Fields {
//		result = append(result, t.Name)
//	}
//	return result
//}
//func Join(a []string, sep string) string {
//	return strings.Join(a, sep)
//}

//// ColumnCount 字段总数
//func (m *Form) ColumnCount() int {
//	return len(m.Fields)
//}
//func MakeQuestionMarkList(num int) string {
//	a := strings.Repeat("?,", num)
//	return a[:len(a)-1]
//}
