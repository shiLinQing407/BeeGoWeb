/*
@Time : 2019/4/1 23:39 
@Author : shilinqing
@File : news
*/
package controllers

import (
	"BeeGoWeb/models/news"
	"BeeGoWeb/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"path"
	"strings"
	"BeeGoWeb/enums"
)

type NewsController struct {
	BaseController
}

/*资讯列表*/
func (this *NewsController) Index() {
	this.display()
}

func (this *NewsController) LoadList() {
	keyword := this.GetString("js_keyword", "")
	startTime := this.GetString("startTime", "")
	endTime := this.GetString("endTime", "")
	filter := make(map[string]interface{})
	if keyword != "" {
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

func (this *NewsController) Edit() {
	if this.isPost() {
		newsClassId, _ := this.GetInt("NewsClassId", 0)
		new := news.News{}
		model := orm.NewOrm()
		var err error
		//获取form里的值并赋值给结构体， name必须和结构对应
		if err = this.ParseForm(&new); err != nil || newsClassId == 0 {
			this.ReturnFailedJson(err, "获取数据失败!")
		}
		new.NewsClass = &news.NewsClass{Id: newsClassId}
		new.CreateTime = time.Now().Unix()
		if new.Id == 0 {
			new.CreateTime = time.Now().Unix()
			if _, err := model.Insert(&new); err != nil {
				beego.Error(err.Error())
				this.ReturnFailedJson(err, "添加资讯失败")
			}
		} else {
			if _, err := model.Update(&new); err != nil {
				this.ReturnFailedJson(err, "更新资讯失败")
			}
		}
		this.ReturnSuccessJson(new)
	} else {
		this.Data["id"], _ = this.GetInt("id", 0)
		this.Data["loadDataAction"] = this.getLoadAction()
		this.display()
	}
}

/**
 加载数据
 */
func (this *NewsController) LoadData() {
	Id, _ := this.GetInt("id", 0)
	news := &news.News{}
	var err error
	if Id > 0 {
		err = news.FindById(Id)
		if err != nil {
			this.ReturnFailedJson(err, "加载数据错误")
		}
		this.ReturnSuccessJson(news)
	} else {
		this.ReturnFailedJson(err, "Id不能为空")
	}
}

func (this *NewsController) UploadImage() {
	file, fileHead, fileErr := this.Ctx.Request.FormFile("file") //上传的文件
	if file == nil || fileErr != nil {
		beego.Error(" 未找到要上传的文件!, ERROR: FormFile获取上传文件失败!")
		this.jsonResult(enums.JRCodeFailed, "未找到要上传的文件!", nil)
	}
	uploadConf := make(map[string]interface{})
	uploadConf["RootPath"] = "resource/" //图库根路径
	uploadConf["SavePath"] = "news/" //分块缓存文件存储路径
	uploadConf["AutoSub"] = false
	uploadConf["TimeSub"] = true  //图片保存按日期分目录
	uploadConf["SaveName"] = strings.Replace(fileHead.Filename, path.Ext(fileHead.Filename),"", -1) //文件名称 不带后缀
	uploadConf["SaveExt"] = path.Ext(fileHead.Filename) //文件后缀

	var upload = utils.Upload{}
	if err := upload.Construct(uploadConf); err == nil {
		if info, err := upload.Upload(fileHead); err != nil {
			beego.Error(err.Error())
			this.jsonResult(enums.JRCodeFailed, fileHead.Filename + "文件上传失败! error message:" + err.Error(), nil)
		} else {
			this.jsonResult(enums.JRCodeSucc, "", info)
		}
	}
}
