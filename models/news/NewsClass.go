/*
@Time : 2019/4/1 23:41 
@Author : shilinqing
@File : news class 资讯分类
*/
package news

import (
	"BeeGoWeb/models/common"
	"github.com/astaxie/beego/orm"
)

type NewsClass struct {
	Id          int    `json:"id" orm:"auto"`
	ClassNameEn string `json:"class_name_en"`
	ClassNameZh string `json:"class_name_zh"`
	CreateTime  int64  `json:"create_time"`
}

var newsClassTable = common.TableName("news_class")

/*资讯分类方法*/
func (c *NewsClass) GetTableName() string {
	return newsClassTable
}

func (c *NewsClass) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *NewsClass) Add() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *NewsClass) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

//资讯分类列表
func GetClassList(page, pageSize int, filters map[string]interface{}) ([]*NewsClass, int64) {
	classes := make([]*NewsClass, 0)
	query := orm.NewOrm().QueryTable(newsClassTable)
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Filter(key, value)
		}
	}
	total, _ := query.Count() //数据总数
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query.Limit(pageSize, offset)
	}
	query.OrderBy("-id").All(&classes)
	return classes, total
}

func ClassListGrid(page, pageSize int, filters map[string]interface{}) ([]NewsClass, int64) {
	data, total := GetClassList(page, pageSize, filters)
	list := make([]NewsClass, len(data))
	for i, item := range data {
		list[i] = *item
	}
	return list, total
}

func (c *NewsClass) FindById(id int) (error) {
	err := orm.NewOrm().QueryTable(c.GetTableName()).Filter("id", id).One(c)
	if err != nil {
		return err
	}
	return nil
}
