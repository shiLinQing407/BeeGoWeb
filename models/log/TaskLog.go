package log

import (
	"github.com/astaxie/beego/orm"
	"BeeGoWeb/models/common"
)

type TaskLog struct {
	Id          int
	TaskId      int
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}
//表名
var taskLogTable = common.TableName("task_log")


func (t *TaskLog) GetTableName() string {
	return taskLogTable
}

func TaskLogAdd(t *TaskLog) (int64, error) {
	return orm.NewOrm().Insert(t)
}

func TaskLogGetList(page, pageSize int, filters ...interface{}) ([]*TaskLog, int64) {
	offset := (page - 1) * pageSize

	logs := make([]*TaskLog, 0)

	query := orm.NewOrm().QueryTable(taskLogTable)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&logs)

	return logs, total
}

func TaskLogGetById(id int) (*TaskLog, error) {
	obj := &TaskLog{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func TaskLogDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(taskLogTable).Filter("id", id).Delete()
	return err
}

func TaskLogDelByTaskId(taskId int) (int64, error) {
	return orm.NewOrm().QueryTable(taskLogTable).Filter("task_id", taskId).Delete()
}
