# Cooooooool-Blog

## 项目介绍
个人自用的Go博客系统的后端服务,采用gin框架+mysql数据库构建,目前正在实现中。

| 引用框架                     | 项目地址                                               |
| ---------------------------- | ------------------------------------------------------ |
| WEB框架: **Gin**             | **Gin**: https://gin-gonic.com/                        |
| ORM框架: **Gorm**            | **Gorm**: https://gorm.io/                             |
| SQL数据库: **Mysql8.0**      | **Mysql**: https://www.mysql.com/                      |
| 用户认证: **Json Web Token**     | **Jwt-go**: https://github.com/golang-jwt/jwt          |
| 验证码系统: **base64Captcha** | **base64Captcha**: https://github.com/mojocn/base64Captcha |
| api文档生成: **gin-swagger** | **gin-swagger**: https://github.com/swaggo/gin-swagger |
| 配置文件解析: **viper**        | **viper**: https://github.com/spf13/viper              |

## 填写配置文件
``` yaml
# configs/config.yaml
Server:
RunMode: debug
HttpPort: 8080
ReadTimeout: 60
WriteTimeout: 60
App:
DefaultPageSize: 10
MaxPageSize: 100
LogSavePath: storage/logs
LogFileName: app
LogFileExt: .log
Database:
DBType: mysql
Username: username  # 填写你的数据库账号
Password: password  # 填写你的数据库密码
Host: 127.0.0.1:3306
DBName: blog
TablePrefix:
Charset: utf8
ParseTime: True
LogMode: info
MaxIdleConns: 10
MaxOpenConns: 30
Jwt:
Secret: Secret# 密钥
Issuer: blog.nmslwsnd.com #签发人
Expire: 72000
```

## 启动项目
```shell
go run main.go
```

### 功能列表

| 已实现功能       | 未实现功能                   |
| ---------------- | ---------------------------- |
| JWT中间件        | 权限管理系统                 |
| 数据库回调       | MarkDown渲染                 |
| 验证码生成及校验 | 标签系统                     |
| 登录,注册功能    | 上传文件                     |
| 文章增删改查     | 自定义验证器                 |
|                  | LetsEncrypt 证书自动签发部署 |
|                  | 日志中间件                   |
|                  | 链路追踪                     |



