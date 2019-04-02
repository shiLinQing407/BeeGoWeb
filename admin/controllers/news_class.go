/*
@Time : 2019/4/2 10:32 
@Author : shilinqing
@File : news_class
*/
package controllers

type NewsClassController struct {
	BaseController
}

func(this *NewsClassController) Index(){
	this.display()
}

func(this *NewsClassController) LoadList(){

}