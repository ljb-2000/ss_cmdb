package controllers

import (
	//"encoding/json"
	"fmt"
	//"github.com/astaxie/beego/orm"
	"ss_cmdb/common/consts"
	"ss_cmdb/models"

	"github.com/astaxie/beego"
)

// Operations about Conf_ini
type Info_iniController struct {
	beego.Controller
}

func (this *Info_iniController) Prepare() {

}

type Error_info struct {
	Project     string `json:"project"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
	Msg         string `json:"msg"`
}

type Succ_info struct {
	Project     string `json:"project"`
	Value       string `json:"value "`
	Section     string `json:"section"`
	Key         string `json:"key"`
	Update_time string `json:"update_time"`
}

// @Title Get
// @Description find project by Conf_ini
// @Param	project		path 	string	true		"the project you want to get"
// @Success 200 {project} models.Conf_ini
// @Failure 403 :project is empty
// @router / [get]
func (this *Info_iniController) Get() {

	//fmt.Println("thi is ini get", this.Ctx.Input)
	//获取输入参数
	project := this.GetString("project")
	environment := this.GetString("huanjing")

	fmt.Println(project, environment)
	//返回格式
	var result = Error_info{Status: consts.SuccCode, Msg: consts.SuccMsg, Project: project, Environment: environment}

	//判断必填参数
	if project == "" {
		result.Status = consts.ERROR_CODE_CMDB_PROJECT
		result.Msg = consts.ERROR_MSG_CMDB_PROJECT
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if environment == "" {
		result.Status = consts.ERROR_CODE_CMDB_ENVIRONMENT
		result.Msg = consts.ERROR_MSG_CMDB_ENVIRONMENT
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	conf := new(models.Conf_ini)
	pro_info, err := conf.Getproject(project)

	for _, m := range pro_info {
		fmt.Println(m["Project"], m["Section"], m["Key"], m[environment], m["Envir_cs"])
	}

	fmt.Println(pro_info, err)
	this.Data["json"] = "{\"Oproject\":\"" + project + "\"}"
	this.ServeJSON()
	return
}

// @Title GetAll
// @Description get all project
// @Success 200 {project} models.Conf_ini
// @Failure 403 :project is empty
// @router / [get]
func (i *Info_iniController) GetAll() {
	obs := models.GetAll()
	i.Data["json"] = obs
	fmt.Println("this is ini controllers", obs)
	i.ServeJSON()
}
