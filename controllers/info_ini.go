package controllers

import (
	//"encoding/json"
	"fmt"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"ss_cmdb/common/consts"
	"ss_cmdb/common/utils"
	"ss_cmdb/models"
	"strconv"
	"time"
)

type InfoIniController struct {
	beego.Controller
}

func (this *InfoIniController) Prepare() {

}

// @Title Get
// @Description find project by Conf_ini
// @Param	project		path 	string	true		"the project you want to get"
// @Success 200 {project} models.Conf_ini
// @Failure 403 :project is empty
// @router /:project [get]
func (this *InfoIniController) Get() {
	type ProInfo struct {
		Section     string    `json:"section"`
		Key         string    `json:"key"`
		Value_ts    string    `json:"ts"`
		Value_cs    string    `json:"cs"`
		Value_yf    string    `json:"yf"`
		Value_zs    string    `json:"zs"`
		Encryption  string    `json:"encryption"`
		Notes       string    `json:"notes"`
		Update_time time.Time `json:"update_time"`
	}
	//获取输入参数
	project := this.GetString(":project")
	var (
		conf                = new(models.ConfIni)
		ProInfos []*ProInfo = []*ProInfo{}
	)

	if get_info, err := conf.Getproject(project); err != nil {
		this.Data["json"] = map[string]interface{}{
			"status": consts.ErraMsg,
			"msg":    consts.ERROR_MSG_CMDB_GETINFO,
		}
		this.ServeJSON()
		return
	} else {

		for _, m := range get_info {
			ProInfos = append(ProInfos, &ProInfo{
				Section:     m["Section"].(string),
				Key:         m["Key"].(string),
				Value_ts:    m["Value_ts"].(string),
				Value_cs:    m["Value_cs"].(string),
				Value_yf:    m["Value_yf"].(string),
				Value_zs:    m["Value_zs"].(string),
				Encryption:  m["Encryption"].(string),
				Notes:       m["Notes"].(string),
				Update_time: m["Update_time"].(time.Time),
			})
		}

		this.Data["json"] = map[string]interface{}{
			project:  ProInfos,
			"status": consts.SuccCode,
			"msg":    consts.SUCC_MSG_CMDB_GETINFO,
		}
		this.ServeJSON()
		return
	}
}

//定义表单结构
type ProInfoForm struct {
	Project    string `form:"project"`
	Section    string `form:"section"`
	Key        string `form:"key"`
	Env        string `form:"environment"`
	Value_ts   string `form:"value_ts"`
	Value_cs   string `form:"value_cs"`
	Value_yf   string `form:"value_yf"`
	Value_zs   string `form:"value_zs"`
	Encryption string `form:"encryption"`
	Notes      string `form:"notes"`
}

var confInfo *ProInfoForm = new(ProInfoForm)

// 新增配置项
// @router / [post]
func (this *InfoIniController) Post() {
	var conf = new(models.ConfIni)
	//
	if err := this.ParseForm(confInfo); err != nil {
		this.Data["json"] = map[string]interface{}{
			"created": "",
			"status":  consts.ErraCode,
			"msg":     consts.ERROR_MSG_CMDB_PROJECT,
		}
		this.ServeJSON()
		return
	}
	//赋值
	confIni := &models.ConfIni{
		Project:     confInfo.Project,
		Section:     confInfo.Section,
		Key:         confInfo.Key,
		Value_ts:    confInfo.Value_ts,
		Value_cs:    confInfo.Value_cs,
		Value_yf:    confInfo.Value_yf,
		Value_zs:    confInfo.Value_zs,
		Encryption:  confInfo.Encryption,
		Notes:       confInfo.Notes,
		Update_time: time.Now(),
	}

	if created, err := conf.InsertProject(confIni); err != nil {
		this.Data["json"] = map[string]interface{}{
			"created": created,
			"status":  consts.ErraCode,
			"msg":     consts.ERROR_MSG_CMDB_INSERTINFO,
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{
			"created": created,
			"status":  consts.SuccCode,
			"msg":     consts.SUCC_MSG_CMDB_INSERTINFO,
		}
		this.ServeJSON()
	}
	return
}

// @router /:pid [delete]
func (this *InfoIniController) Delete() {
	pid := this.GetString(":pid")
	var conf = new(models.ConfIni)

	if Id, erro := strconv.Atoi(pid); erro != nil {
		this.Data["json"] = map[string]interface{}{
			"status": consts.ErraCode,
			"msg":    consts.ERROR_MSG_CMDB_ID,
		}
		this.ServeJSON()
		return

	} else {
		if err := conf.DeleteProject(Id); err == nil {
			this.Data["json"] = map[string]interface{}{
				"status": consts.SuccCode,
				"msg":    consts.SUCC_MSG_CMDB_DELETEINFO,
			}
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": consts.ErraCode,
		"msg":    consts.ERROR_MSG_CMDB_DELETEINFO,
	}
	this.ServeJSON()
	return
}

// @router / [patch]
func (this *InfoIniController) Patch() {
	pid := this.GetString("pid")
	this.ParseForm(confInfo)
	var (
		Patch_k = ""
		Patch_v = ""
		conf    = new(models.ConfIni)
	)

	switch {
	case confInfo.Value_ts != "":
		Patch_k = "value_ts"
		Patch_v = confInfo.Value_ts
	case confInfo.Value_cs != "":
		Patch_k = "value_cs"
		Patch_v = confInfo.Value_cs
	case confInfo.Value_yf != "":
		Patch_k = "value_yf"
		Patch_v = confInfo.Value_yf
	case confInfo.Value_zs != "":
		Patch_k = "value_zs"
		Patch_v = confInfo.Value_yf
	default:
		this.Data["json"] = map[string]interface{}{
			"status": consts.ErraCode,
			"msg":    consts.ERROR_MSG_CMDB_UPDATE,
		}
		this.ServeJSON()
		return
	}

	if Id, erro := strconv.Atoi(pid); erro != nil {
		this.Data["json"] = map[string]interface{}{
			"status": consts.ErraCode,
			"msg":    consts.ERROR_MSG_CMDB_ID,
		}
		this.ServeJSON()
		return
	} else {

		if num, err := conf.PatchProject(Id, Patch_k, Patch_v); err == nil {
			fmt.Println("xxxx", num, err)
			this.Data["json"] = map[string]interface{}{
				"status": consts.SuccCode,
				"num":    num,
				"msg":    consts.SUCC_MSG_CMDB_UPDATEINFO,
			}
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = map[string]interface{}{
				"status": consts.ErraCode,
				"msg":    consts.ERROR_MSG_CMDB_UPDATEINFO,
			}
			this.ServeJSON()
			return
		}
	}
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *InfoIniController) GetAll() {
	var conf = new(models.ConfIni)

	if lists, err := conf.GetAllProjectName(); err == nil {
		result_list := utils.RemoveDuplicatesAndEmpty(lists)
		this.Data["json"] = map[string]interface{}{
			"project_list": result_list,
			"status":       consts.SuccCode,
			"msg":          consts.SUCC_MSG_CMDB_GET_ALL,
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"status": consts.ErraCode,
		"msg":    consts.ERROR_MSG_CMDB_GET_ALL,
	}
	this.ServeJSON()
	return
}
