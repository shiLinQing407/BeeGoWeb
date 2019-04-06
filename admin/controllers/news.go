/*
@Time : 2019/4/1 23:39 
@Author : shilinqing
@File : news
*/
package controllers

import (
	"BeeGoWeb/models/news"
	"BeeGoWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type NewsController struct {
	BaseController
}

/*资讯列表*/
func (this *NewsController) Index(){
	this.display()
}

func(this *NewsController) LoadList(){
	keyword := this.GetString("js_keyword", "")
	startTime := this.GetString("startTime", "")
	endTime := this.GetString("endTime", "")
	filter := make(map[string]interface{})
	if keyword != ""{
		filter["title__icontains"] = keyword //模糊查询 title
	}
	if startTime != "" {
		filter["create_time__gt"] = utils.TimeParseInLocation(startTime)
	}
	if endTime != "" {
		filter["create_time__lt"] = utils.TimeParseInLocation(endTime)
	}
	data, count := news.NewsListGrid(this.page, this.pageSize, filter)
	this.toDataGrid(data, count)
}

func(this *NewsController) Edit(){
	if this.isPost(){
		news := news.News{}
		model := orm.NewOrm()
		var err error
		//获取form里的值并赋值给结构体， name必须和结构对应
		if err = this.ParseForm(&news); err != nil {
			this.ReturnFailedJson(err, "获取数据失败")
		}
		fmt.Println(news)
		news.CreateTime = time.Now().Unix()
		if news.Id == 0{
			news.CreateTime = time.Now().Unix()
			if _, err := model.Insert(&news); err != nil {
				beego.Error(err.Error())
				this.ReturnFailedJson(err, "添加资讯失败")
			}
		}else{
			if _, err := model.Update(&news); err != nil {
				this.ReturnFailedJson(err, "更新资讯失败")
			}
		}
		this.ReturnSuccessJson(news)
	}else{
		this.Data["id"] ,_ = this.GetInt("id",0)
		this.Data["loadDataAction"] = this.getLoadAction()
		this.display()
	}
}

/**
 加载数据
 */
func(this *NewsController)	LoadData(){
	Id, _ := this.GetInt("id", 0)
	news := &news.News{}
	var err error
	if Id > 0 {
		news, err = news.FindById(Id)
		if err != nil {
			this.ReturnFailedJson(err,"加载数据错误")
		}
		this.ReturnSuccessJson(news)
	} else {
		this.ReturnFailedJson(err,"Id不能为空")
	}
}