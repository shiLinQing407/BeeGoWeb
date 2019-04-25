package utils

import (
	"github.com/astaxie/beego"
	"mime/multipart"
	"errors"
	"os"
	"path"
	"io"
	"strings"
	"tests/utils"
	"time"
)

type Upload struct {
	Mimes        map[string]interface{}      //允许上传的文件MiMe类型
	MaxSize      int64                       //上传的文件大小限制 (0-不做限制)
	Exts         map[string]string           //允许上传的文件后缀
	AutoSub      bool                        //自动子目录保存文件
	TimeSub      bool                        //保存文件划分时间目录
	SubName      map[int]string              //子目录创建方式，[0]-函数名，[1]-参数，多个参数使用数组  'subName' => array('date', 'Ymd'),
	RootPath     string                      // BASE_DIR . '/public/uploads/', //保存根路径
	SavePath     string                      //保存路径
	SaveName     string                      //上传文件命名规则，[0]-函数名，[1]-参数，多个参数使用数组
	SaveExt      string                      //文件保存后缀，空则使用原后缀
	Replace      bool                        //存在同名是否覆盖
	Hash         bool                        //是否生成hash编码
	CallBack     bool                        //检测文件是否存在回调，如果存在返回文件信息数组
	Driver       string                      // 文件上传驱动
	DriverConfig map[interface{}]interface{} // 上传驱动配置
	error        error
	FileType     string //文件类型  png、jpg ...
}

/**
 *	对实例化的结构体进行初始化
 *  默认配置   利用反射给结构体赋值
 */
func (u *Upload) Construct(params map[string]interface{}) error {
	//初始化文件上传配置
	u.MaxSize = 0          //上传的文件大小限制 (0-不做限制)
	u.AutoSub = true       //自动子目录保存文件
	u.RootPath = "/upload" // BASE_DIR . '/public/uploads/', //保存根路径
	u.SavePath = ""        //保存路径
	u.SaveExt = ""         //文件保存后缀，空则使用原后缀
	u.Replace = false      //存在同名是否覆盖
	u.Hash = true          //是否生成hash编码
	u.CallBack = false     //检测文件是否存在回调，如果存在返回文件信息数组
	u.Driver = ""          // 文件上传驱动
	u.FileType = ".gif,.jpg,.jpeg,.bmp,.png,.swf,.tmp"
	u.SubName = make(map[int]string) //子目录创建方式，[0]-函数名，[1]-参数，多个参数使用数组  'subName' => array('date', 'Ymd'),
	u.SubName[0] = "date"
	u.SubName[1] = "Ymd"

	if len(params) > 0 {
		//根据传入的参数初始化upload结构体
		if err := InitStruct(u, params); err == nil {
			return nil
		} else {
			return err
		}
	}
	return nil
}

/**
 * 创建目录
 * @param  string $savepath 要创建的目录
 * @return boolean          创建状态，true-成功，false-失败
 */
func (u *Upload) Mkdir(savePath string) bool {
	dir := u.RootPath + savePath
	if !IsDir(dir) {
		if err := os.Mkdir(savePath, os.ModePerm); err != nil {
			u.error = errors.New("上传目录" + savePath + "创建失败!")
			return false
		}
	}
	return true
}

/**
 * 获取最后一次上传错误信息
 * @return string 错误信息
 */
func (u *Upload) GetError() string {
	return u.error.Error()
}

/**
 * 保存指定文件
 * @param  array   $file    保存的文件信息
 * @param  boolean $replace 同名文件是否覆盖
 * @return boolean          保存状态，true-成功，false-失败
 */
func (u *Upload) saveFile(file *multipart.FileHeader, fileInfo map[string]interface{}, replace bool) bool {
	temFile, err := file.Open()
	defer temFile.Close()
	if err != nil {
		u.error = errors.New("上传文件获取失败")
		return false
	} else {
		fileName := String(fileInfo["src"])
		if !replace && IsFile(fileName) {
			/* 不覆盖同名文件 */
			//u.error = errors.New("存在同名文件" + String(fileInfo["title"]))
			//return false
			return true
		} else if replace {
			/* 覆盖同名文件 */
			beego.Debug("覆盖")

		} else {
			//新上传文件
			f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
			defer f.Close()
			if err != nil {
				u.error = errors.New("create file error:" + err.Error())
				return false
			}
			if _, err := io.Copy(f, temFile); err != nil {
				u.error = errors.New("文件上传失败! 文件名：" + file.Filename)
				return false
			}
		}
		return true
	}
}

/**
 * 检查文件大小是否合法
 * @param integer $size 数据
 */
func (u *Upload) CheckSize(size int64) bool {
	return size <= u.MaxSize
}

/**
 * 检测上传根目录 是否存在且可写
 * @return boolean true-检测通过，false-检测失败
 */
func (u *Upload) CheckRootPath() bool {
	if !(IsDir(u.RootPath) && IsWritable(u.RootPath)) {
		u.error = errors.New("上传根目录不存在或无可写权限,请检查后重新尝试！" + u.RootPath)
		return false
	} else {
		return true
	}
}

/**
 * 检测上传目录
 * @param  string $savepath 上传目录
 * @return boolean          检测结果，true-通过，false-失败
 */
func (u *Upload) CheckSavePath() bool {
	path := u.RootPath + u.SavePath
	if !IsExist(path) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			u.error = errors.New("上传目录" + path + "创建失败!")
			return false
		}
	} else if !IsWritable(path) {
		u.error = errors.New("上传目录" + u.SavePath + "不可写!")
		return false
	}
	return true
}

/**
 * 检查上传的文件MIME类型是否合法
 * @param string $mime 数据
 */
//private function checkMime($mime)
//{
//return empty($this->config['mimes']) ? true : in_array(strtolower($mime), $this->mimes);
//}

/**
 * 检查上传的文件后缀是否合法
 * @param string $ext 后缀
 */
func (u *Upload) CheckExt(ext string) bool {
	if Empty(u.FileType) {
		return true
	} else if InArray(Explode(",", u.FileType), strings.ToLower(ext)) {
		return true
	} else {
		return false
	}
}

/**
 * 根据指定的规则获取文件或目录名称
 * @param  array  $rule     规则
 * @param  string $filename 原文件名
 * @return string           文件或目录名称
 */
func (u *Upload) getName(saveName, fileName string) string {
	return saveName + fileName
}

//private function getName($rule, $filename)
//{
//$name = '';
//if (is_array($rule)) { //数组规则
//$func = $rule[0];
//$param = (array)$rule[1];
//foreach ($param as &$value) {
//$value = str_replace('__FILE__', $filename, $value);
//}
//if($func=='uuid'){
//$random=new \Phalcon\Security\Random();
//$func=[$random,'uuid'];
//}
//$name = call_user_func_array($func, $param);
//} elseif (is_string($rule)) { //字符串规则
//if (function_exists($rule)) {
//$name = call_user_func($rule);
//} else {
//$name = $rule;
//}
//}
//return $name;
//}

/**
 * 根据上传文件命名规则取得保存文件名
 * @param string $file 文件信息
 */
func (u *Upload) getSaveName(file *multipart.FileHeader) string {
	rule := u.SaveName
	var saveName string
	if Empty(rule) {
		saveName = file.Filename
	} else {
		saveName = u.getName(rule, "")
		if Empty(saveName) {
			u.error = errors.New("File name rule error!")
			return ""
		}
	}
	if Empty(u.SaveExt) {
		return saveName + path.Ext(file.Filename)
	} else {
		return saveName + u.SaveExt
	}
}

/**
 * 获取子目录的名称
 * @param array $file  上传的文件信息
 */
func (u *Upload) getSubPath(fileName string) string {
	var subPath string
	subPath = ""
	rule := u.SavePath
	if u.AutoSub && rule != "" {
		subPath = u.getName(rule, fileName) + "/"
	}
	if Empty(subPath) || !u.Mkdir(u.SavePath + subPath) {
		return ""
	}
	return subPath
}

func (u *Upload) getTimeSubPath() string {
	path := u.RootPath + u.SavePath
	var timeSubPath string
	timeSubPath = ""
	if u.TimeSub {
		timeSubPath = utils.Date("Ymd", time.Now()) + "/"
	}
	if Empty(timeSubPath){
		return ""
	} else if !IsExist(path + timeSubPath) && !u.Mkdir(path + timeSubPath)  {
		return ""
	}
	return timeSubPath
}

func (u *Upload) getHash(temName string) string {
	return Md5(temName)
}

func (u *Upload) upload(file *multipart.FileHeader) (map[string]interface{}, bool) {
	/* 检测上传根目录 */
	if !u.CheckRootPath() {
		return nil, false
	}

	/* 检查上传目录 */
	if !u.CheckSavePath() {
		return nil, false
	}

	var info = make(map[string]interface{})
	info["key"] = ""
	info["type"] = ""                                                              //文件类型
	info["ext"] = path.Ext(file.Filename)                                          //文件扩展
	info["name"] = strings.Replace(file.Filename, path.Ext(file.Filename), "", -1) //文件名称
	info["size"] = file.Size                                                       //文件大小

	/* 无效上传 */
	if Empty(file.Filename) {
		u.error = errors.New("无效上传!")
		return nil, false
	}

	/* 检查文件大小 */
	if (u.CheckSize(file.Size)) {
		u.error = errors.New("Upload file size inconsistent!")
		return nil, false
	}

	/* 检查是否合法上传 */
	//if (!$file->isUploadedFile()) {
	//	throw new \Exception('Illegal file upload!');
	//}

	/* 检查文件Mime类型 */
	//if (!$this->checkMime($info['type'])) {
	//throw new \Exception('Upload file MIME type not allowed!');
	//}

	/* 检查文件后缀 */
	if !u.CheckExt(String(info["ext"])) {
		u.error = errors.New("File suffix name error!'")
		return nil, false
	}

	/* 哈希验证 */
	if u.Hash {
		info["md5"] = u.getHash(u.SaveName)
	}

	/* 调用回调函数检测文件是否存在 */
	//if($this->callback){
	//	$data = call_user_func($this->callback, $info);
	//	if ($this->callback && $data) {
	//		if (file_exists('.' . $data['path'])) {
	//		return $data;
	//		} elseif ($this->removeTrash) {
	//			call_user_func($this->removeTrash, $data);//删除垃圾据
	//		}
	//	}
	//}

	/* 生成保存文件名 */
	info["title"] = u.getSaveName(file);

	/* 检测并创建子目录 */
	info["src"] = u.RootPath + u.SavePath + u.getSubPath(String(info["name"])) + u.getTimeSubPath() + String(info["title"])
	//
	/* 保存文件 并记录保存成功的文件 */
	if u.saveFile(file, info, u.Replace) {
		return info, true
	} else {
		return nil, false
	}
}

/*
	上传文件并返回文件存储路径等信息
*/
func (u *Upload) Upload(file *multipart.FileHeader) (map[string]interface{}, error) {
	if info, k := u.upload(file); !k {
		return nil, u.error
	} else {
		return info, nil
	}
}
