# 配置管理
## ini配置文件管理
### 表结构

| Field       | Type           | Null | Key  | Default | Extra           | 
| -------     | ----           | ---- | ---- | ------- | -----           |
| id          | int(11)        | NO   | PRI  |         | auto_increment  |
| project     | string(80)     | NO   |      | NULL    |                 |
| section     | string(50)     | NO   |      | NULL    |                 |
| key         | string(50)     | NO   |      | NULL    |                 |
| value_ts    | varchar(2000)  | YES  |      | NULL    |                 |
| value_cs    | varchar(2000)  | YES  |      | NULL    |                 |
| value_yf    | varchar(2000)  | YES  |      | NULL    |                 |
| value_zs    | varchar(2000)  | YES  |      | NULL    |                 |
| encryption  | varchar(2)     | NO   |      | 0       |                 |
| notes       | string(1000)   | NO   |      | NULL    |                 |
| update_time | datetime       | NO   |      | NULL    |                 |
###
### API访问
- 获取配置  
  - 接口地址：`/v1/conf/ini` 
  - 请求方式：GET

###
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 |
| environment | string | 是       | 环境名称 | 

示例:  
url:/v1/conf/ini/exam_server/?huanjing=cs
```json
{
  "env": "cs",
}
```

- 更新配置
 - 接口地址 `/v1/conf/ini`  
 - 请求方式：PATCH

| 参数名称    | 类型   | 是否必要 | 备注     |
| --         | --     | --       | --      |
| value_ts   | string | 否       | 调试值   |
| value_cs   | string | 否       | 测试值   | 
| value_yf   | string | 否       | 预发布   |
| value_zs   | string | 否       | 正式值   |

示例:  
url: /v1/conf/ini?pid=35&value_cs=80 
```json
{
  "id": "30",
  "value_cs": "80"
}
```

- 删除配置
 - 接口地址 `/v1/conf/ini`
 - 请求方式：DELETE

| 参数名称    | 类型   | 是否必要  | 备注     |
| --         | --     | --       | --       |
| id         | string | 是       | 主键ID |
  

示例:    
```json
{
  "id": "30",
}
```

- 新建配置
 - 接口地址 `/v1/conf/ini`
 - 请求方式：POST

| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 | 
| section     | string | 是       | 节       |
| key         | string | 是       | 键       |
| value_ts    | string | 否       | 值       |
| value_cs    | string | 否       | 值       |
| value_yf    | string | 否       | 值       |
| value_zs    | string | 否       | 值       |
| encryption  | int    | 是       | 加密类型 |

示例:  
url: /v1/conf/ini?project=exam_server07&section=default&key=service_sso2_port&value_cs=80&encryption=0 
```json
{
  "project": "exam_server",
  "section": "default",
  "key": "app_name",
  "value_cs": "exam"
  "encryption": "0",
}
```
###安装
go get github.com/astaxie/beego
go get github.com/go-sql-driver/mysql
go get github.com/astaxie/beego/session/redis