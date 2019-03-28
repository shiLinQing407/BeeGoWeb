package system

import (
	"github.com/astaxie/beego/orm"
	"BeeGoWeb/models/common"
	)

type Config struct {
	Id int	`orm: "auto"` //autoincrement
	Name string	//配置名称
	Desc string	//配置说明
	GroupType int //配置分组
	Settings string //配置值
	Type int	//类型
	Status int	//状态
	Sort int	//排序
	CreateTime int64 //创建时间
	UpdateTime int64	//更新时间
}

var configTable = common.TableName("config")

func (c *Config) TableName() string {
	return configTable
}

func (c *Config) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Config) Add() error{
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) Delete() error{
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func ConfigGetById(id int) (*Config, error) {
	c := new(Config)
	err := orm.NewOrm().QueryTable(configTable).Filter("id", id).One(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ConfigGetByName(name string) (*Config, error) {
	c := new(Config)
	err := orm.NewOrm().QueryTable(configTable).Filter("name", name).One(c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func ConfigList(page, pageSize int, filters ...interface{}) ([]*Config, int64){
	offset := (page - 1) * pageSize
	configs := make([]*Config, 0)
	query := orm.NewOrm().QueryTable(configTable)
	if len(filters)  > 0 {
		l := len(filters)
		for i := 0; i < l; i+=2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-sort", "-id").Limit(pageSize, offset).All(&configs)
	return configs, total
}