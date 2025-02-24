#### 描述

该接口提供版本：v1.0.0+

查询模版被套餐引用详情

#### 输入参数

| 参数名称          | 参数类型 | 必选 | 描述       |
| ----------------- | -------- | ---- | ---------- |
| biz_id            | uint32   | 是   | 业务ID     |
| template_space_id | uint32   | 是   | 模版空间ID |
| template_id       | uint32   | 是   | 模版ID     |
| start             | uint32   | 是   | 分页起始值 |
| limit             | uint32   | 是   | 分页大小   |

#### 调用示例

```json

```

#### 响应示例

```json
{
  "data": {
    "count": 2,
    "details": [
      {
        "template_set_id": 1,
        "template_set_name": "template_set001"
      },
      {
        "template_set_id": 2,
        "template_set_name": "template_set001"
      }
    ]
  }
}
```

#### 响应参数说明

| 参数名称 | 参数类型 | 描述     |
| -------- | -------- | -------- |
| data     | object   | 响应数据 |

#### data

| 参数名称 | 参数类型 | 描述                         |
| -------- | -------- | ---------------------------- |
| count    | uint32   | 当前规则能匹配到的总记录条数 |
| detail   | array    | 查询返回的数据               |

#### data.detail[n]

| 参数名称          | 参数类型 | 描述         |
| ----------------- | -------- | ------------ |
| template_set_id   | uint32   | 模版套餐ID   |
| template_set_name | uint32   | 模版套餐名称 |

