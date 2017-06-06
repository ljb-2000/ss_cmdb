package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	//"ss_cmdb/common/consts"
)

var (
	Conf_inis map[string]*Conf_ini
)

type Conf_ini struct {
	Id          int       `orm:"column(conf_ini_id);auto"`
	Project     string    `orm:"column(project);size(80);null"`
	Section     string    `orm:"column(section);size(50);null"`
	Key         string    `orm:"column(key);size(50);null"`
	Envir_ts    string    `orm:"column(envir_ts);size(1500);null"`
	Envir_cs    string    `orm:"column(envir_cs);size(1500);null"`
	Envir_zs    string    `orm:"column(envir_zs);size(1500);null"`
	Encryption  string    `orm:"column(encryption);size(2);0"`
	Notes       string    `orm:"column(notes);size(300);null"`
	Update_time time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *Conf_ini) TableName() string {
	return "conf_ini"
}

func init() {
	orm.RegisterModelWithPrefix("cmdb_", new(Conf_ini))
}

func GetAll() map[string]*Conf_ini {
	return Conf_inis
}

func (this *Conf_ini) Getproject(Pro string) ([]orm.Params, error) {

	o := orm.NewOrm()
	//读操作用只读库
	o.Using("read")
	c := new(Conf_ini)

	var maps []orm.Params
	num, err := o.QueryTable(c).Filter("project", Pro).Values(&maps)

	if err != nil {
		beego.Debug(num)
		return maps, err
	}
	return maps, err

}

func GetOne(Id string) (conf *Conf_ini, err error) {
	if v, ok := Conf_inis[Id]; ok {
		return v, nil
	}
	return nil, errors.New("Id Not Exist")
}

func Delete(Id string) {
	delete(Conf_inis, Id)
}
