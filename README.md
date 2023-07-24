# bytetrack_server

## 简介

此项目是使用**gin**和**gorm**的gen模式实现对监控程序用户数据的管理。

前端部分：[CodeRikka/Bytetrack_Server (github.com)](https://github.com/CodeRikka/Bytetrack_Server)

## 主要功能

- 设备信息的创建/修改/删除/查询
- 用户的注册/登录
- (Todo) 监控到的异常行为的记录

## 项目结构

```
bytetrack_server/
├─ api
├─ build
├─ logs
├─ repository
│  ├─ dal
│  ├─ dao
│  └─ model
├─ routes
├─ service
└─ util
    ├─ applog
    └─ response
```

- api: controller层
- build: 使用gen模式生成model
- logs: 日志
- repository: 存储仓库
- repository/dal: gen模式生成出的接口
- repository/dao: dao层
- repository/model: 生成出的数据库表model层
- routes: 路由逻辑
- service: 接口具体实现
- util: 辅助工具
- util/applog: 日志生成工具
- util/response: 结构化response
