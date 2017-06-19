package models

import (
	//"errors"
	//"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	ConfInis map[string]*ConfIni
)

type ConfIni struct {
	Id          int       `orm:"column(ini_id);auto"`
	Project     string    `orm:"column(project);size(80)"`
	Section     string    `orm:"column(section);size(50)"`
	Key         string    `orm:"column(key);size(50)"`
	Value_ts    string    `orm:"column(value_ts);size(2000);kkkk"`
	Value_cs    string    `orm:"column(value_cs);size(2000);kkkk"`
	Value_yf    string    `orm:"column(value_yf);size(2000);kkkk"`
	Value_zs    string    `orm:"column(value_zs);size(2000);kkkk"`
	Encryption  string    `orm:"column(encryption);size(2);0"`
	Notes       string    `orm:"column(notes);size(300);kkkk"`
	Update_time time.Time `orm:"column(update_time);type(datetime)"`
}

//定义数据表名
func (t *ConfIni) TableName() string {
	return "conf_ini"
}

func init() {
	orm.RegisterModelWithPrefix("cmdb_", new(ConfIni))
}

func GetAll() map[string]*ConfIni {
	return ConfInis
}

func (this *ConfIni) Getproject(Pro string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	//读操作用只读库
	o.Using("read")
	c := new(ConfIni)
	//var maps []orm.Params
	if _, err = o.QueryTable(c).Filter("project", Pro).Values(&maps); err != nil {
		return
	}
	return
}

func (this *ConfIni) InsertProject(confIni *ConfIni) (created bool, err error) {
	o := orm.NewOrm()
	//开启事务
	o.Begin()
	if created, _, err = o.ReadOrCreate(confIni, "Project", "Section", "Key"); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

//通过ID删除
func (this *ConfIni) DeleteProject(Id int) (err error) {
	c := new(ConfIni)
	_, err = orm.NewOrm().QueryTable(c).Filter("Id", Id).Delete()
	return
}

//通过自增ID更新value
func (this *ConfIni) PatchProject(Id int, pk string, pv string) (num int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	c := new(ConfIni)
	if num, err = o.QueryTable(c).Filter("Id", Id).Update(orm.Params{pk: pv,
		"Update_time": time.Now()}); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

func (this *ConfIni) GetAllProjectName() (info_list []string, err error) {
	o := orm.NewOrm()
	o.Using("read")
	c := new(ConfIni)
	confs := []*ConfIni{}
	if _, err = o.QueryTable(c).All(&confs, "Project"); err == nil {
		for _, p := range confs {
			info_list = append(info_list, p.Project)
		}
		return
	}
	return
}
