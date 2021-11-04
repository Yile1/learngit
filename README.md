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
