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
    "cn": "�����������뵥",
    "en": "Additional application program",
    "tw": "��ʽ������Ո��"
  },
  "fields": [
   {
      "name": "ID",
      "flag": "ID",
      "lang": {
        "cn": "���",
        "en": "ID",
        "tw": "��̖"
      },
      "type": "int"
    },
    {
      "name": "BenefitEstimate",
      "flag": "BenefitEstimate",
      "lang": {
        "cn": "Ч������",
        "en": "Benefit evaluation",
        "tw": "Ч���u��"
      },
      "type": "text"
    },
    {
      "name": "BillNo20",
      "flag": "BillNo20",
      "lang": {
        "cn": "�������޵����",
        "en": "Program amendment order number",
        "tw": "��ʽ���ކξ�̖"
      },
      "type": "text"
    },
    {
      "name": "systemname",
      "flag": "Systemname",
      "lang": {
        "cn": "��������ϵͳ����",
        "en": "The program reads the system name",
        "tw": "��ʽ����ϵ�y���Q"
      },
      "type": "text"
    },
    {
      "name": "Enter_Time",
      "flag": "EnterTime",
      "lang": {
        "cn": "��������",
        "en": "Enter_Time",
        "tw": "��Ո����"
      },
      "type": "datetime"
    },
    {
      "name": "Enter_User",
      "flag": "EnterUser",
      "lang": {
        "cn": "������",
        "en": "Enter_User",
        "tw": "��Ո��"
      },
      "type": "text"
    },
    {
      "name": "FineshTime",
      "flag": "FineshTime",
      "lang": {
        "cn": "Ԥ�����ʱ��",
        "en": "Estimated time of completion",
        "tw": "�AӋ��ɕr�g"
      },
      "type": "datetime"
    },
    {
      "name": "sjwcrw",
      "flag": "Sjwcrw",
      "lang": {
        "cn": "ʵ���������",
        "en": "actual finishing date",
        "tw": "���H�������"
      },
      "type": "datetime"
    },
    {
      "name": "ItemMoney",
      "flag": "ItemMoney",
      "lang": {
        "cn": "��Ŀ���",
        "en": "Item amount",
        "tw": "�Ŀ���~"
      },
      "type": "decimal"
    },
    {
      "name": "maintenance_cause",
      "flag": "MaintenanceCause",
      "lang": {
        "cn": "ά��ԭ�����������",
        "en": "Description of maintenance cause or phenomenon",
        "tw": "�S�oԭ���F������"
      },
      "type": "text"
    },
    {
      "name": "outcome",
      "flag": "Outcome",
      "lang": {
        "cn": "��Ҫ�ﵽ�������",
        "en": "Need to achieve results description",
        "tw": "��Ҫ�_���Y������"
      },
      "type": "text"
    },
    {
      "name": "PayCost",
      "flag": "PayCost",
      "lang": {
        "cn": "���ѳɱ�",
        "en": "Cost cost",
        "tw": "���M�ɱ�"
      },
      "type": "decimal"
    },
    {
      "name": "PhoneNO",
      "flag": "PhoneNO",
      "lang": {
        "cn": "��ϵ�绰",
        "en": "Contact number",
        "tw": "�M�Ԓ"
      },
      "type": "text"
    },
    {
      "name": "SH_Status",
      "flag": "SHStatus",
      "lang": {
        "cn": "���״̬",
        "en": "SH_Status",
        "tw": "���ˠ�B"
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
	BenefitEstimate         string     `gorm:"column:BenefitEstimate" json:"BenefitEstimate"`                     // Ч������
	BillNo20                string     `gorm:"column:BillNo20" json:"BillNo20"`                                   // �������޵����
	EnterTime               *time.Time `gorm:"column:Enter_Time" json:"Enter_Time"`                               // ��������
	EnterUser               string     `gorm:"column:Enter_User" json:"Enter_User"`                               // ������
	FineshTime              *time.Time `gorm:"column:FineshTime" json:"FineshTime"`                               // Ԥ�����ʱ��
	ID                      int        `gorm:"column:ID;primary_key" json:"ID"`                                   // ���
	ItemMoney               float64    `gorm:"column:ItemMoney" json:"ItemMoney"`                                 // ��Ŀ���
	MaintenanceCause        string     `gorm:"column:maintenance_cause" json:"maintenance_cause"`                 // ά��ԭ�����������
	Outcome                 string     `gorm:"column:outcome" json:"outcome"`                                     // ��Ҫ�ﵽ�������
	PayCost                 float64    `gorm:"column:PayCost" json:"PayCost"`                                     // ���ѳɱ�
	PhoneNO                 string     `gorm:"column:PhoneNO" json:"PhoneNO"`                                     // ��ϵ�绰
	SHStatus                int        `gorm:"column:SH_Status" json:"SH_Status"`                                 // ���״̬
	Sjwcrw                  *time.Time `gorm:"column:sjwcrw" json:"sjwcrw"`                                       // ʵ���������
	Systemname              string     `gorm:"column:systemname" json:"systemname"`                               // ��������ϵͳ����

	TblinformationReworkapplySub2ID           []TblinformationReworkapplySub2           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub2_ID"`
	TblinformationReworkapplySub3ID           []TblinformationReworkapplySub3           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub3_ID"`
	VwcomputerdevelopUsersDesigner            []VwcomputerdevelopUsers                  `gorm:"ForeignKey:staff_name;AssociationForeignKey:designer;save_associations:false" json:"vwcomputerdevelop_users_designer"`
	TblinformationReworkapplyFunctionDetailID []TblinformationReworkapplyFunctionDetail `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_FunctionDetail_ID"`
	TblinformationReworkapplySub4ID           []TblinformationReworkapplySub4           `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:"tblinformation_reworkapply_Sub4_ID"`
}

//TableName ��TblinformationReworkapplyӳ��Ϊtblinformation_reworkapply
func (table TblinformationReworkapply) TableName() string {
	return "test.dbo.tblinformation_reworkapply"
}

func NewTblinformationReworkapply() *TblinformationReworkapply {
	table := new(TblinformationReworkapply)

	return table
}

//GetTblinformationReworkapplys: ��ȡ���г����������뵥��¼
func GetTblinformationReworkapplys(qs map[string]interface{}, fields []string) ([]TblinformationReworkapply, error) {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return nil, fmt.Errorf("����orm����ʧ��, %s", err)
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
		return nil, fmt.Errorf("��ȡ���г����������뵥��¼ʧ�ܣ�%s", query.Error)
	}

	return records, nil
}

//GetTblinformationReworkapplyByPK: ����������ȡ�����������뵥��¼
func GetTblinformationReworkapplyByPK(ID int) (*TblinformationReworkapply, error) {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return nil, fmt.Errorf("����orm����ʧ��, %s", err)
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
		return nil, fmt.Errorf("��ȡ�����������뵥��¼ʧ�ܣ�%s", query.Error)
	}

	return &record, nil
}

//Update: ���³����������뵥��¼
func (table *TblinformationReworkapply) Update() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("����orm����ʧ��, %s", err)
	}
	defer db.Close()

	//��������
	tx := db.Begin()

	up := tx.Save(&table)
	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("���³����������뵥ʧ��, %s", up.Error)
	}
	tx.Commit()

	return nil
}

//Insert: �½������������뵥��¼
func (table *TblinformationReworkapply) Insert() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("����orm����ʧ��, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("�½������������뵥��¼ʧ��, %s", save.Error)
	}
	return nil
}

//Delete: ɾ�������������뵥��¼
func (table TblinformationReworkapply) Delete() error {
	db, err := gorm.Open("mssql", "server=127.0.0.1;database=test;user id=sa;password=123456;encrypt=disable")
	if err != nil {
		return fmt.Errorf("����orm����ʧ��, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("ɾ�������������뵥��¼ʧ��, %s", del.Error)
	}

	return nil
}
```