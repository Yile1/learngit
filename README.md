# techtrainingcamp-AppUpgrade
 AppUpgrade By Group 12


### 参考文档
[Gin中文文档](https://gin-gonic.com/zh-cn/docs)

[Go语言工程代码示例](https://github.com/haydenzhourepo/gin-vue-gin-essential)

[白名单的一种方式](https://www.zhihu.com/question/343047416/answer/809796259)

[Gorm数据库连接](https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/gorm/)

[Gin + Gorm + Redis(1)](https://zhuanlan.zhihu.com/p/147663215)

[Gin + Gorm + Redis(2)](https://zhuanlan.zhihu.com/p/148423815)

[Gin + Gorm + Redis(3)](https://zhuanlan.zhihu.com/p/148670832)

[Redis + Mysql配合使用](https://www.cnblogs.com/jocongmin/articles/7879470.html)

### 待解决问题
1. Mysql高并发（我们是读多写少）（读写分离加锁）
2. 白名单如何控制（我现在有个初步的想法待讨论）
3. 服务端接口部署
4. Redis预热（读入白名单数据）（Mysql的数据用Json转换？）
5. 配置与检查接口分离（安全性）

### 接下来的任务
1. 配置接口与页面（新建规则、暂停规则、重启规则、删除规则）（分离？）（配置保证安全性）
2. 版本号比较（就写一个函数，工作量很小）√
3. Mysql与Redis的配合，Redis预热加载数据（有一定难度）
4. Mysql读写锁（读写分离）
5. 服务器部署接口与云数据库

### 需要关注的问题
1. 参数校验的问题，避免Redis缓存穿透
2. Redis大Key
3. Redis热点Rule存储与更新

### 代码框架

> @csy 下面是我就着我自己对咱们这个项目的理解写的一个文档，因为我有一些地方也不是特别理解，如果同学们发现有什么问题麻烦改一下，3q

```
|- main.go: 程序代码入口，做初始化工作并让gin跑起来
|
|- router.go: 路由，根据GET或POST方法以及输入路由到不同的函数
|
|- common
|   |- database.go: 与mysql数据库连接
|   |- redis.go: redis预热并连接
|
|- config
|   |- application.yml: 配置文件(用了viper)
|
|- controller
|   |- pong_controller.go: ping函数-检查是否ping通，Form函数-检查表单
|   |- rule_controller.go: 与rule有关的几个函数，检查rule,获取rule,添加和删除rule,启用和禁用rule
|                           (这些都是与rule有关的行为，也许可以放在rule.go里面作为一个rule的class?)
|
|- front-end
|   |- appupdate-admin: 管理员接口部分(这个我不太了解)
|
|- html
|   |- form.html: 一个表单页面
|
|- middlewares 中间件
|   |- cors.go: 前端有关接口?
|
|- model
|   |- rule.go: rule结构体，定义了规则的字段
|
|- resource
|   |- whiteList.txt: 白名单
|
|- utils
|   |- general_tools: 一些其他的发可以用到的函数
```


