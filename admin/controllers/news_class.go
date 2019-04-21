/*
@Time : 2019/4/2 10:32 
@Author : shilinqing
@File : news_class
*/
package controllers

import (
	"BeeGoWeb/models/news"
	"BeeGoWeb/utils"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)

type NewsClassController struct {
	BaseController
}

func(this *NewsClassController) Index(){
	this.display()
}

func(this *NewsClassController) LoadList(){
	isJson ,_ := this.GetInt("is_json", 0)
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
	if isJson == 0{
		this.toDataGrid(data, count)
	}else{
		this.ReturnSuccessJson(data)
	}
}


func(this *NewsClassController) Edit(){
	if this.isPost(){
		newsClass := news.NewsClass{}
		o := orm.NewOrm()
		var err error
		//获取form里的值并赋值给结构体， name必须和结构对应
		if err = this.ParseForm(&newsClass); err != nil {
			this.ReturnFailedJson(err, "获取数据失败!")
		}
		newsClass.CreateTime = time.Now().Unix()
		if newsClass.Id == 0 {
			newsClass.CreateTime = time.Now().Unix()
			if _, err := o.Insert(&newsClass); err != nil {
				beego.Error(err.Error())
				this.ReturnFailedJson(err, "添加资讯类别失败!")
			}
		} else {
			if _, err := o.Update(&newsClass); err != nil {
				this.ReturnFailedJson(err, "更新资讯类别失败!")
			}
		}
		this.ReturnSuccessJson(newsClass)
	} else {
		this.Data["id"] ,_ = this.GetInt("id",0)
		this.Data["loadDataAction"] = this.getLoadAction()
		this.display()
	}
}

/**
 加载数据
 */
func(this *NewsClassController)	LoadData(){
	Id, _ := this.GetInt("id", 0)
	newsClass := &news.NewsClass{}
	var err error
	if Id > 0 {
		err = newsClass.FindById(Id)
		if err != nil {
			this.ReturnFailedJson(err,"加载数据错误")
		}
		this.ReturnSuccessJson(newsClass)
	} else {
		this.ReturnFailedJson(err,"Id不能为空")
	}
}