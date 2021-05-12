# go-vben-admin
项目简介
go-vben-admin 是一个中后台应用框架，基于（gin, gorm, Casbin,zap,Jwt, vben, antd...）实现。

遵循 RESTful API 设计规范

基于 Go WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪 ID 等）

基于 Casbin 的 RBAC 访问控制模型

基于 JWT 认证

基于 GORM 的数据库存储，可扩展多种类型数据库

#项目拉取

`git clone --recursive  https://github.com/coderchiang/go-vben-admin.git  `    

注意，项目中使用了git子项目管理，所以git clone 时必须携带--recursive参数




#项目启动


拉取成功后进入项目根目录

`cd  go-ven-admin`

go-ven-admin目录 执行如下命令 赋予可执行权限

`chmod a+x up.sh `

执行启动脚本

`./up.sh`

到这里项目已经启动完成，启动逻辑在shell脚本中有明确解释，如果部分启动函数对您没用，可以注释掉即可
##项目演示地址


http://2wm.top

###目录结构说明
暂时还没时间写，先放下吧

## 交流

`go-vben-Admin` 是完全开源免费的项目，在帮助开发者更方便地进行中大型管理系统开发，同时也提供 QQ 交流群使用问题欢迎在群内提问。

- QQ 群 `1055067008`

如果本项目对您有帮助，请记得star！您的star是我们长期维护的动力。
## License

[MIT © go-vben-admin](./LICENSE)

