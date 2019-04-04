/*
@Time : 2019/4/1 23:41 
@Author : shilinqing
@File : news 资讯
*/
package news

import (
	"BeeGoWeb/models/common"
	"github.com/astaxie/beego/orm"
)

//资讯
type News struct {
	Id          	int		`json:"id" orm:"auto"`
	Class			*Class	`json:"class_id" orm:"rel(fk)"`
	Tags			string	`json:"tags"`
	Title			string	`json:"title"`
	Content			string	`json:"content"`
	CreateTime  	int64	`json:"create_time"`
}

//资讯分类
type Class struct{
	Id          	int		`orm:"auto"`
	ClassNameEn		string	`json:"class_name_en"`
	ClassNameZh		string	`json:"class_name_zh"`
	CreateTime  	int64	`json:"create_time"`
}

var newsTable = common.TableName("news")

var newsClassTable = common.TableName("news_class")

/*资讯方法*/
func (c *News) GetTableName() string {
	return newsTable
}

func (c *News) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *News) Add() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *News) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func GetList(page, pageSize int, filters ...interface{}) ([]*News, int64){
	offset := (page - 1) * pageSize
	news := make([]*News, 0)
	query := orm.NewOrm().QueryTable(newsTable)
	if len(filters) > 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count() //数据总数
	query.OrderBy("-id").Limit(pageSize, offset).All(&news)
	return news, total
}

func NewsListGrid(page, pageSize int, filters ...interface{}) ([]News, int64) {
	data, total := GetList(page, pageSize)
	list := make([]News, len(data))
	for i, item := range data {
		list[i] = *item
	}
	return list, total
}

func (c *News) FindById(id int) (*News, error) {
	err := orm.NewOrm().QueryTable(newsTable).Filter("id", id).One(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*资讯方法*/

/*资讯分类方法*/
func (c *Class) GetTableName() string {
	return newsClassTable
}
func (c *Class) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Class) Add() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Class) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}
//资讯分类列表
func (c *Class) GetClassList(filters ...interface{}) ([]*Class, int64){
	classes := make([]*Class, 0)
	query := orm.NewOrm().QueryTable(newsClassTable)
	if len(filters) > 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-sort", "-id").All(&classes)
	return classes, total
}

func (c *Class) FindById(id int) (*Class, error) {
	err := orm.NewOrm().QueryTable(newsTable).Filter("id", id).One(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*资讯分类方法*/