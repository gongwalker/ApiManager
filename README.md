# 项目介绍

## 什么是接口文档管理工具?
采用golang 基于Gin框架编写的在线API文档管理系统；其致力于快速解决团队内部接口文档的编写、维护、存档，以及减少团队协作开发的沟通成本。

## 在线体验地址
体验地址 [http://apidoc.gwalker.cn/](http://apidoc.gwalker.cn/)
> 账号: guest 密码: 123456 [注: 此体验账号为游客模式,只有只读权限]
> 不要改密码哟 :)

## 安装方法
### 方法一 直接使用 (推荐)
* 进入 https://github.com/gongwalker/ApiManager/releases

```
根据您的系统环境，下载 相应系统环境的可执行文件压缩包
linux ApiManager-linux-amd64.zip
windows ApiManager-windows-amd64.zip
mac ApiManager-darwin-amd64.zip
```
* 进入解压目录，设置数据库(创建数据库，导入db.sql) 与 配置文件(config/config.conf)
* 运行 ./run.sh start|stop
* 浏览器打开 http://your_host:8000 (账号: root 密码: 123456)


### 方法二 编译安装
* 在MySQL中创建mysql数据库，db.sql导入
* 修改 /config/config.ini 配置文件，修改数据库连接信息
* git clone git@github.com:gongwalker/ApiManager.git
  或者
  git clone git@gitee.com:gongwalker/ApiManager.git
* go mod vendor
* go build
* 运行 ./run.sh start|stop
* 浏览器打开 http://your_host:8000 (账号: root 密码: 123456)


### 使用说明

1. 系统有三个角色,超级管理、管理员、游客
    - **超级管理员** 拥有一切权限(api分类管理,api管理)
    - **管理** 可创建编辑API
    - **游客** 只能查看接口分类和接口信息 __无增删改权限__
    
2. 默认的超级管理员 账号root 密码:123456

### 用到了哪些技术及项目价值
1. 前端页面采用 layui-v2.5.6,Bootstrap v3.3.4，后端采用gin框架编写Restful Api 接口,前后端分离。
2. 后端用到了表单校验，权限校验中间件定义,mysql数据存档，项目session可以支持cookie与redis两个储存方案
3. 适用于中小团队api文档管理使用
4. 可以作为一个基础角手架进行使用，快速开发。初学者可以作为熟悉gin框架学习所用



### 系统长什么样
![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/1.jpg?raw=true)

![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/2.jpg?raw=true)

![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/3.jpg?raw=true)

![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/4.jpg?raw=true)

![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/5.jpg?raw=true)



### 作者信息
* Email:	gongcoder@gmail.com
* Blog:	[https://www.gwalker.cn](https://www.gwalker.cn)


### 写在最后
非常欢迎大家贡献代码，让这个项目成长的更好。

![image](https://github.com/gongwalker/ApiManager/blob/master/app/static/show/me.jpg?raw=true)
