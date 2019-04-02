/*
@Time : 2019/4/1 23:39 
@Author : shilinqing
@File : news
*/
package controllers

import "BeeGoWeb/models/news"

type NewsController struct {
	BaseController
}

/*资讯列表*/
func (this *NewsController) Index(){
	this.display()
}

func(this *NewsController) LoadList(){
	this.page, _ = this.GetInt("page")
	this.pageSize, _ = this.GetInt("limit")
	list := news.NewsListGrid(this.page, this.pageSize)
	this.jsonResult(0, "", list)
}