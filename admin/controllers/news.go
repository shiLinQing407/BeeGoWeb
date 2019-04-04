/*
@Time : 2019/4/1 23:39 
@Author : shilinqing
@File : news
*/
package controllers

import (
	"BeeGoWeb/models/news"
)

type NewsController struct {
	BaseController
}

/*资讯列表*/
func (this *NewsController) Index(){
	this.display()
}

func(this *NewsController) LoadList(){
	keyword := this.GetStrings("js_keyword")
	//startTime := this.GetStrings("startTime")
	//endTime := this.GetStrings("endTime")
	filter := make(map[string]interface{})
	if keyword != nil{
		filter["title__icontains"] = keyword[0] //模糊查询 title
	}

	data, count := news.NewsListGrid(this.page, this.pageSize, filter)
	this.toDataGrid(data, count)
}