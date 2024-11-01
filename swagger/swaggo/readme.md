执行目录下 init.sh 初始化 swagger.json，再启动项目

# swaggo
## 安装 swaggo
```
go install github.com/swaggo/swag/cmd/swag@latest
```
会将执行程序安装在你的 GOPATH/bin 下。
## 初始化 / 更新文档
执行 init.sh 脚本，在 docs 目录下重新生成 docs.go 与 json 文件。
```
sh init.sh
```
## 启动
run main() 启动服务，访问 `http://ip:port/swagger/index.html`。