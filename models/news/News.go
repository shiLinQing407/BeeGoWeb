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
	ClassId		int		`json:"class_id"`
	Tags			string	`json:"tags"`
	Title			string	`json:"title"`
	Content			string	`json:"content"`
	CreateTime  	int64	`json:"create_time"`
}

//资讯分类
type NewsClass struct{
	Id          	int		`json:"id" orm:"auto"`
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

func GetList(page, pageSize int, filters map[string]interface{}) ([]*News, int64){
	news := make([]*News, 0)
	query := orm.NewOrm().QueryTable(newsTable)
	if len(filters) > 0 {
		for key, value := range filters{
			query = query.Filter(key, value)
		}
	}
	total, _ := query.Count() //数据总数
	if page > 0 && pageSize > 0{
		offset := (page - 1) * pageSize
		query.Limit(pageSize, offset)
	}
	query.OrderBy("-id").All(&news)
	return news, total
}

func NewsListGrid(page, pageSize int, filters map[string]interface{}) ([]News, int64) {
	data, total := GetList(page, pageSize, filters)
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
func GetClassList(page, pageSize int, filters map[string]interface{}) ([]*NewsClass, int64){
	classes := make([]*NewsClass, 0)
	query := orm.NewOrm().QueryTable(newsClassTable)
	if len(filters) > 0 {
		for key, value := range filters{
			query = query.Filter(key, value)
		}
	}
	total, _ := query.Count() //数据总数
	if page > 0 && pageSize > 0{
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

func (c *NewsClass) FindById(id int) (*NewsClass, error) {
	err := orm.NewOrm().QueryTable(newsTable).Filter("id", id).One(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*资讯分类方法*/