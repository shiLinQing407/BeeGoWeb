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

type News struct {
	Id          int    `json:"id" orm:"auto"`
	Tags        string `json:"tags"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreateTime  int64  `json:"create_time"`
	NewsClass *NewsClass  `json:"news_class" orm:"rel(one);on_delete(do_nothing);"`
}

var newsTable = common.TableName("news")

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

func GetList(page, pageSize int, filters map[string]interface{}) ([]*News, int64) {
	news := make([]*News, 0)
	query := orm.NewOrm().QueryTable(newsTable)
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Filter(key, value)
		}
	}
	total, _ := query.Count() //数据总数
	var offset int
	if page > 0 && pageSize > 0 {
		offset = (page - 1) * pageSize
	}

	query.OrderBy("-id").Limit(pageSize, offset).RelatedSel().All(&news)

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

func (c *News) FindById(id int) (error) {
	err := orm.NewOrm().QueryTable(newsTable).Filter("id", id).One(c)
	if err != nil {
		return err
	}
	return nil
}
