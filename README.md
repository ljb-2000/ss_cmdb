# 配置管理
## ini配置文件管理
### 表结构

| Field       | Type          | Null | Key  | Default | Extra           | 
| -------     | ----          | ---- | ---- | ------- | -----           |
| id          | int(11)       | NO   | PRI  |         | auto_increment  |
| project     | string(80)    | NO   |      | NULL    |                 |
| section     | string(50)    | NO   |      | NULL    |                 |
| key         | string(50)    | NO   |      | NULL    |                 |
| envir_ts    | string(1500)  | YES  |      | NULL    |                 |
| envir_cs    | string(1500)  | YES  |      | NULL    |                 |
| envir_zs    | string(1500)  | YES  |      | NULL    |                 |
| encryption  | int(2)        | NO   |      | 0       |                 |
| notes       | string(1000)  | NO   |      | NULL    |                 |
| update_time | datetime      | NO   |      | NULL    |                 |
###
### API访问
- 获取配置  
  - 接口地址：`/v1/conf/info_ini` 
  - 请求方式：GET

###
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 |
| environment | string | 是       | 环境名称 | 
| section     | string | 否       | 节       |
| key         | string | 否       | 键       |

示例:    
```json
{
  "project": "exam_server",
  "environment": "envir_cs",
  "section": "default",
  "key": "app_name",
}
```

- 更新配置
 - 接口地址 `/v1/conf/info_ini` 
 - 请求方式：PATCH

| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 |
| environment | string | 否       | 环境名称 | 
| section     | string | 否       | 节       |
| key         | string | 否       | 键       |
| encryption  | int    | 否       | 加密类型 |

示例:    
```json
{
  "project": "exam_server",
  "environment": "envir_cs",
  "section": "default",
  "key": "app_name",
  "encryption": 0,
}
```

- 删除配置
 - 接口地址 `/v1/conf/info_int`
 - 请求方式：DELETE

| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 |
| environment | string | 否       | 环境名称 | 
| section     | string | 否       | 节       |
| key         | string | 否       | 键       |

示例:    
```json
{
  "project": "exam_server",
  "environment": "envir_cs",
  "section": "default",
  "key": "app_name",
}
```

- 新建配置
 - 接口地址 `/v1/conf/info_ini`
 - 请求方式：POST

| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| project     | string | 是       | 项目名称 |
| environment | string | 否       | 环境名称 | 
| section     | string | 否       | 节       |
| key         | string | 否       | 键       |
| encryption  | int    | 否       | 加密类型 |

示例:    
```json
{
  "project": "exam_server",
  "environment": "envir_cs",
  "section": "default",
  "key": "app_name",
  "encryption": 0,
}
```
