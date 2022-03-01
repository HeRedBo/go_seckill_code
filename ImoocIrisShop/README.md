# ImoocIrisShop
慕课网秒杀项目
使用 `iris` 框架搭建的 Web 商城项目
## 软件架构
使用 Golang `iris`框架搭建，同时使用了 `RabbitMQ`,`MYSQL` 完成技术搭建，使用 `gorm` 包操作数据库
目前完成如下功能
- 商品信息管理
- 秒杀订单管理
- 用户注册与登录
- 商品秒杀功能 （使用RabbitMQ 控制秒杀）


## 安装
### 必须

1. Mysql  > 5.6 
2. RabbitMQ 

```
-- 下载代码
git clone https://github.com/HeRedBo/go_seckill_code 
-- 进入下载目录 的 go_seckill_code 目录 
cd ImoocIrisShop
-- 安装必要的包 
go mod tidy 
-- 复制包的项目 vendor 目录
go mod vendor  
```

### 准备

创建数据库 执行项目目录 databases `database.sql` 的相关SQL文件 修改 项目 `conf/app.ini` 相关MYSQL配置
创建一个 RbbbitMQ 可访问的账号秘密 修改 项目 `rabbitmq\rabbitmq.go` 的  `MQURL` 配置项即可


创建 

### 配置``

你应该修改 `conf/app.ini` 配置文件

```
[database]
Type = mysql
User = root
Password = root
Host = 127.0.0.1:3306
Name = imooc_shop
TablePrefix =

```


### 运行

backend   fronted
进入项目根目录 
``` 
--  启动后台部分  进入项目 backend 目录 执行  
go run main.go 
-- 启动项目前端 进入 fronted 目录 执行 
go run main.go 

```

