package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// getDatadictionaryByName 根据表名获取数据字典
func getDatadictionaryByName(name string) (*Form, error) {
	db, err := gorm.Open(newDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var form *Form

	switch ds.Driver {
	case "mysql":
		form, err = getDataDictionaryFromMysql(db, ds.Name, name)
		if err != nil {
			return nil, err
		}
	case "mssql":
		form, err = getDataDictionaryFromSqlserver(db, name)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("数据库驱动不支持%s", ds.Driver)
	}
	return form, form.removalDuplicates()
}

// removalDuplicates 移除重复的父表和子表
func (form *Form) removalDuplicates() error {
	if form.Name == "" || form.Base == "" {
		return fmt.Errorf("表单或数据库名为空!")
	}
	form.Childs = getUniqueChilds(form.Name, form.Childs)
	form.Parents = getUniqueParents(form.Name, form.Parents)

	return nil
}

// getUniqueParents 排除父表名为空、和当前表单名相同、已存在父表集中的父表
func getUniqueParents(form string, parents []Parent) []Parent {
	for i := 0; i < len(parents); i++ {
		if parents[i].Name == "" ||
			parents[i].Name == form || //父表名和当前表名相同
			parentIsExist(parents[i+1:], parents[i].Name) { //从下一个父表集查找相同
			parents = append(parents[:i], parents[i+1:]...)
			i--
		}
	}
	return parents
}

// getUniqueChilds 排除子表名为空、和当前表单名相同、已存在子表集中的子表
func getUniqueChilds(form string, childs []Child) []Child {
	for i := 0; i < len(childs); i++ {
		if childs[i].Name == "" ||
			childs[i].Name == form ||
			childIsExist(childs[i+1:], childs[i]) {
			childs = append(childs[:i], childs[i+1:]...)
			i--
		}
	}
	return childs
}

// parentIsExist 判断父表名称是否有相同的
func parentIsExist(parents []Parent, name string) bool {
	for _, parent := range parents {
		if name == parent.Name {
			return true
		}
	}
	return false
}

// childIsExist 判断子表名称和主键是否相同
func childIsExist(childs []Child, child Child) bool {
	for _, c := range childs {
		if c.Name == child.Name && c.Link == child.Link {
			return true
		}
	}
	return false
}

// setControllers 设置所有控制器名称
func (form *Form) setControllers(dir string) error {
	var c controller

	names, err := readnames(dir + "/" + c.getDir())
	if err != nil {
		return err
	}
	form.Controllers = names

	return nil
}

// getChilds 获取所有关联的子表,tables必须用make()初始化
func (form *Form) getChilds(tables map[string]string) {
	for _, child := range form.Childs {
		if form.Name == child.Name {
			continue
		}
		tables[child.Name] = child.Self

		table, err := getDatadictionaryByName(child.Name)
		if err != nil {
			return
		}
		if len(table.Childs) != 0 {
			table.getChilds(tables)
		}
	}
}

//// removalDuplicate 去掉重复父表和子表
//func (form *Form) removalDuplicate() {
//	var parents []Parent

//	for _, parent := range form.Parents {
//		//如果父表名和当前表名相同
//		if parent.Name == form.Name {
//			continue
//		}
//		if parentIsExist(parents, parent.Name) {
//			continue
//		}
//		parents = append(parents, parent)
//	}
//	var childs []Child

//	for _, child := range form.Childs {
//		//如果子表名和当前表名相同
//		if child.Name == form.Name {
//			continue
//		}
//		if childIsExist(childs, child) {
//			continue
//		}
//		childs = append(childs, child)
//	}
//	form.Parents = parents
//	form.Childs = childs
//}
