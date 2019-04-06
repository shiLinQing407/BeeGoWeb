/*
@Time : 2019/4/2 10:32 
@Author : shilinqing
@File : news_class
*/
package controllers

import (
	"BeeGoWeb/models/news"
	"BeeGoWeb/utils"
)

type NewsClassController struct {
	BaseController
}

func(this *NewsClassController) Index(){
	this.display()
}

func(this *NewsClassController) LoadList(){
	is_json ,_ := this.GetInt("is_json", 0)
	keyword := this.GetString("js_keyword", "")
	startTime := this.GetString("startTime", "")
	endTime := this.GetString("endTime", "")
	filter := make(map[string]interface{})
	if keyword != ""{
		filter["class_name_zh__icontains"] = keyword //模糊查询
	}
	if startTime != "" {
		filter["create_time__gt"] = utils.TimeParseInLocation(startTime)
	}
	if endTime != "" {
		filter["create_time__lt"] = utils.TimeParseInLocation(endTime)
	}
	data, count := news.ClassListGrid(this.page, this.pageSize, filter)
	if is_json == 0{
		this.toDataGrid(data, count)
	}else{
		this.ReturnSuccessJson(data)
	}
}