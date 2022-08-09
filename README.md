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
| 权限管理: **Casbin** | **Casbin**: https://casbin.org/ |
| **Casbin-gorm-adapter** | **Casbin-gorm-adapter**: https://github.com/casbin/gorm-adapter |
| 支付功能:**Gopay** | **Gopay**: https://github.com/go-pay/gopay |
| 雪花算法:**SnowFlake** | **SnowFlake**: https://github.com/bwmarrin/snowflake |

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
  # 上传服务
  UploadSavePath: storage/uploads
  UploadServerUrl: http://127.0.0.1:8080/static
  UploadImageMaxSize: 5  # MB
  UploadImageAllowExts:
    - .jpg
    - .jpeg
    - .png
  UploadMarkdownAllowExts:
    - .md
  QrCodeSavePath: storage/qrcode
  PrefixUrl: 127.0.0.1:8080
Database:
  DBType: mysql
  Username: username  # 填写你的数据库账号
  Password: password  # 填写你的数据库密码
  Host: host:3306 # 填写你的数据库地址
  DBName: blog
  TablePrefix:
  Charset: utf8
  ParseTime: True
  LogMode: info
  MaxIdleConns: 10
  MaxOpenConns: 30
Jwt:
  Secret: initcool # 密钥
  Issuer: blog.nmslwsnd.com
  Expire: 72000
Pay:
  AlipayPrivateKey: # 支付宝私钥
  AlipayAppId: 2021000119643838
  AlipayPublicKey: # 支付宝公钥
Redis:
  Host: 127.0.0.1:6379
  Password: 
  MaxIdle: 30 # 最大空闲连接数
  MaxActive: 0 # 在给定时间内，允许分配的最大连接数（当为零时，没有限制）
  IdleTimeout: 0 # 在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
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

| 已实现功能         | 未实现功能   |
| ------------------ | ------------ |
| JWT中间件          | 链路追踪     |
| 数据库回调         | MarkDown文件绑定文章ID |
| 验证码生成及校验   | 日志中间件 |
| 登录,注册功能      |    |
| 文章增删改查       |                        |
| 上传文件           |              |
| api文档生成        |              |
| 接口限流           |              |
| 系统管理员识别     |              |
|自定义验证器  | |
| casbin权限管理系统 |              |
| 支付宝支付         |              |
| MarkDown渲染                   |              |
| 标签系统 | |



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

### Casbin 注解:

```yaml
# Casbin-gorm-adapter 
# 即使使用的是Mysql数据库,Casbin-gorm-adapter也需要安装sqlserver及postgres依赖
go get gorm.io/driver/postgres
go get gorm.io/driver/sqlserver
# casbin数据库字段解析
- ptype: 类型: g代表角色继承关系,p代表路由访问控制策略
- v0: RoleID 用于判断用户组别
- v1: /api/v1/* 代表RoleID可以访问v1内的路由
- V2: GET 支持RoleID可以进行的访问方式
```

### bootstrap初始化 注解：

```yaml
# 在main.go下导入bootstrap
_ "github.com/limitcool/blog/bootstrap"
```

```json
// 注册请求示例
{
    "username":"admin",
    "password":"password",
    "profile":{
        "desc":"个人信息描述",
        "img":"图片URL"
    }
}
```

### air热重启

```bash
# 安装air
go install github.com/cosmtrek/air@latest
# 新建文件.air.conf
touch .air.conf
# 在.air.conf添加以下内容
go run main.go
# 运行air
air
```

