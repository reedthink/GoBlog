## 使用Go开发，遵循RESTful风格设计API
功能：用户注册、登录，标签和文章的增删查改

## 待完善：
- 搜索功能
- 实现本地图片的上传，七牛SDK上传图片以及评论功能。

## 文件夹内容说明
- conf: 配置文件
- docs: API文档生成
- middleware: 中间件，包含jwt认证功能
- models: 操作数据库
- pkg: 各种功能包
- routes: 主要代码
- runtime: 运行日志

 
## 使用的开源库:
- gin
- gorm
- jwt-go
- swaggo
- viper


## 如何启动：
`go build && ./GoBlog`

## 查看API文档: ~~http://114.116.239.137:8081/swagger/index.html~~
