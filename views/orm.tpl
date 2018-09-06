{{$modelName 	:=	.Name | ExportStructTag}}
{{$modelTile 	:=	.Chinese}}
{{$modelDriver	:=	.Driver}}
{{$modelBase	:=	.Base | ToLower}}
{{$modelConn 	:=	.DataSource.Name | ToLower}}

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type {{$modelName}} struct {
{{range .Fields}} {{Tags .}} 
{{end}}
{{range .Childs}} {{AssoTags .}}
{{end}}
}

// TableName 将{{$modelName}}映射为{{.Name}}
func (table {{$modelName}}) TableName() string {
	{{if eq $modelBase ""}}
		return "{{.Name}}"
	{{else}}
    	return {{$modelBase}}DbName()+"{{.FormName}}"
	{{end}}
}

func New{{$modelName}}() *{{$modelName}} {
    table := new({{$modelName}})

    return table
}

// Get{{$modelName}}s: 获取所有{{$modelTile}}记录
func Get{{$modelName}}s(qs map[string]interface{},joins,fields []string,sortby string,offset int64,limit int64) ([]{{$modelName}}, error) {
    db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return nil, fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

    if sortby != "" {
		db=db.Order(sortby)
	}  
	for _, join := range joins {
		db=db.Joins(join)
	}
    {{if eq  $modelDriver "mssql"}}
		if offset > 0 {
		var ids []string
		var id string

		rows, _  := db.
			Table("{{.Name}}").
			Select("TOP " + strconv.FormatInt(offset, 10) + " {{range .PkColumns}}{{.}}{{end}}").
			Where(qs).
			Rows()
		
		for rows.Next() {
			rows.Scan(&id)
			ids = append(ids, id)
		}
		db = db.Where("{{range .PkColumns}}{{.}}{{end}} NOT IN(?)", ids)
	}
		limitStr := strconv.FormatInt(limit, 10)
		fields[0] = "top " + limitStr + fields[0]
	{{else}}
		db=db.Offset(offset).Limit(limit)
	{{end}}
	var records []{{$modelName}}  
   
    query := db.
	{{range .Childs}}
		Preload("{{.Name | ExportStructTag}}{{.Link | ExportStructTag}}").
	{{end}}
		Select(fields).
		Where(qs).
		Find(&records)

    if query.Error != nil {
        return nil, fmt.Errorf("获取所有{{$modelTile}}记录失败，%s", query.Error)
    }           
    return records, nil  
}

// Get{{$modelName}}ByCondition: 根据给定条件获取{{$modelTile}}记录
func Get{{$modelName}}ByCondition(qs interface{}, args ...interface{}) ([]{{$modelName}}, error) {
	db, err := gorm.Open({{$modelConn}}DataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []{{$modelName}}

	query := db.
		{{range .Childs}}
			Preload("{{.Name | ExportStructTag}}{{.Link | ExportStructTag}}").
		{{end}}
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取{{$modelTile}}记录失败，%s", query.Error)
	}
	return records, nil
}

{{if .HavePk}}
// Get{{$modelName}}ByPK: 根据主键获取{{$modelTile}}记录
func Get{{$modelName}}ByPK({{.PkColumnsSchema | ColumnAndType}}) (*{{$modelName}}, error) {
    db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return nil, fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

    var record {{$modelName}}

    query := db.
	{{range .Childs}}
		Preload("{{.Name | ExportStructTag}}{{.Link | ExportStructTag}}").
	{{end}}
		Where("{{ColumnWithPostfix .PkColumns "=?" " and "}}", {{range .PkColumns}}{{.}},{{end}}).
		Find(&record)

    if query.Error != nil {
        return nil, fmt.Errorf("获取{{$modelTile}}记录失败，%s", query.Error)
    }
    return &record, nil
}
{{end}}

// Update: 更新{{$modelTile}}记录
func (table *{{$modelName}}) Update() (error) {
    db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

    //开启事务
    tx := db.Begin()
    up := tx.Save(&table)
	
    if up.Error != nil {
        tx.Rollback()
        return fmt.Errorf("更新{{$modelTile}}失败, %s", up.Error)
    }
    tx.Commit()

    return nil
}

// UpdateByCondition: 根据指定条件更新{{$modelTile}}记录
func (table *{{$modelName}}) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&{{$modelName}}{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新{{$modelTile}}失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// Insert: 新建{{$modelTile}}记录
func (table *{{$modelName}}) Insert() (error) {
    db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

	save := db.Save(&table)
			
    if save.Error != nil {
        return fmt.Errorf("新建{{$modelTile}}记录失败, %s", save.Error)
    }
    return nil
}

// Delete: 删除{{$modelTile}}记录
func (table {{$modelName}}) Delete() (error) {
    db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

    del := db.Delete(&table)

    if del.Error != nil {
       return fmt.Errorf("删除{{$modelTile}}记录失败, %s", del.Error)
    }
    return nil
}

// DeleteByCondition: 根据指定条件删除{{$modelTile}}记录
func (table {{$modelName}}) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open({{$modelConn}}DataSource())
    if err != nil {
        return fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除{{$modelTile}}记录失败, %s", del.Error)
	}
	return nil
}