# igolang
golang web 脚手架，纯后端项目，封装了一些常用的项目基本功能，主要包括：
- ibeego：基于 beego 框架扩展，封装了 config，controllers，orm，logger，routers 功能； 
- igin： 基于 gin 框架扩展，整合 iconfig，ilogger，iorm 等模块；
- iecho：基于 echo 框架扩展，整合 iconfig，ilogger，iorm 等模块；
- iconfig：配置读取，可选择使用 jinzhu/configor 或者 viper；
- ilogger：日志打印，日志文件分割，基于 zap 扩展；
- iorm：提供实体泛型配置，通用 crud 方法，事务，软删除机制，基于 gorm 扩展；
- swagger：web 框架接口文档生成。

使用示例请参考各模块目录下 readme.md。