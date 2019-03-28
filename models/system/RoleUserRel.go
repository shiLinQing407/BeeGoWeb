package system

import (
	"github.com/astaxie/beego/orm"
	"BeeGoWeb/models/common"
)

type RoleUserRel struct {
	Id         int `orm: "auto"`
	Role       *Role `orm:"rel(fk)"`  //外键
	User       *User `orm:"rel(fk)" ` // 外键
	CreateTime int64 //创建时间
}

var roleUserRelTable = common.TableName("role_user_rel")

func (c *RoleUserRel) GetTableName() string {
	return roleUserRelTable
}

func (c *RoleUserRel) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *RoleUserRel) Add() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *RoleUserRel) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func RoleUserRelGetById(id int) (*RoleUserRel, error) {
	c := new(RoleUserRel)
	err := orm.NewOrm().QueryTable(menuTable).Filter("id", id).One(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func RoleUserRelGetByName(name string) (*RoleUserRel, error) {
	c := new(RoleUserRel)
	err := orm.NewOrm().QueryTable(menuTable).Filter("name", name).One(c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func RoleUserRelList(page, pageSize int, filters ...interface{}) ([]*RoleUserRel, int64) {
	offset := (page - 1) * pageSize
	menus := make([]*RoleUserRel, 0)
	query := orm.NewOrm().QueryTable(menuTable)
	if len(filters) > 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-sort", "-id").Limit(pageSize, offset).All(&menus)
	return menus, total
}
