# apigen
According to the data dictionary automatically generate beego framework api, model layer using gorm

# Installation
> go get github.com/hwfy/apigen

# Usage
If the data dictionary is stored in redis, where "flag" is the gorm model field, "base" is the database name, "flow" is the name of the process, "self" in childs is foreign key, "link" is the associated primary table field:
```json
{
  "name": "tblinformation_reworkapply",
  "flag": "TblinformationReworkapply",
  "base": "test",
  "flow": "myexample:2:40",
  "lang": {
    "cn": "程序增修申请单",
    "en": "Additional application program",
    "tw": "程式增修申請單"
  },
  "fields": [
   {
      "name": "ID",
      "flag": "ID",
      "lang": {
        "cn": "序号",
        "en": "ID",
        "tw": "序號"
      },
      "type": "int"
    },
    {
      "name": "BenefitEstimate",
      "flag": "BenefitEstimate",
      "lang": {
        "cn": "效益评估",
        "en": "Benefit evaluation",
        "tw": "效益評估"
      },
      "type": "text"
    },
    {
      "name": "BillNo20",
      "flag": "BillNo20",
      "lang": {
        "cn": "程序增修单编号",
        "en": "Program amendment order number",
        "tw": "程式增修單編號"
      },
      "type": "text"
    },
    {
      "name": "systemname",
      "flag": "Systemname",
      "lang": {
        "cn": "程序增修系统名称",
        "en": "The program reads the system name",
        "tw": "程式增修系統名稱"
      },
      "type": "text"
    },
    {
      "name": "Enter_Time",
      "flag": "EnterTime",
      "lang": {
        "cn": "申请日期",
        "en": "Enter_Time",
        "tw": "申請日期"
      },
      "type": "datetime"
    },
    {
      "name": "Enter_User",
      "flag": "EnterUser",
      "lang": {
        "cn": "申请人",
        "en": "Enter_User",
        "tw": "申請人"
      },
      "type": "text"
    },
    {
      "name": "FineshTime",
      "flag": "FineshTime",
      "lang": {
        "cn": "预计完成时间",
        "en": "Estimated time of completion",
        "tw": "預計完成時間"
      },
      "type": "datetime"
    },
    {
      "name": "sjwcrw",
      "flag": "Sjwcrw",
      "lang": {
        "cn": "实际完成日期",
        "en": "actual finishing date",
        "tw": "實際完成日期"
      },
      "type": "datetime"
    },
    {
      "name": "ItemMoney",
      "flag": "ItemMoney",
      "lang": {
        "cn": "项目金额",
        "en": "Item amount",
        "tw": "項目金額"
      },
      "type": "decimal"
    },
    {
      "name": "maintenance_cause",
      "flag": "MaintenanceCause",
      "lang": {
        "cn": "维护原因或现象描述",
        "en": "Description of maintenance cause or phenomenon",
        "tw": "維護原因或現象描述"
      },
      "type": "text"
    },
    {
      "name": "outcome",
      "flag": "Outcome",
      "lang": {
        "cn": "需要达到结果描述",
        "en": "Need to achieve results description",
        "tw": "需要達到結果描述"
      },
      "type": "text"
    },
    {
      "name": "PayCost",
      "flag": "PayCost",
      "lang": {
        "cn": "花费成本",
        "en": "Cost cost",
        "tw": "花費成本"
      },
      "type": "decimal"
    },
    {
      "name": "PhoneNO",
      "flag": "PhoneNO",
      "lang": {
        "cn": "联系电话",
        "en": "Contact number",
        "tw": "聯繫電話"
      },
      "type": "text"
    },
    {
      "name": "SH_Status",
      "flag": "SHStatus",
      "lang": {
        "cn": "审核状态",
        "en": "SH_Status",
        "tw": "稽核狀態"
      },
      "type": "int"
    }
  ],
  "parent": "",
  "childs": [
    {
      "name": "tblinformation_reworkapply_Sub2",
      "self": "ParentID",
      "link": "ID"
    },
    {
      "name": "tblinformation_reworkapply_Sub3",
      "self": "ParentID",
      "link": "ID"
    },
    {
      "name": "vwcomputerdevelop_users",
      "self": "staff_name",
      "link": "designer"
    },
    {
      "name": "tblinformation_reworkapply_FunctionDetail",
      "self": "ParentID",
      "link": "ID"
    },
    {
      "name": "tblinformation_reworkapply_Sub4",
      "self": "ParentID",
      "link": "ID"
    }
  ]
}
```
There are two ways to generate api, one from the command line generation, one from the file batch generation:
##1.command line
![Alt text](./11111111.png)

##2.file
replace -form with -forms = file name, you can also enter "apigen" generated directly, the other command parameters from "app.conf" load


----------
After the controller will generate control files, routers produce routing files, models produce all the associated table model file, for example:
```go
package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type TblinformationReworkapply struct {
	BenefitEstimate         string     `gorm:"column:BenefitEstimate" json:"BenefitEstimate"`                     // 效益评估
	BillNo20                string     `gorm:"column:BillNo20" json:"BillNo20"`                                   // 程序增修单编号
	EnterTime               *time.Time `gorm:"column:Enter_Time" json:"Enter_Time"`                               // 申请日期
	EnterUser               string     `gorm:"column:Enter_User" json:"Enter_User"`                               // 申请人
	FineshTime              *time.Time `gorm:"column:FineshTime" json:"FineshTime"`                               // 预计完成时间
	ID                      int        `gorm:"column:ID;primary_key" json:"ID"`                                   // 序号
	ItemMoney               float64    `gorm:"column:ItemMoney" json:"ItemMoney"`                                 // 项目金额
	MaintenanceCause        string     `gorm:"column:maintenance_cause" json:"maintenance_cause"`                 // 维护原因或现象描述
	Outcome                 string     `gorm:"column:outcome" json:"outcome"`                                     // 需要达到结果描述
	PayCost                 float64    `gorm:"column:PayCost" json:"PayCost"`                                     // 花费成本
	PhoneNO                 string     `gorm:"column:PhoneNO" json:"PhoneNO"`                                     // 联系电话
	SHStatus                int        `gorm:"column:SH_Status" json:"SH_Status"`                                 // 审核状态
	Sjwcrw                  *time.Time `gorm:"column:sjwcrw" json:"sjwcrw"`                                       // 实际完成日期
	Systemname              string     `gorm:"column:systemname" json:"systemname"`                               // 程序增修系统名称

	TblinformationReworkapplySub2ID           []TblinformationReworkapplySub2           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub2_ID"`
	TblinformationReworkapplySub3ID           []TblinformationReworkapplySub3           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub3_ID"`
	VwcomputerdevelopUsersDesigner            []VwcomputerdevelopUsers                  `gorm:"ForeignKey:staff_name;AssociationForeignKey:designer;save_associations:false" json:"vwcomputerdevelop_users_designer"`
	TblinformationReworkapplyFunctionDetailID []TblinformationReworkapplyFunctionDetail `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_FunctionDetail_ID"`
	TblinformationReworkapplySub4ID           []TblinformationReworkapplySub4           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub4_ID"`
}

//TableName 将TblinformationReworkapply映射为tblinformation_reworkapply
func (table TblinformationReworkapply) TableName() string {
	return "test.dbo.tblinformation_reworkapply"
}

func NewTblinformationReworkapply() *TblinformationReworkapply {
	table := new(TblinformationReworkapply)

	return table
}

//GetTblinformationReworkapplys: 获取所有程序增修申请单记录
func GetTblinformationReworkapplys(qs map[string]interface{}, fields []string) ([]TblinformationReworkapply, error) {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []TblinformationReworkapply

	query := db.
		Preload("TblinformationReworkapplySub2ID").
		Preload("TblinformationReworkapplySub3ID").
		Preload("VwcomputerdevelopUsersDesigner").
		Preload("TblinformationReworkapplyFunctionDetailID").
		Preload("TblinformationReworkapplySub4ID").
		Where(qs).
		Select(fields).
		Find(&records)
	if query.Error != nil {
		return nil, fmt.Errorf("获取所有程序增修申请单记录失败，%s", query.Error)
	}

	return records, nil
}

//GetTblinformationReworkapplyByPK: 根据主键获取程序增修申请单记录
func GetTblinformationReworkapplyByPK(ID int) (*TblinformationReworkapply, error) {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record TblinformationReworkapply

	query := db.
		Preload("TblinformationReworkapplySub2ID").
		Preload("TblinformationReworkapplySub3ID").
		Preload("VwcomputerdevelopUsersDesigner").
		Preload("TblinformationReworkapplyFunctionDetailID").
		Preload("TblinformationReworkapplySub4ID").
		Where("ID=?", ID).
		Find(&record)
	if query.Error != nil {
		return nil, fmt.Errorf("获取程序增修申请单记录失败，%s", query.Error)
	}

	return &record, nil
}

//Update: 更新程序增修申请单记录
func (table *TblinformationReworkapply) Update() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	//开启事务
	tx := db.Begin()

	up := tx.Save(&table)
	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新程序增修申请单失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

//Insert: 新建程序增修申请单记录
func (table *TblinformationReworkapply) Insert() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("新建程序增修申请单记录失败, %s", save.Error)
	}
	return nil
}

//Delete: 删除程序增修申请单记录
func (table TblinformationReworkapply) Delete() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除程序增修申请单记录失败, %s", del.Error)
	}

	return nil
}
```