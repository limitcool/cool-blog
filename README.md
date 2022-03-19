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
# swagger 初始化
swag init
# 默认接口文档地址
http://127.0.0.1:8080/swagger/index.html
```

### 功能列表

| 已实现功能       | 未实现功能                   |
| ---------------- | ---------------------------- |
| JWT中间件        | 权限管理系统                 |
| 数据库回调       | MarkDown渲染                 |
| 验证码生成及校验 | 标签系统                     |
| 登录,注册功能    | 日志中间件                   |
| 文章增删改查     | 自定义验证器                 |
| 上传文件         | LetsEncrypt 证书自动签发部署 |
| api文档生成      | 链路追踪                     |



##  Swagger 注解：



```yaml
# 需要确保导入了生成的docs/docs.go文件，这样特定的配置文件才会被初始化
_ "github.com/limitcool/blog/docs"
```



| 注解     | 描述                                                         |
| -------- | ------------------------------------------------------------ |
| @Summary | 摘要                                                         |
| @Produce | API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等 |
| @Param   | 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释 |
| @Success | 响应成功，从左到右分别为：状态码、参数类型、数据类型、注释   |
| @Failure | 响应失败，从左到右分别为：状态码、参数类型、数据类型、注释   |
| @Router  | 路由，从左到右分别为：路由地址，HTTP 方法                    |
