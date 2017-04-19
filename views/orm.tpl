{{$modelName := .Flag }}
{{$modelTile := .Lang.Cn}}
{{$modelConn := .Conn}}
{{$modelDriver := .Driver}}

package {{.Model}}

type {{$modelName}} struct {
{{range .Fields}} {{.Flag}} {{.Type | TypeConvert}} {{.Name | Tags}} // {{.Lang.Cn}}
{{end}}
{{range .Childs}} {{.Name | ExportColumn}}{{.Link | ExportColumn}} []{{.Name | ExportColumn}} {{AssoTags .}}
{{end}}
}

// TableName 将{{$modelName}}映射为{{.Name}}
func (table {{$modelName}}) TableName() string {
    return "{{.Base}}.dbo.{{.Name}}"
}

func New{{$modelName}}() *{{$modelName}} {
    table := new({{$modelName}})

    return table
}

// Get{{$modelName}}s: 获取所有{{$modelTile}}记录
func Get{{$modelName}}s(qs map[string]interface{},fields []string) ([]{{$modelName}}, error) {
    db, err := gorm.Open("{{$modelDriver}}","{{$modelConn}}")
    if err != nil {
        return nil, fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()
        
    var records []{{$modelName}}
        
    query := db.
	{{range .Childs}}
		Preload("{{.Name | ExportColumn}}{{.Link | ExportColumn}}").
	{{end}}
		Where(qs).
		Select(fields).
		Find(&records)
    if query.Error != nil {
        return nil, fmt.Errorf("获取所有{{$modelTile}}记录失败，%s", query.Error)
    }           
    return records, nil  
}

{{if .HavePk}}
// Get{{$modelName}}ByPK: 根据主键获取{{$modelTile}}记录
func Get{{$modelName}}ByPK({{.PkColumnsSchema | ColumnAndType}}) (*{{$modelName}}, error) {
    db, err := gorm.Open("{{$modelDriver}}","{{$modelConn}}")
    if err != nil {
        return nil, fmt.Errorf("创建orm引擎失败, %s", err)
    }
    defer db.Close()

    var record {{$modelName}}

    query := db.
	{{range .Childs}}
		Preload("{{.Name | ExportColumn}}{{.Link | ExportColumn}}").
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
    db, err := gorm.Open("{{$modelDriver}}","{{$modelConn}}")
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

// Insert: 新建{{$modelTile}}记录
func (table *{{$modelName}}) Insert() (error) {
    db, err := gorm.Open("{{$modelDriver}}","{{$modelConn}}")
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
    db, err := gorm.Open("{{$modelDriver}}","{{$modelConn}}")
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
